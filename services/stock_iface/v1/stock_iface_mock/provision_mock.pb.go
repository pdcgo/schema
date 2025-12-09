package stock_iface_mock

import (
    "context"
    "github.com/stretchr/testify/mock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/stock_iface/v1"
)

type MockStockProvisionService struct {
    mock.Mock
}

func (m *MockStockProvisionService) StockProvision(ctx context.Context, req *connect.Request[v1.StockProvisionRequest]) (*connect.Response[v1.StockProvisionResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.StockProvisionResponse]), args.Error(1)
}

