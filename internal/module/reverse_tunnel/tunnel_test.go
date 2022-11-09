package reverse_tunnel

import (
	"context"
	"errors"
	"io"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v15/internal/module/reverse_tunnel/rpc"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v15/internal/tool/grpctool"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v15/internal/tool/testing/matcher"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v15/internal/tool/testing/mock_reverse_tunnel_rpc"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v15/internal/tool/testing/mock_rpc"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v15/internal/tool/testing/testhelpers"
	statuspb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	_ Tunnel = &tunnel{}
)

func TestTunnel_ForwardStream_VisitorErrorIsReturnedOnErrorMessageAndReadError(t *testing.T) {
	ctrl := gomock.NewController(t)
	tunnelRetErr := make(chan error)
	tunnelStreamVisitor, err := grpctool.NewStreamVisitor(&rpc.ConnectRequest{})
	require.NoError(t, err)
	connectServer := mock_reverse_tunnel_rpc.NewMockReverseTunnel_ConnectServer(ctrl)
	incomingStream := mock_rpc.NewMockServerStream(ctrl)
	sts := mock_rpc.NewMockServerTransportStream(ctrl)
	incomingCtx := grpc.NewContextWithServerTransportStream(context.Background(), sts)
	cb := NewMockTunnelDataCallback(ctrl)
	incomingStream.EXPECT().
		Context().
		Return(incomingCtx).
		MinTimes(1)
	sts.EXPECT().
		Method().
		Return("some method")

	gomock.InOrder(
		connectServer.EXPECT().
			Send(gomock.Any()),
		incomingStream.EXPECT().
			RecvMsg(gomock.Any()).
			DoAndReturn(func(x interface{}) error {
				<-tunnelRetErr // wait until the other goroutine finished
				return errors.New("failed read")
			}),
	)

	stat := &statuspb.Status{
		Code:    int32(codes.DataLoss),
		Message: "expected data loss",
	}
	gomock.InOrder(
		connectServer.EXPECT().
			RecvMsg(gomock.Any()).
			Do(testhelpers.RecvMsg(&rpc.ConnectRequest{
				Msg: &rpc.ConnectRequest_Error{
					Error: &rpc.Error{
						Status: stat,
					},
				},
			})),
		cb.EXPECT().
			Error(matcher.ProtoEq(t, stat)),
		connectServer.EXPECT().
			RecvMsg(gomock.Any()).
			Return(errors.New("correct error")),
	)
	c := tunnel{
		tunnel:              connectServer,
		tunnelStreamVisitor: tunnelStreamVisitor,
		tunnelRetErr:        tunnelRetErr,
		onForward: func(t *tunnel) error {
			return nil
		},
		onDone: func(t *tunnel) {},
	}
	err = c.ForwardStream(nil, nil, incomingStream, cb)
	assert.EqualError(t, err, "correct error")
}

func TestTunnel_ForwardStream_IsUnblockedWhenIncomingStreamContextIsCancelledAfterSendingAllData(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	tunnelRetErr := make(chan error, 1)
	tunnelStreamVisitor, err := grpctool.NewStreamVisitor(&rpc.ConnectRequest{})
	require.NoError(t, err)
	connectServer := mock_reverse_tunnel_rpc.NewMockReverseTunnel_ConnectServer(ctrl)
	incomingStream := mock_rpc.NewMockServerStream(ctrl)
	sts := mock_rpc.NewMockServerTransportStream(ctrl)
	incomingCtx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	incomingCtx = grpc.NewContextWithServerTransportStream(incomingCtx, sts)
	cb := NewMockTunnelDataCallback(ctrl)
	incomingStream.EXPECT().
		Context().
		Return(incomingCtx).
		MinTimes(1)
	sts.EXPECT().
		Method().
		Return("some method")

	gomock.InOrder( //
		connectServer.EXPECT().
			Send(gomock.Any()), // ConnectResponse_RequestInfo
		incomingStream.EXPECT().
			RecvMsg(gomock.Any()).
			Return(io.EOF),
		connectServer.EXPECT().
			Send(gomock.Any()), // ConnectResponse_CloseSend
	)

	recvChan := make(chan struct{})
	connectServer.EXPECT().
		RecvMsg(gomock.Any()).
		DoAndReturn(func(msg interface{}) error {
			<-recvChan
			return status.Error(codes.DataLoss, "boom")
		})

	c := tunnel{
		tunnel:              connectServer,
		tunnelStreamVisitor: tunnelStreamVisitor,
		tunnelRetErr:        tunnelRetErr,
		onForward: func(t *tunnel) error {
			return nil
		},
		onDone: func(t *tunnel) {},
	}
	err = c.ForwardStream(nil, nil, incomingStream, cb)
	assert.EqualError(t, err, "rpc error: code = DeadlineExceeded desc = context deadline exceeded")
	err = <-tunnelRetErr
	assert.EqualError(t, err, "rpc error: code = DeadlineExceeded desc = context deadline exceeded")
	close(recvChan) // unblock recv
}
