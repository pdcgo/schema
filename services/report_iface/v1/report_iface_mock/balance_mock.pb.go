package report_iface_mock

import (
    "context"
    "github.com/stretchr/testify/mock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/report_iface/v1"
)

type MockBalanceService struct {
    mock.Mock
}

func (m *MockBalanceService) BalanceResync(ctx context.Context, req *connect.Request[v1.BalanceResyncRequest]) (*connect.Response[v1.BalanceResyncResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.BalanceResyncResponse]), args.Error(1)
}

