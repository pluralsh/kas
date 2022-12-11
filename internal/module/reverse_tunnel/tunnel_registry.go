package reverse_tunnel

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v15/internal/api"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v15/internal/module/reverse_tunnel/rpc"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v15/internal/module/reverse_tunnel/tracker"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v15/internal/tool/errz"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v15/internal/tool/grpctool"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v15/internal/tool/logz"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// IOFunc is a function that performs I/O.
// It should be called when not holding a mutex to avoid blocking other goroutines trying to acquire the mutex
// for the duration of the I/O.
type IOFunc func()

type TunnelHandler interface {
	// HandleTunnel is called with server-side interface of the reverse tunnel.
	// It registers the tunnel and blocks, waiting for a request to proxy through the tunnel.
	// The method returns the error value to return to gRPC framework.
	// ctx can be used to unblock the method if the tunnel is not being used already.
	// ctx should be a child of the server's context.
	HandleTunnel(ctx context.Context, agentInfo *api.AgentInfo, server rpc.ReverseTunnel_ConnectServer) error
}

type TunnelFinder interface {
	// FindTunnel finds a tunnel to a matching agentk.
	// It waits for a matching tunnel to proxy a connection through. When a matching tunnel is found, it is returned.
	// It only returns errors from the context or context.Canceled if the finder is shutting down.
	// service and method and gRPC service and method that the agent must support.
	FindTunnel(ctx context.Context, agentId int64, service, method string) (Tunnel, error)
}

type findTunnelRequest struct {
	agentId         int64
	service, method string
	retTun          chan<- *tunnel
}

type TunnelRegistry struct {
	log                 *zap.Logger
	errRep              errz.ErrReporter
	tunnelRegisterer    tracker.Registerer
	ownPrivateApiUrl    string
	tunnelStreamVisitor *grpctool.StreamVisitor

	mu                    sync.Mutex
	tuns                  map[*tunnel]struct{}
	tunsByAgentId         map[int64]map[*tunnel]struct{}
	findRequestsByAgentId map[int64]map[*findTunnelRequest]struct{}
}

func NewTunnelRegistry(log *zap.Logger, errRep errz.ErrReporter, tunnelRegisterer tracker.Registerer, ownPrivateApiUrl string) (*TunnelRegistry, error) {
	tunnelStreamVisitor, err := grpctool.NewStreamVisitor(&rpc.ConnectRequest{})
	if err != nil {
		return nil, err
	}
	return &TunnelRegistry{
		log:                   log,
		errRep:                errRep,
		tunnelRegisterer:      tunnelRegisterer,
		ownPrivateApiUrl:      ownPrivateApiUrl,
		tunnelStreamVisitor:   tunnelStreamVisitor,
		tuns:                  make(map[*tunnel]struct{}),
		tunsByAgentId:         make(map[int64]map[*tunnel]struct{}),
		findRequestsByAgentId: make(map[int64]map[*findTunnelRequest]struct{}),
	}, nil
}

func (r *TunnelRegistry) FindTunnel(ctx context.Context, agentId int64, service, method string) (Tunnel, error) {
	// Buffer 1 to not block on send when a tunnel is found before find request is registered.
	retTun := make(chan *tunnel, 1) // can receive nil from it if Stop() is called
	ftr := &findTunnelRequest{
		agentId: agentId,
		service: service,
		method:  method,
		retTun:  retTun,
	}
	doIO := func() IOFunc {
		r.mu.Lock()
		defer r.mu.Unlock()

		// 1. Check if we have a suitable tunnel
		for tun := range r.tunsByAgentId[agentId] {
			if !tun.agentDescriptor.SupportsServiceAndMethod(service, method) {
				continue
			}
			// Suitable tunnel found!
			tun.state = stateFound
			retTun <- tun // must not block because the reception is below
			return r.unregisterTunnelLocked(tun)
		}

		// 2. No suitable tunnel found, add to the queue
		findRequestsForAgentId := r.findRequestsByAgentId[agentId]
		if findRequestsForAgentId == nil {
			findRequestsForAgentId = make(map[*findTunnelRequest]struct{}, 1)
			r.findRequestsByAgentId[agentId] = findRequestsForAgentId
		}
		findRequestsForAgentId[ftr] = struct{}{}
		return noIO
	}()
	doIO()
	select {
	case <-ctx.Done():
		doIO = func() IOFunc {
			r.mu.Lock()
			defer r.mu.Unlock()
			close(retTun)
			tun := <-retTun // will get nil if there was nothing in the channel or if registry is shutting down.
			if tun != nil {
				// Got the tunnel, but it's too late so return it to the registry.
				return r.onTunnelDoneLocked(tun)
			} else {
				r.deleteFindRequest(ftr)
				return noIO
			}
		}()
		doIO()
		return nil, ctx.Err()
	case tun := <-retTun:
		if tun == nil {
			return nil, status.Error(codes.Unavailable, "kas is shutting down")
		}
		return tun, nil
	}
}

func (r *TunnelRegistry) HandleTunnel(ctx context.Context, agentInfo *api.AgentInfo, server rpc.ReverseTunnel_ConnectServer) error {
	recv, err := server.Recv()
	if err != nil {
		return err
	}
	descriptor, ok := recv.Msg.(*rpc.ConnectRequest_Descriptor_)
	if !ok {
		return status.Errorf(codes.InvalidArgument, "Invalid oneof value type: %T", recv.Msg)
	}
	retErr := make(chan error, 1)
	tun := &tunnel{
		tunnel:              server,
		tunnelStreamVisitor: r.tunnelStreamVisitor,
		tunnelRetErr:        retErr,
		agentId:             agentInfo.Id,
		agentDescriptor:     descriptor.Descriptor_.AgentDescriptor,
		state:               stateReady,
		onForward:           r.onTunnelForward,
		onDone:              r.onTunnelDone,
	}
	// Register
	doIO := func() IOFunc {
		r.mu.Lock()
		defer r.mu.Unlock()
		return r.registerTunnelLocked(tun)
	}()
	doIO()
	// Wait for return error or for cancellation
	select {
	case <-ctx.Done():
		// Context canceled
		r.mu.Lock()
		switch tun.state {
		case stateReady:
			tun.state = stateContextDone
			doIO = r.unregisterTunnelLocked(tun) // nolint: contextcheck
			r.mu.Unlock()
			doIO()
			return nil
		case stateFound:
			// Tunnel was found but hasn't been used yet, Done() hasn't been called.
			// Set state to stateContextDone so that ForwardStream() errors out without doing any I/O.
			tun.state = stateContextDone
			r.mu.Unlock()
			return nil
		case stateForwarding:
			// I/O on the stream will error out, just wait for the return value.
			r.mu.Unlock()
			return <-retErr
		case stateDone:
			// Forwarding has finished and then ctx signaled done. Return the result value from forwarding.
			r.mu.Unlock()
			return <-retErr
		case stateContextDone:
			// Cannot happen twice.
			r.mu.Unlock()
			panic(errors.New("unreachable"))
		case invalid:
			fallthrough // satisfy 'exhaustive' linter
		default:
			// Should never happen
			r.mu.Unlock()
			panic(fmt.Errorf("invalid state: %d", tun.state))
		}
	case err = <-retErr:
		return err
	}
}

func (r *TunnelRegistry) registerTunnelLocked(toReg *tunnel) IOFunc {
	agentId := toReg.agentId
	// 1. Before registering the tunnel see if there is a find tunnel request waiting for it
	findRequestsForAgentId := r.findRequestsByAgentId[agentId]
	for ftr := range findRequestsForAgentId {
		if !toReg.agentDescriptor.SupportsServiceAndMethod(ftr.service, ftr.method) {
			continue
		}
		// Waiting request found!
		toReg.state = stateFound
		ftr.retTun <- toReg      // Satisfy the waiting request ASAP
		r.deleteFindRequest(ftr) // Remove it from the queue
		return noIO
	}

	// 2. Register the tunnel
	toReg.state = stateReady
	r.tuns[toReg] = struct{}{}
	tunsByAgentId := r.tunsByAgentId[agentId]
	if tunsByAgentId == nil {
		tunsByAgentId = make(map[*tunnel]struct{}, 1)
		r.tunsByAgentId[agentId] = tunsByAgentId
	}
	tunsByAgentId[toReg] = struct{}{}
	return r.registerTunnelIO(agentId)
}

func (r *TunnelRegistry) unregisterTunnelLocked(toUnreg *tunnel) IOFunc {
	agentId := toUnreg.agentId
	delete(r.tuns, toUnreg)
	tunsByAgentId := r.tunsByAgentId[agentId]
	delete(tunsByAgentId, toUnreg)
	if len(tunsByAgentId) == 0 {
		delete(r.tunsByAgentId, agentId)
	}
	return r.unregisterTunnelIO(agentId)
}

func (r *TunnelRegistry) onTunnelForward(tun *tunnel) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	switch tun.state {
	case stateReady:
		return status.Error(codes.Internal, "unreachable: ready -> forwarding should never happen")
	case stateFound:
		tun.state = stateForwarding
		return nil
	case stateForwarding:
		return status.Error(codes.Internal, "ForwardStream() called more than once")
	case stateDone:
		return status.Error(codes.Internal, "ForwardStream() called after Done()")
	case stateContextDone:
		return context.Canceled
	case invalid:
		fallthrough // satisfy 'exhaustive' linter
	default:
		return status.Errorf(codes.Internal, "unreachable: invalid state: %d", tun.state)
	}
}

func (r *TunnelRegistry) onTunnelDone(tun *tunnel) {
	var doIO IOFunc
	func() {
		r.mu.Lock()
		defer r.mu.Unlock()
		doIO = r.onTunnelDoneLocked(tun)
	}()
	doIO()
}

func (r *TunnelRegistry) onTunnelDoneLocked(tun *tunnel) IOFunc {
	switch tun.state {
	case stateReady:
		panic(errors.New("unreachable: ready -> done should never happen"))
	case stateFound:
		// Tunnel was found but was not used, Done() was called. Just put it back.
		return r.registerTunnelLocked(tun)
	case stateForwarding:
		tun.state = stateDone
	case stateDone:
		panic(errors.New("Done() called more than once"))
	case stateContextDone:
	// Done() called after cancelled context in HandleTunnel(). Nothing to do.
	case invalid:
		fallthrough // satisfy 'exhaustive' linter
	default:
		// Should never happen
		panic(fmt.Errorf("invalid state: %d", tun.state))
	}
	return noIO
}

func (r *TunnelRegistry) deleteFindRequest(ftr *findTunnelRequest) {
	findRequestsForAgentId := r.findRequestsByAgentId[ftr.agentId]
	delete(findRequestsForAgentId, ftr)
	if len(findRequestsForAgentId) == 0 {
		delete(r.findRequestsByAgentId, ftr.agentId)
	}
}

// Stop aborts any open tunnels.
// It should not be necessary to abort tunnels when registry is used correctly i.e. this method is called after
// all tunnels have terminated gracefully.
func (r *TunnelRegistry) Stop() {
	r.stopInternal()
}

// stopInternal is used for testing.
func (r *TunnelRegistry) stopInternal() (int, int) {
	// Abort all tunnels
	r.mu.Lock()
	defer r.mu.Unlock()
	tl := len(r.tuns)
	fl := len(r.findRequestsByAgentId)
	if tl == 0 && fl == 0 {
		return 0, 0 // Avoid logging a warning
	}
	r.log.Warn("Stopping tunnels and find requests", logz.NumberOfTunnels(len(r.tuns)), logz.NumberOfTunnelFindRequests(len(r.findRequestsByAgentId)))
	for tun := range r.tuns {
		tun.state = stateDone
		tun.tunnelRetErr <- nil // nil so that HandleTunnel() returns cleanly and agent immediately retries
		doIO := r.unregisterTunnelLocked(tun)
		doIO()
	}
	r.log.Warn("Done stopping tunnels")
	// Abort all waiting new stream requests
	for _, findRequestsForAgentId := range r.findRequestsByAgentId {
		for ftr := range findRequestsForAgentId {
			ftr.retTun <- nil // respond ASAP, then do all the bookkeeping
			r.deleteFindRequest(ftr)
		}
	}
	return tl, fl
}

func (r *TunnelRegistry) registerTunnelIO(agentId int64) IOFunc {
	return func() {
		err := r.tunnelRegisterer.RegisterTunnel(context.Background(), agentId) // don't pass context to always register
		if err != nil {
			r.errRep.HandleProcessingError(context.Background(), r.log.With(logz.AgentId(agentId)), "Failed to register tunnel", err)
		}
	}
}

func (r *TunnelRegistry) unregisterTunnelIO(agentId int64) IOFunc {
	return func() {
		err := r.tunnelRegisterer.UnregisterTunnel(context.Background(), agentId) // don't pass context to always unregister
		if err != nil {
			r.errRep.HandleProcessingError(context.Background(), r.log.With(logz.AgentId(agentId)), "Failed to unregister tunnel", err)
		}
	}
}

func noIO() {
}
