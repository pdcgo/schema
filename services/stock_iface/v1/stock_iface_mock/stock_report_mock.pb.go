package stock_iface_mock

import (
    "context"
    "github.com/stretchr/testify/mock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/stock_iface/v1"
)

type MockStockReportService struct {
    mock.Mock
}

func (m *MockStockReportService) StockTeam(ctx context.Context, req *connect.Request[v1.StockTeamRequest]) (*connect.Response[v1.StockTeamResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.StockTeamResponse]), args.Error(1)
}

