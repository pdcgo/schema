package stock_iface_mock

import (
    "context"
    "github.com/stretchr/testify/mock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/stock_iface/v1"
)

type MockStockService struct {
    mock.Mock
}

func (m *MockStockService) InboundCreate(ctx context.Context, req *connect.Request[v1.InboundCreateRequest]) (*connect.Response[v1.InboundCreateResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.InboundCreateResponse]), args.Error(1)
}

func (m *MockStockService) InboundUpdate(ctx context.Context, req *connect.Request[v1.InboundUpdateRequest]) (*connect.Response[v1.InboundUpdateResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.InboundUpdateResponse]), args.Error(1)
}

func (m *MockStockService) InboundAccept(ctx context.Context, req *connect.Request[v1.InboundAcceptRequest]) (*connect.Response[v1.InboundAcceptResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.InboundAcceptResponse]), args.Error(1)
}

func (m *MockStockService) StockAdjustment(ctx context.Context, req *connect.Request[v1.StockAdjustmentRequest]) (*connect.Response[v1.StockAdjustmentResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.StockAdjustmentResponse]), args.Error(1)
}

func (m *MockStockService) TransferToWarehouse(ctx context.Context, req *connect.Request[v1.TransferToWarehouseRequest]) (*connect.Response[v1.TransferToWarehouseResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.TransferToWarehouseResponse]), args.Error(1)
}

func (m *MockStockService) TransferToWarehouseAccept(ctx context.Context, req *connect.Request[v1.TransferToWarehouseAcceptRequest]) (*connect.Response[v1.TransferToWarehouseAcceptResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.TransferToWarehouseAcceptResponse]), args.Error(1)
}

func (m *MockStockService) TransferToWarehouseCancel(ctx context.Context, req *connect.Request[v1.TransferToWarehouseCancelRequest]) (*connect.Response[v1.TransferToWarehouseCancelResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.TransferToWarehouseCancelResponse]), args.Error(1)
}

func (m *MockStockService) StockProblemCreate(ctx context.Context, req *connect.Request[v1.StockProblemCreateRequest]) (*connect.Response[v1.StockProblemCreateResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.StockProblemCreateResponse]), args.Error(1)
}

