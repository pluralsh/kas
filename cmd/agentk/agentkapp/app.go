package agentkapp

import (
	"bytes"
	"context"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/ash2k/stager"
	"github.com/go-logr/zapr"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/spf13/cobra"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v15/cmd"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v15/internal/api"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v15/internal/module/agent_configuration/rpc"
	gitlab_access_rpc "gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v15/internal/module/gitlab_access/rpc"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v15/internal/module/gitops/agent/chartops"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v15/internal/module/gitops/agent/manifestops"
	kubernetes_api_agent "gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v15/internal/module/kubernetes_api/agent"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v15/internal/module/modagent"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v15/internal/module/modshared"
	observability_agent "gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v15/internal/module/observability/agent"
	reverse_tunnel_agent "gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v15/internal/module/reverse_tunnel/agent"
	starboard_vulnerability "gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v15/internal/module/starboard_vulnerability/agent"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v15/internal/tool/errz"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v15/internal/tool/grpctool"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v15/internal/tool/httpz"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v15/internal/tool/logz"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v15/internal/tool/retry"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v15/internal/tool/tlstool"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v15/internal/tool/wstunnel"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v15/pkg/agentcfg"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zapgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/encoding/gzip"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/keepalive"
	core_v1 "k8s.io/api/core/v1"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes/scheme"
	client_core_v1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/record"
	"k8s.io/klog/v2"
	"k8s.io/kubectl/pkg/cmd/util"
	"nhooyr.io/websocket"
)

const (
	defaultLogLevel     agentcfg.LogLevelEnum = 0 // whatever is 0 is the default value
	defaultGrpcLogLevel                       = agentcfg.LogLevelEnum_error

	defaultObservabilityListenNetwork = "tcp"
	defaultObservabilityListenAddress = ":8080"
	defaultMaxMessageSize             = 10 * 1024 * 1024
	agentName                         = "gitlab-agent"

	envVarPodNamespace       = "POD_NAMESPACE"
	envVarPodName            = "POD_NAME"
	envVarServiceAccountName = "SERVICE_ACCOUNT_NAME"

	getConfigurationInitBackoff   = 10 * time.Second
	getConfigurationMaxBackoff    = 5 * time.Minute
	getConfigurationResetDuration = 10 * time.Minute
	getConfigurationBackoffFactor = 2.0
	getConfigurationJitter        = 1.0
)

type App struct {
	Log          *zap.Logger
	LogLevel     zap.AtomicLevel
	GrpcLogLevel zap.AtomicLevel
	AgentMeta    *modshared.AgentMeta
	AgentId      *AgentIdHolder
	// KasAddress specifies the address of kas.
	KasAddress                 string
	KasCACertFile              string
	KasHeaders                 []string
	ServiceAccountName         string
	ObservabilityListenNetwork string
	ObservabilityListenAddress string
	ObservabilityCertFile      string
	ObservabilityKeyFile       string
	TokenFile                  string
	K8sClientGetter            genericclioptions.RESTClientGetter
}

func (a *App) Run(ctx context.Context) (retErr error) {
	// TODO Tracing
	tp := trace.NewNoopTracerProvider()
	p := propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{})

	// Construct gRPC connection to gitlab-kas
	kasConn, err := a.constructKasConnection(ctx, tp, p)
	if err != nil {
		return err
	}
	defer errz.SafeClose(kasConn, &retErr)

	// Construct internal gRPC server
	internalSrv, err := newInternalServer(a.Log, tp, p) // nolint: contextcheck
	if err != nil {
		return err
	}
	defer errz.SafeClose(internalSrv.conn, &retErr)

	// Construct Kubernetes tools.
	k8sFactory := util.NewFactory(a.K8sClientGetter)
	kubeClient, err := k8sFactory.KubernetesClientSet()
	if err != nil {
		return err
	}

	// Construct event recorder
	eventBroadcaster := record.NewBroadcaster()
	eventRecorder := eventBroadcaster.NewRecorder(scheme.Scheme, core_v1.EventSource{Component: agentName})

	// Construct leader runner
	lr := newLeaderRunner(&leaseLeaderElector{
		namespace: a.AgentMeta.PodNamespace,
		name: func(ctx context.Context) (string, error) {
			id, err := a.AgentId.get(ctx) // nolint: govet
			if err != nil {
				return "", err
			}
			// We use agent id as part of lock name so that agentk Pods of different id don't compete with
			// each other. Only Pods with same agent id should compete for a lock. Put differently, agentk Pods
			// with same agent id have the same lock name but with different id have different lock name.
			return fmt.Sprintf("agent-%d-lock", id), nil
		},
		identity:           a.AgentMeta.PodName,
		coordinationClient: kubeClient.CoordinationV1(),
		eventRecorder:      eventRecorder,
	})

	// Construct agent modules
	beforeServersModules, afterServersModules, err := a.constructModules(internalSrv.server, kasConn, internalSrv.conn, k8sFactory, lr)
	if err != nil {
		return err
	}
	runner := a.newModuleRunner(kasConn)
	beforeServersModulesRun := runner.RegisterModules(beforeServersModules)
	afterServersModulesRun := runner.RegisterModules(afterServersModules)

	// Start events processing pipeline.
	loggingWatch := eventBroadcaster.StartStructuredLogging(0)
	defer loggingWatch.Stop()
	eventBroadcaster.StartRecordingToSink(&client_core_v1.EventSinkImpl{Interface: kubeClient.CoreV1().Events("")})
	defer eventBroadcaster.Shutdown()

	// Start things up. Stages are shut down in reverse order.
	return stager.RunStages(ctx,
		func(stage stager.Stage) {
			stage.Go(func(ctx context.Context) error {
				// Start leader runner.
				lr.Run(ctx)
				return nil
			})
		},
		func(stage stager.Stage) {
			// Start modules.
			stage.Go(beforeServersModulesRun)
		},
		func(stage stager.Stage) {
			// Start internal gRPC server. It is used by internal modules, so it is shut down after them.
			internalSrv.Start(stage)
		},
		func(stage stager.Stage) {
			// Start modules that use internal server.
			stage.Go(afterServersModulesRun)
		},
		func(stage stager.Stage) {
			// Start configuration refresh.
			stage.Go(runner.RunConfigurationRefresh)
		},
	)
}

func (a *App) newModuleRunner(kasConn *grpc.ClientConn) *moduleRunner {
	return &moduleRunner{
		log: a.Log,
		configurationWatcher: &rpc.ConfigurationWatcher{
			Log:       a.Log,
			AgentMeta: a.AgentMeta,
			Client:    rpc.NewAgentConfigurationClient(kasConn),
			PollConfig: retry.NewPollConfigFactory(0, retry.NewExponentialBackoffFactory(
				getConfigurationInitBackoff,
				getConfigurationMaxBackoff,
				getConfigurationResetDuration,
				getConfigurationBackoffFactor,
				getConfigurationJitter,
			)),
			ConfigPreProcessor: func(data rpc.ConfigurationData) error {
				return a.AgentId.set(data.Config.AgentId)
			},
		},
	}
}

func (a *App) constructModules(internalServer *grpc.Server, kasConn, internalServerConn grpc.ClientConnInterface,
	k8sFactory util.Factory, lr *leaderRunner) ([]modagent.Module, []modagent.Module, error) {
	accessClient := gitlab_access_rpc.NewGitlabAccessClient(kasConn)
	factories := []modagent.Factory{
		&observability_agent.Factory{
			LogLevel:            a.LogLevel,
			GrpcLogLevel:        a.GrpcLogLevel,
			ListenNetwork:       a.ObservabilityListenNetwork,
			ListenAddress:       a.ObservabilityListenAddress,
			CertFile:            a.ObservabilityCertFile,
			KeyFile:             a.ObservabilityKeyFile,
			DefaultGrpcLogLevel: defaultGrpcLogLevel,
		},
		&manifestops.Factory{},
		&chartops.Factory{},
		&starboard_vulnerability.Factory{},
		&reverse_tunnel_agent.Factory{
			InternalServerConn: internalServerConn,
		},
		&kubernetes_api_agent.Factory{},
	}
	var beforeServersModules, afterServersModules []modagent.Module
	for _, f := range factories {
		moduleName := f.Name()
		module, err := f.New(&modagent.Config{
			Log:       a.Log.With(logz.ModuleName(moduleName)),
			AgentMeta: a.AgentMeta,
			Api: &agentAPI{
				moduleName: moduleName,
				agentId:    a.AgentId,
				client:     accessClient,
			},
			K8sUtilFactory:     k8sFactory,
			KasConn:            kasConn,
			Server:             internalServer,
			AgentName:          agentName,
			ServiceAccountName: a.ServiceAccountName,
		})
		if err != nil {
			return nil, nil, err
		}
		module = lr.MaybeWrapModule(module)
		phase := f.StartStopPhase()
		switch phase {
		case modshared.ModuleStartBeforeServers:
			beforeServersModules = append(beforeServersModules, module)
		case modshared.ModuleStartAfterServers:
			afterServersModules = append(afterServersModules, module)
		default:
			return nil, nil, fmt.Errorf("invalid StartStopPhase from factory %s: %d", moduleName, phase)
		}
	}
	return beforeServersModules, afterServersModules, nil
}

func (a *App) constructKasConnection(ctx context.Context, tp trace.TracerProvider, p propagation.TextMapPropagator) (*grpc.ClientConn, error) {
	tokenData, err := os.ReadFile(a.TokenFile)
	if err != nil {
		return nil, fmt.Errorf("token file: %w", err)
	}
	tokenData = bytes.TrimSuffix(tokenData, []byte{'\n'})
	tlsConfig, err := tlstool.DefaultClientTLSConfigWithCACert(a.KasCACertFile)
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(a.KasAddress)
	if err != nil {
		return nil, fmt.Errorf("invalid gitlab-kas address: %w", err)
	}
	kasHeaders, err := parseHeaders(a.KasHeaders)
	if err != nil {
		return nil, err
	}
	userAgent := fmt.Sprintf("%s/%s/%s", agentName, a.AgentMeta.Version, a.AgentMeta.CommitId)
	opts := []grpc.DialOption{
		grpc.WithDefaultCallOptions(grpc.UseCompressor(gzip.Name)),
		grpc.WithUserAgent(userAgent),
		// keepalive.ClientParameters must be specified at least as large as what is allowed by the
		// server-side grpc.KeepaliveEnforcementPolicy
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			// kas allows min 20 seconds, trying to stay below 60 seconds (typical load-balancer timeout) and
			// above kas' server keepalive Time so that kas pings the client sometimes. This helps mitigate
			// reverse-proxies' enforced server response timeout.
			Time:                55 * time.Second,
			PermitWithoutStream: true,
		}),
		grpc.WithChainStreamInterceptor(
			grpc_prometheus.StreamClientInterceptor,
			otelgrpc.StreamClientInterceptor(otelgrpc.WithTracerProvider(tp), otelgrpc.WithPropagators(p)),
			grpctool.StreamClientValidatingInterceptor,
		),
		grpc.WithChainUnaryInterceptor(
			grpc_prometheus.UnaryClientInterceptor,
			otelgrpc.UnaryClientInterceptor(otelgrpc.WithTracerProvider(tp), otelgrpc.WithPropagators(p)),
			grpctool.UnaryClientValidatingInterceptor,
		),
	}
	var addressToDial string
	// "grpcs" is the only scheme where encryption is done by gRPC.
	// "wss" is secure too but gRPC cannot know that, so we tell it it's not.
	secure := u.Scheme == "grpcs"
	switch u.Scheme {
	case "ws", "wss":
		addressToDial = a.KasAddress
		dialer := net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}
		kasHeaders.Set(httpz.UserAgentHeader, userAgent)
		opts = append(opts, grpc.WithContextDialer(wstunnel.DialerForGRPC(defaultMaxMessageSize, &websocket.DialOptions{
			HTTPClient: &http.Client{
				Transport: &http.Transport{
					Proxy:                 http.ProxyFromEnvironment,
					DialContext:           dialer.DialContext,
					TLSClientConfig:       tlsConfig,
					MaxIdleConns:          10,
					IdleConnTimeout:       90 * time.Second,
					TLSHandshakeTimeout:   10 * time.Second,
					ResponseHeaderTimeout: 20 * time.Second,
				},
				CheckRedirect: func(req *http.Request, via []*http.Request) error {
					return http.ErrUseLastResponse
				},
			},
			HTTPHeader: kasHeaders,
		})))
	case "grpc":
		addressToDial = grpcHostWithPort(u)
		opts = append(opts, grpc.WithPerRPCCredentials(grpctool.NewHeaderMetadata(kasHeaders, !secure)))
	case "grpcs":
		addressToDial = grpcHostWithPort(u)
		opts = append(opts,
			grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)),
			grpc.WithPerRPCCredentials(grpctool.NewHeaderMetadata(kasHeaders, !secure)),
		)
	default:
		return nil, fmt.Errorf("unsupported scheme in GitLab Kubernetes Agent Server address: %q", u.Scheme)
	}
	if !secure {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}
	opts = append(opts, grpc.WithPerRPCCredentials(grpctool.NewTokenCredentials(api.AgentToken(tokenData), !secure)))
	conn, err := grpc.DialContext(ctx, addressToDial, opts...)
	if err != nil {
		return nil, fmt.Errorf("gRPC.dial: %w", err)
	}
	return conn, nil
}

func NewCommand() *cobra.Command {
	kubeConfigFlags := genericclioptions.NewConfigFlags(true)
	a := App{
		AgentMeta: &modshared.AgentMeta{
			Version:  cmd.Version,
			CommitId: cmd.Commit,
		},
		AgentId:            NewAgentIdHolder(),
		ServiceAccountName: os.Getenv(envVarServiceAccountName),
		K8sClientGetter:    kubeConfigFlags,
	}
	c := &cobra.Command{
		Use:   "agentk",
		Short: "GitLab Agent for Kubernetes",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) (retErr error) {
			podNs := os.Getenv(envVarPodNamespace)
			if podNs == "" {
				return fmt.Errorf("%s environment variable is required but is empty", envVarPodNamespace)
			}
			podName := os.Getenv(envVarPodName)
			if podName == "" {
				return fmt.Errorf("%s environment variable is required but is empty", envVarPodName)
			}
			a.AgentMeta.PodNamespace = podNs
			a.AgentMeta.PodName = podName
			lockedSyncer := zapcore.Lock(logz.NoSync(os.Stderr))
			var err error
			a.Log, a.LogLevel, err = a.logger(defaultLogLevel, lockedSyncer)
			if err != nil {
				return err
			}
			defer errz.SafeCall(a.Log.Sync, &retErr)

			var grpcLog *zap.Logger
			grpcLog, a.GrpcLogLevel, err = a.logger(defaultGrpcLogLevel, lockedSyncer)
			if err != nil {
				return err
			}
			defer errz.SafeCall(grpcLog.Sync, &retErr)

			grpclog.SetLoggerV2(zapgrpc.NewLogger(grpcLog)) // pipe gRPC logs into zap
			// Kubernetes uses klog so here we pipe all logs from it to our logger via an adapter.
			klog.SetLogger(zapr.NewLogger(a.Log))

			return a.Run(cmd.Context())
		},
		SilenceErrors: true,
		SilenceUsage:  true,
	}
	f := c.Flags()
	f.StringVar(&a.KasAddress, "kas-address", "", "GitLab Kubernetes Agent Server address")
	f.StringVar(&a.TokenFile, "token-file", "", "File with access token")

	f.StringVar(&a.KasCACertFile, "ca-cert-file", "", "Optional file with X.509 certificate authority certificate in PEM format. Used for verifying cert of agent server")
	f.StringArrayVar(&a.KasHeaders, "kas-header", []string{}, "Optional HTTP headers to set when connecting to the agent server")

	f.StringVar(&a.ObservabilityListenNetwork, "observability-listen-network", defaultObservabilityListenNetwork, "Observability network to listen on")
	f.StringVar(&a.ObservabilityListenAddress, "observability-listen-address", defaultObservabilityListenAddress, "Observability address to listen on")
	f.StringVar(&a.ObservabilityCertFile, "observability-cert-file", "", "Optional file with X.509 certificate in PEM format. User for observability endpoint TLS")
	f.StringVar(&a.ObservabilityKeyFile, "observability-key-file", "", "Optional file with X.509 key in PEM format. User for observability endpoint TLS")

	kubeConfigFlags.AddFlags(f)
	cobra.CheckErr(c.MarkFlagRequired("kas-address"))
	cobra.CheckErr(c.MarkFlagRequired("token-file"))
	return c
}

func grpcHostWithPort(u *url.URL) string {
	port := u.Port()
	if port != "" {
		return u.Host
	}
	switch u.Scheme {
	case "grpc":
		return net.JoinHostPort(u.Host, "80")
	case "grpcs":
		return net.JoinHostPort(u.Host, "443")
	default:
		// Function called with unknown scheme, just return the original host.
		return u.Host
	}
}

func parseHeaders(raw []string) (http.Header, error) {
	header := http.Header{}
	for _, h := range raw {
		k, v, ok := strings.Cut(h, ":")
		if !ok {
			return nil, fmt.Errorf("invalid header supplied: %s", h)
		}
		k, v = strings.Trim(k, " "), strings.Trim(v, " ")
		if len(k) < 1 || len(v) < 1 {
			return nil, fmt.Errorf("invalid header supplied: %s", h)
		}
		header.Add(k, v)
	}
	return header, nil
}
