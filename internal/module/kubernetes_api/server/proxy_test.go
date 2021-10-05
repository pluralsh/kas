package server

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	gapi "gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v14/internal/gitlab/api"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v14/internal/module/kubernetes_api/rpc"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v14/internal/module/modserver"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v14/internal/tool/cache"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v14/internal/tool/grpctool"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v14/internal/tool/prototool"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v14/internal/tool/testing/matcher"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v14/internal/tool/testing/mock_gitlab"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v14/internal/tool/testing/mock_kubernetes_api"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v14/internal/tool/testing/mock_modserver"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v14/internal/tool/testing/mock_usage_metrics"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v14/internal/tool/testing/testhelpers"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v14/pkg/agentcfg"
	"gitlab.com/gitlab-org/labkit/metrics"
	"go.uber.org/zap/zaptest"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"k8s.io/apimachinery/pkg/util/wait"
)

const (
	jobToken        = "asdfgasdfxadf"
	requestPath     = "/api/bla"
	requestPayload  = "asdfndaskjfadsbfjsadhvfjhavfjasvf"
	responsePayload = "jknkjnjkasdnfkjasdnfkasdnfjnkjn"
	queryParamValue = "query-param-value with a space"
	queryParamName  = "q with a space"
)

func TestProxy_JobTokenErrors(t *testing.T) {
	tests := []struct {
		name     string
		auth     []string
		response string
	}{
		{
			name:     "missing header",
			response: "Authorization header: expecting token\n",
		},
		{
			name:     "multiple headers",
			auth:     []string{"a", "b"},
			response: "Authorization header: expecting a single header, got 2\n",
		},
		{
			name:     "invalid format1",
			auth:     []string{"Token asdfadsf"},
			response: "Authorization header: expecting Bearer token\n",
		},
		{
			name:     "invalid format2",
			auth:     []string{"Bearer asdfadsf"},
			response: "Authorization header: invalid value\n",
		},
		{
			name:     "invalid agent id",
			auth:     []string{"Bearer ci:asdf:as"},
			response: "Authorization header: failed to parse: strconv.ParseInt: parsing \"asdf\": invalid syntax\n",
		},
		{
			name:     "empty token",
			auth:     []string{"Bearer ci:1:"},
			response: "Authorization header: empty token\n",
		},
		{
			name:     "unknown token type",
			auth:     []string{"Bearer blabla:1:asd"},
			response: "Authorization header: unknown token type\n",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, _, client, req, _ := setupProxyWithHandler(t, "/", func(w http.ResponseWriter, r *http.Request) {
				t.Fail() // unexpected invocation
			})
			req.Header.Del("Authorization")
			if len(tc.auth) > 0 {
				req.Header["Authorization"] = tc.auth
			}
			resp, err := client.Do(req)
			require.NoError(t, err)
			defer resp.Body.Close()
			assert.EqualValues(t, http.StatusUnauthorized, resp.StatusCode)
			assert.Equal(t, tc.response, string(readAll(t, resp.Body)))
		})
	}
}

func TestProxy_AllowedAgentsError(t *testing.T) {
	tests := []struct {
		allowedAgentsHttpStatus int
		expectedHttpStatus      int
		captureErr              bool
	}{
		{
			allowedAgentsHttpStatus: http.StatusUnauthorized, // token is invalid
			expectedHttpStatus:      http.StatusUnauthorized,
		},
		{
			allowedAgentsHttpStatus: http.StatusForbidden, // token is forbidden
			expectedHttpStatus:      http.StatusForbidden,
		},
		{
			allowedAgentsHttpStatus: http.StatusNotFound, // agent is not found
			expectedHttpStatus:      http.StatusNotFound,
		},
		{
			allowedAgentsHttpStatus: http.StatusBadGateway, // some weird error
			expectedHttpStatus:      http.StatusInternalServerError,
			captureErr:              true,
		},
	}
	for _, tc := range tests {
		t.Run(strconv.Itoa(tc.allowedAgentsHttpStatus), func(t *testing.T) {
			api, _, client, req, _ := setupProxyWithHandler(t, "/", func(w http.ResponseWriter, r *http.Request) {
				assertToken(t, r)
				w.WriteHeader(tc.allowedAgentsHttpStatus)
			})
			if tc.captureErr {
				api.EXPECT().
					HandleProcessingError(gomock.Any(), gomock.Any(), testhelpers.AgentId, gomock.Any(),
						matcher.ErrorEq(fmt.Sprintf("HTTP status code: %d", tc.allowedAgentsHttpStatus)))
			}
			resp, err := client.Do(req)
			require.NoError(t, err)
			defer resp.Body.Close()
			assert.EqualValues(t, tc.expectedHttpStatus, resp.StatusCode)
			assert.Empty(t, string(readAll(t, resp.Body)))
		})
	}
}

func TestProxy_NoExpectedUrlPathPrefix(t *testing.T) {
	_, _, client, req, _ := setupProxyWithHandler(t, "/bla/", configGitLabHandler(t, nil))
	req.URL.Path = requestPath
	resp, err := client.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()
	assert.EqualValues(t, http.StatusBadRequest, resp.StatusCode)
	assert.Empty(t, string(readAll(t, resp.Body)))
}

func TestProxy_ForbiddenAgentId(t *testing.T) {
	_, _, client, req, _ := setupProxy(t)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s:%d:%s", tokenTypeCi, 15 /* disallowed id */, jobToken))
	resp, err := client.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()
	assert.EqualValues(t, http.StatusForbidden, resp.StatusCode)
	assert.Empty(t, string(readAll(t, resp.Body)))
}

func TestProxy_HappyPath(t *testing.T) {
	tests := []struct {
		name          string
		urlPathPrefix string
		config        *gapi.Configuration
		expectedExtra *rpc.HeaderExtra
	}{
		{
			name:          "no prefix",
			urlPathPrefix: "/",
			expectedExtra: &rpc.HeaderExtra{
				ImpConfig: &rpc.ImpersonationConfig{},
			},
		},
		{
			name:          "with prefix",
			urlPathPrefix: "/bla/",
			expectedExtra: &rpc.HeaderExtra{
				ImpConfig: &rpc.ImpersonationConfig{},
			},
		},
		{
			name:          "impersonate agent",
			urlPathPrefix: "/",
			config: &gapi.Configuration{
				AccessAs: &agentcfg.CiAccessAsCF{
					As: &agentcfg.CiAccessAsCF_Agent{
						Agent: &agentcfg.CiAccessAsAgentCF{},
					},
				},
			},
			expectedExtra: &rpc.HeaderExtra{
				ImpConfig: &rpc.ImpersonationConfig{},
			},
		},
		{
			name:          "impersonate",
			urlPathPrefix: "/",
			config: &gapi.Configuration{
				AccessAs: &agentcfg.CiAccessAsCF{
					As: &agentcfg.CiAccessAsCF_Impersonate{
						Impersonate: &agentcfg.CiAccessAsImpersonateCF{
							Username: "user1",
							Groups:   []string{"g1", "g2"},
							Uid:      "uid",
							Extra: []*agentcfg.ExtraKeyValCF{
								{
									Key: "k1",
									Val: []string{"v1", "v2"},
								},
							},
						},
					},
				},
			},
			expectedExtra: &rpc.HeaderExtra{
				ImpConfig: &rpc.ImpersonationConfig{
					Username: "user1",
					Groups:   []string{"g1", "g2"},
					Uid:      "uid",
					Extra: []*rpc.ExtraKeyVal{
						{
							Key: "k1",
							Val: []string{"v1", "v2"},
						},
					},
				},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			testProxyHappyPath(t, tc.urlPathPrefix, tc.expectedExtra, configGitLabHandler(t, tc.config))
		})
	}
}

func testProxyHappyPath(t *testing.T, urlPathPrefix string, expectedExtra *rpc.HeaderExtra, handler func(http.ResponseWriter, *http.Request)) {
	_, k8sClient, client, req, requestCount := setupProxyWithHandler(t, urlPathPrefix, handler)
	requestCount.EXPECT().Inc()
	mrClient := mock_kubernetes_api.NewMockKubernetesApi_MakeRequestClient(gomock.NewController(t))
	mrCall := k8sClient.EXPECT().
		MakeRequest(gomock.Any()).
		DoAndReturn(func(ctx context.Context, opts ...grpc.CallOption) (rpc.KubernetesApi_MakeRequestClient, error) {
			requireCorrectOutgoingMeta(t, ctx)
			return mrClient, nil
		})
	extra, err := anypb.New(expectedExtra)
	require.NoError(t, err)
	gomock.InOrder(append([]*gomock.Call{mrCall}, mockSendStream(t, mrClient,
		&grpctool.HttpRequest{
			Message: &grpctool.HttpRequest_Header_{
				Header: &grpctool.HttpRequest_Header{
					Request: &prototool.HttpRequest{
						Method: http.MethodPost,
						Header: map[string]*prototool.Values{
							"Req-Header": {
								Value: []string{"x1", "x2"},
							},
							"Accept-Encoding": { // added by the Go client
								Value: []string{"gzip"},
							},
							"User-Agent": {
								Value: []string{"test-agent"},
							},
							"Content-Length": { // added by the Go client
								Value: []string{strconv.Itoa(len(requestPayload))},
							},
							"Via": {
								Value: []string{"gRPC/1.0 sv1"},
							},
						},
						UrlPath: requestPath,
						Query: map[string]*prototool.Values{
							queryParamName: {
								Value: []string{queryParamValue},
							},
						},
					},
					Extra: extra,
				},
			},
		},
		&grpctool.HttpRequest{
			Message: &grpctool.HttpRequest_Data_{
				Data: &grpctool.HttpRequest_Data{
					Data: []byte(requestPayload),
				},
			},
		},
		&grpctool.HttpRequest{
			Message: &grpctool.HttpRequest_Trailer_{
				Trailer: &grpctool.HttpRequest_Trailer{},
			},
		},
	)...)...)
	gomock.InOrder(append([]*gomock.Call{mrCall}, mockRecvStream(mrClient,
		&grpctool.HttpResponse{
			Message: &grpctool.HttpResponse_Header_{
				Header: &grpctool.HttpResponse_Header{
					Response: &prototool.HttpResponse{
						StatusCode: http.StatusOK,
						Status:     http.StatusText(http.StatusOK),
						Header: map[string]*prototool.Values{
							"Resp-Header": {
								Value: []string{"a1", "a2"},
							},
							"Content-Type": {
								Value: []string{"application/octet-stream"},
							},
							"Date": {
								Value: []string{"NOW!"},
							},
						},
					},
				},
			},
		},
		&grpctool.HttpResponse{
			Message: &grpctool.HttpResponse_Data_{
				Data: &grpctool.HttpResponse_Data{
					Data: []byte(responsePayload),
				},
			},
		},
		&grpctool.HttpResponse{
			Message: &grpctool.HttpResponse_Trailer_{
				Trailer: &grpctool.HttpResponse_Trailer{},
			},
		},
	)...)...)

	req.Header.Set("Req-Header", "x1")
	req.Header.Add("Req-Header", "x2")
	req.Header.Set("User-Agent", "test-agent") // added manually to override what is added by the Go client
	resp, err := client.Do(req)
	require.NoError(t, err)
	defer func() {
		assert.NoError(t, resp.Body.Close())
	}()
	assert.EqualValues(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, responsePayload, string(readAll(t, resp.Body)))
	resp.Header.Del("Date")
	assert.NotEmpty(t, resp.Header.Get("X-Request-Id"))
	resp.Header.Del("X-Request-Id")
	assert.Empty(t, cmp.Diff(map[string][]string{
		"Resp-Header":    {"a1", "a2"},
		"Content-Length": {"31"},
		"Content-Type":   {"application/octet-stream"},
		"Via":            {"gRPC/1.0 sv1"},
	}, (map[string][]string)(resp.Header)))
}

func TestProxy_RecvHeaderError(t *testing.T) {
	api, k8sClient, client, req, requestCount := setupProxy(t)
	requestCount.EXPECT().Inc()
	mrClient := mock_kubernetes_api.NewMockKubernetesApi_MakeRequestClient(gomock.NewController(t))
	mrClient.EXPECT().
		Send(gomock.Any()).
		Return(io.EOF) // error on the stream
	gomock.InOrder(
		k8sClient.EXPECT().
			MakeRequest(gomock.Any()).
			Return(mrClient, nil),
		mrClient.EXPECT().
			RecvMsg(gomock.Any()).
			Do(testhelpers.RecvMsg(&grpctool.HttpResponse{
				Message: &grpctool.HttpResponse_Header_{
					Header: &grpctool.HttpResponse_Header{
						Response: &prototool.HttpResponse{
							StatusCode: http.StatusOK,
						},
					},
				},
			})),
		mrClient.EXPECT().
			RecvMsg(gomock.Any()).
			Return(errors.New("expected error")),
		api.EXPECT().
			HandleProcessingError(gomock.Any(), gomock.Any(), testhelpers.AgentId, gomock.Any(), matcher.ErrorEq("expected error")),
	)
	_, err := client.Do(req)               // nolint: bodyclose
	assert.True(t, errors.Is(err, io.EOF)) // dropped connection is io.EOF
}

func TestProxy_ErrorAfterHeaderWritten(t *testing.T) {
	api, k8sClient, client, req, requestCount := setupProxy(t)
	requestCount.EXPECT().Inc()
	mrClient := mock_kubernetes_api.NewMockKubernetesApi_MakeRequestClient(gomock.NewController(t))
	mrClient.EXPECT().
		Send(gomock.Any()).
		Return(io.EOF) // error on the stream
	gomock.InOrder(
		k8sClient.EXPECT().
			MakeRequest(gomock.Any()).
			Return(mrClient, nil),
		mrClient.EXPECT().
			RecvMsg(gomock.Any()).
			Return(errors.New("expected error 2")),
		api.EXPECT().
			HandleProcessingError(gomock.Any(), gomock.Any(), testhelpers.AgentId, gomock.Any(), matcher.ErrorEq("expected error 2")),
	)
	resp, err := client.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()
	assert.EqualValues(t, http.StatusBadGateway, resp.StatusCode)
	assert.Equal(t, "Proxy failed to read response from agent: expected error 2\n", string(readAll(t, resp.Body)))
}

func requireCorrectOutgoingMeta(t *testing.T, ctx context.Context) {
	md, _ := metadata.FromOutgoingContext(ctx)
	vals := md.Get(modserver.RoutingAgentIdMetadataKey)
	require.Len(t, vals, 1)
	agentId, err := strconv.ParseInt(vals[0], 10, 64)
	require.NoError(t, err)
	require.Equal(t, testhelpers.AgentId, agentId)
}

func assertToken(t *testing.T, r *http.Request) bool {
	return assert.Equal(t, jobToken, r.Header.Get("Job-Token"))
}

func setupProxy(t *testing.T) (*mock_modserver.MockApi, *mock_kubernetes_api.MockKubernetesApiClient, *http.Client, *http.Request, *mock_usage_metrics.MockCounter) {
	return setupProxyWithHandler(t, "/", configGitLabHandler(t, nil))
}

func configGitLabHandler(t *testing.T, config *gapi.Configuration) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if !assertToken(t, r) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		testhelpers.RespondWithJSON(t, w, &prototool.JsonBox{Message: &gapi.AllowedAgentsForJob{
			AllowedAgents: []*gapi.AllowedAgent{
				{
					Id: testhelpers.AgentId,
					ConfigProject: &gapi.ConfigProject{
						Id: 5,
					},
					Configuration: config,
				},
			},
			Job: &gapi.Job{
				Id: 1,
			},
			Pipeline: &gapi.Pipeline{
				Id: 2,
			},
			Project: &gapi.Project{
				Id: 3,
			},
			User: &gapi.User{
				Id:       3,
				Username: "testuser",
			},
		}})
	}
}

func setupProxyWithHandler(t *testing.T, urlPathPrefix string, handler func(http.ResponseWriter, *http.Request)) (*mock_modserver.MockApi, *mock_kubernetes_api.MockKubernetesApiClient, *http.Client, *http.Request, *mock_usage_metrics.MockCounter) {
	ctrl := gomock.NewController(t)
	mockApi := mock_modserver.NewMockApi(ctrl)
	k8sClient := mock_kubernetes_api.NewMockKubernetesApiClient(ctrl)
	requestCount := mock_usage_metrics.NewMockCounter(ctrl)

	p := kubernetesApiProxy{
		log:                 zaptest.NewLogger(t),
		api:                 mockApi,
		kubernetesApiClient: k8sClient,
		gitLabClient:        mock_gitlab.SetupClient(t, gapi.AllowedAgentsApiPath, handler),
		allowedAgentsCache:  cache.NewWithError(0, 0, func(err error) bool { return false }),
		requestCount:        requestCount,
		metricsHttpHandlerFactory: func(next http.Handler, opts ...metrics.HandlerOption) http.Handler {
			return next
		},
		serverName:    "sv1",
		urlPathPrefix: urlPathPrefix,
	}
	listener := grpctool.NewDialListener()
	var wg wait.Group
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(func() {
		cancel()
		wg.Wait()
		listener.Close()
	})
	wg.Start(func() {
		assert.NoError(t, p.Run(ctx, listener))
	})
	client := &http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				return listener.DialContext(ctx, addr)
			},
		},
	}
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		"http://any_host_will_do.local"+path.Join(urlPathPrefix, requestPath)+"?"+url.QueryEscape(queryParamName)+"="+url.QueryEscape(queryParamValue),
		strings.NewReader(requestPayload),
	)
	require.NoError(t, err)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s:%d:%s", tokenTypeCi, testhelpers.AgentId, jobToken))
	return mockApi, k8sClient, client, req, requestCount
}

func mockRecvStream(server *mock_kubernetes_api.MockKubernetesApi_MakeRequestClient, msgs ...proto.Message) []*gomock.Call {
	res := make([]*gomock.Call, 0, len(msgs)+1)
	for _, msg := range msgs {
		call := server.EXPECT().
			RecvMsg(gomock.Any()).
			Do(testhelpers.RecvMsg(msg))
		res = append(res, call)
	}
	call := server.EXPECT().
		RecvMsg(gomock.Any()).
		Return(io.EOF)
	res = append(res, call)
	return res
}

func mockSendStream(t *testing.T, client *mock_kubernetes_api.MockKubernetesApi_MakeRequestClient, msgs ...*grpctool.HttpRequest) []*gomock.Call {
	res := make([]*gomock.Call, 0, len(msgs)+1)
	for _, msg := range msgs {
		call := client.EXPECT().
			Send(matcher.ProtoEq(t, msg))
		res = append(res, call)
	}
	res = append(res, client.EXPECT().CloseSend())
	return res
}

func readAll(t *testing.T, r io.Reader) []byte {
	data, err := io.ReadAll(r)
	require.NoError(t, err)
	return data
}
