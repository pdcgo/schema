package revenue_iface_mock

import (
    "context"
    "github.com/stretchr/testify/mock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/revenue_iface/v1"
)

type MockRevenueService struct {
    mock.Mock
}

func (m *MockRevenueService) OnOrder(ctx context.Context, req *connect.Request[v1.OnOrderRequest]) (*connect.Response[v1.OnOrderResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.OnOrderResponse]), args.Error(1)
}

func (m *MockRevenueService) OrderCancel(ctx context.Context, req *connect.Request[v1.OrderCancelRequest]) (*connect.Response[v1.OrderCancelResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.OrderCancelResponse]), args.Error(1)
}

func (m *MockRevenueService) OrderReturnAsync(ctx context.Context, req *connect.Request[v1.OrderReturnAsyncRequest]) (*connect.Response[v1.OrderReturnAsyncResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.OrderReturnAsyncResponse]), args.Error(1)
}

func (m *MockRevenueService) OrderReturn(ctx context.Context, req *connect.Request[v1.OrderReturnRequest]) (*connect.Response[v1.OrderReturnResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.OrderReturnResponse]), args.Error(1)
}

func (m *MockRevenueService) OrderCompleted(ctx context.Context, req *connect.Request[v1.OrderCompletedRequest]) (*connect.Response[v1.OrderCompletedResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.OrderCompletedResponse]), args.Error(1)
}

func (m *MockRevenueService) RevenueAdjustment(ctx context.Context, req *connect.Request[v1.RevenueAdjustmentRequest]) (*connect.Response[v1.RevenueAdjustmentResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.RevenueAdjustmentResponse]), args.Error(1)
}

func (m *MockRevenueService) Withdrawal(ctx context.Context, req *connect.Request[v1.WithdrawalRequest]) (*connect.Response[v1.WithdrawalResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.WithdrawalResponse]), args.Error(1)
}

func (m *MockRevenueService) RevenueStream(ctx context.Context, req *connect.Request[v1.RevenueStreamRequest]) (*connect.Response[v1.RevenueStreamResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.RevenueStreamResponse]), args.Error(1)
}

