package rpc_test

import (
	"context"
	"io"
	"testing"
	"time"

	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/module/gitops/rpc"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/tool/retry"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/tool/testing/matcher"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/tool/testing/mock_rpc"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/tool/testing/testhelpers"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap/zaptest"
)

var (
	_ rpc.ObjectsToSynchronizeWatcherInterface = &rpc.ObjectsToSynchronizeWatcher{}
)

const (
	projectId = "bla123/bla-1"
	revision  = "rev12341234"
)

func TestObjectsToSynchronizeWatcherResumeConnection(t *testing.T) {
	pathsCfg := []*rpc.PathCF{
		{
			Path: &rpc.PathCF_Glob{
				Glob: "*.yaml",
			},
		},
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ctrl := gomock.NewController(t)
	client := mock_rpc.NewMockGitopsClient(ctrl)
	stream1 := mock_rpc.NewMockGitops_GetObjectsToSynchronizeClient(ctrl)
	stream2 := mock_rpc.NewMockGitops_GetObjectsToSynchronizeClient(ctrl)
	req := &rpc.ObjectsToSynchronizeRequest{
		ProjectId: projectId,
		Paths:     pathsCfg,
	}
	gomock.InOrder(
		client.EXPECT().
			GetObjectsToSynchronize(gomock.Any(), matcher.ProtoEq(t, req), gomock.Any()).
			Return(stream1, nil),
		stream1.EXPECT().
			RecvMsg(gomock.Any()).
			Do(testhelpers.RecvMsg(&rpc.ObjectsToSynchronizeResponse{
				Message: &rpc.ObjectsToSynchronizeResponse_Header_{
					Header: &rpc.ObjectsToSynchronizeResponse_Header{
						CommitId: revision,
					},
				},
			})),
		stream1.EXPECT().
			RecvMsg(gomock.Any()).
			Do(testhelpers.RecvMsg(&rpc.ObjectsToSynchronizeResponse{
				Message: &rpc.ObjectsToSynchronizeResponse_Trailer_{
					Trailer: &rpc.ObjectsToSynchronizeResponse_Trailer{},
				},
			})),
		stream1.EXPECT().
			RecvMsg(gomock.Any()).
			Return(io.EOF),
		client.EXPECT().
			GetObjectsToSynchronize(gomock.Any(), matcher.ProtoEq(t, &rpc.ObjectsToSynchronizeRequest{
				ProjectId: projectId,
				CommitId:  revision,
				Paths:     pathsCfg,
			}), gomock.Any()).
			Return(stream2, nil),
		stream2.EXPECT().
			RecvMsg(gomock.Any()).
			DoAndReturn(func(msg interface{}) error {
				cancel()
				return io.EOF
			}),
	)
	w := rpc.ObjectsToSynchronizeWatcher{
		Log:          zaptest.NewLogger(t),
		GitopsClient: client,
		PollConfig:   retry.NewPollConfigFactory(0, retry.NewExponentialBackoffFactory(10*time.Millisecond, time.Minute, time.Minute, 2, 1)),
	}
	w.Watch(ctx, req, func(ctx context.Context, data rpc.ObjectsToSynchronizeData) {
		// Don't care
	})
}

func TestObjectsToSynchronizeWatcherInvalidStream(t *testing.T) {
	tests := []struct {
		name   string
		stream []*rpc.ObjectsToSynchronizeResponse
		eof    bool
	}{
		{
			name: "empty stream", // Not invalid, but should not call callback
			eof:  true,
		},
		{
			name: "missing header",
			stream: []*rpc.ObjectsToSynchronizeResponse{
				{
					Message: &rpc.ObjectsToSynchronizeResponse_Trailer_{
						Trailer: &rpc.ObjectsToSynchronizeResponse_Trailer{},
					},
				},
			},
		},
		{
			name: "unexpected header",
			stream: []*rpc.ObjectsToSynchronizeResponse{
				{
					Message: &rpc.ObjectsToSynchronizeResponse_Header_{
						Header: &rpc.ObjectsToSynchronizeResponse_Header{
							CommitId: revision,
						},
					},
				},
				{
					Message: &rpc.ObjectsToSynchronizeResponse_Header_{
						Header: &rpc.ObjectsToSynchronizeResponse_Header{
							CommitId: revision,
						},
					},
				},
			},
		},
		{
			name: "missing trailer",
			stream: []*rpc.ObjectsToSynchronizeResponse{
				{
					Message: &rpc.ObjectsToSynchronizeResponse_Header_{
						Header: &rpc.ObjectsToSynchronizeResponse_Header{
							CommitId: revision,
						},
					},
				},
			},
			eof: true,
		},
		{
			name: "trailer then header",
			stream: []*rpc.ObjectsToSynchronizeResponse{
				{
					Message: &rpc.ObjectsToSynchronizeResponse_Trailer_{
						Trailer: &rpc.ObjectsToSynchronizeResponse_Trailer{},
					},
				},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			ctrl := gomock.NewController(t)
			client := mock_rpc.NewMockGitopsClient(ctrl)
			stream1 := mock_rpc.NewMockGitops_GetObjectsToSynchronizeClient(ctrl)
			req := &rpc.ObjectsToSynchronizeRequest{
				ProjectId: projectId,
				Paths: []*rpc.PathCF{
					{
						Path: &rpc.PathCF_Glob{
							Glob: "*.yaml",
						},
					},
				},
			}
			calls := []any{
				client.EXPECT().
					GetObjectsToSynchronize(gomock.Any(), matcher.ProtoEq(t, req), gomock.Any()).
					Return(stream1, nil),
			}
			if tc.eof {
				for _, streamItem := range tc.stream {
					calls = append(calls, stream1.EXPECT().
						RecvMsg(gomock.Any()).
						Do(testhelpers.RecvMsg(streamItem)),
					)
				}
				calls = append(calls, stream1.EXPECT().
					RecvMsg(gomock.Any()).
					DoAndReturn(func(msg interface{}) error {
						cancel()
						return io.EOF
					}))
			} else {
				for i := 0; i < len(tc.stream)-1; i++ {
					streamItem := tc.stream[i]
					calls = append(calls, stream1.EXPECT().
						RecvMsg(gomock.Any()).
						Do(testhelpers.RecvMsg(streamItem)),
					)
				}
				calls = append(calls, stream1.EXPECT().RecvMsg(gomock.Any()).Do(func(msg any) error {
					testhelpers.SetValue(msg, tc.stream[len(tc.stream)-1])
					cancel()
					return nil
				}))
			}
			calls = append(calls, stream1.EXPECT().Header().MaxTimes(1)) // exact number not important for this test
			gomock.InOrder(calls...)
			w := rpc.ObjectsToSynchronizeWatcher{
				Log:          zaptest.NewLogger(t),
				GitopsClient: client,
				PollConfig:   retry.NewPollConfigFactory(0, retry.NewExponentialBackoffFactory(10*time.Millisecond, time.Minute, time.Minute, 2, 1)),
			}
			w.Watch(ctx, req, func(ctx context.Context, data rpc.ObjectsToSynchronizeData) {
				// Must not be called
				t.FailNow()
			})
		})
	}
}
