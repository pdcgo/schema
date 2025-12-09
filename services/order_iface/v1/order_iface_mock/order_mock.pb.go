package order_iface_mock

import (
    "context"
    "github.com/stretchr/testify/mock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/order_iface/v1"
)

type MockOrderService struct {
    mock.Mock
}

func (m *MockOrderService) OrderFundSet(ctx context.Context, req *connect.Request[v1.OrderFundSetRequest]) (*connect.Response[v1.OrderFundSetResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.OrderFundSetResponse]), args.Error(1)
}

func (m *MockOrderService) OrderTagRemove(ctx context.Context, req *connect.Request[v1.OrderTagRemoveRequest]) (*connect.Response[v1.OrderTagRemoveResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.OrderTagRemoveResponse]), args.Error(1)
}

func (m *MockOrderService) OrderTagAdd(ctx context.Context, req *connect.Request[v1.OrderTagAddRequest]) (*connect.Response[v1.OrderTagAddResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.OrderTagAddResponse]), args.Error(1)
}

func (m *MockOrderService) OrderList(ctx context.Context, req *connect.Request[v1.OrderListRequest]) (*connect.Response[v1.OrderListResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.OrderListResponse]), args.Error(1)
}

func (m *MockOrderService) OrderOverview(ctx context.Context, req *connect.Request[v1.OrderOverviewRequest]) (*connect.Response[v1.OrderOverviewResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.OrderOverviewResponse]), args.Error(1)
}

