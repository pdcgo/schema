package accounting_iface_mock

import (
    "context"
    "github.com/stretchr/testify/mock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/accounting_iface/v1"
)

type MockCoreService struct {
    mock.Mock
}

func (m *MockCoreService) AccountKeyList(ctx context.Context, req *connect.Request[v1.AccountKeyListRequest]) (*connect.Response[v1.AccountKeyListResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.AccountKeyListResponse]), args.Error(1)
}

func (m *MockCoreService) TypeLabelList(ctx context.Context, req *connect.Request[v1.TypeLabelListRequest]) (*connect.Response[v1.TypeLabelListResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.TypeLabelListResponse]), args.Error(1)
}

