package warehouse_iface_mock

import (
    "context"
    "github.com/stretchr/testify/mock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/warehouse_iface/v1"
)

type MockTransactionService struct {
    mock.Mock
}

func (m *MockTransactionService) TransactionItemSearch(ctx context.Context, req *connect.Request[v1.TransactionItemSearchRequest]) (*connect.Response[v1.TransactionItemSearchResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.TransactionItemSearchResponse]), args.Error(1)
}

