package withdrawal_iface_mock

import (
    "context"
    "github.com/stretchr/testify/mock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/withdrawal_iface/v1"
)

type MockWithdrawalService struct {
    mock.Mock
}

func (m *MockWithdrawalService) SubmitWithdrawal(ctx context.Context, req *connect.Request[v1.SubmitWithdrawalRequest]) (*connect.Response[v1.SubmitWithdrawalResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.SubmitWithdrawalResponse]), args.Error(1)
}

func (m *MockWithdrawalService) GetTaskList(ctx context.Context, req *connect.Request[v1.GetTaskListRequest]) (*connect.Response[v1.GetTaskListResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.GetTaskListResponse]), args.Error(1)
}

func (m *MockWithdrawalService) Run(ctx context.Context, req *connect.Request[v1.RunRequest]) (*connect.Response[v1.RunResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.RunResponse]), args.Error(1)
}

func (m *MockWithdrawalService) Stop(ctx context.Context, req *connect.Request[v1.StopRequest]) (*connect.Response[v1.StopResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.StopResponse]), args.Error(1)
}

func (m *MockWithdrawalService) HealthCheck(ctx context.Context, req *connect.Request[v1.HealthCheckRequest]) (*connect.Response[v1.HealthCheckResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.HealthCheckResponse]), args.Error(1)
}

