package common_mock

import (
    "context"
    "github.com/stretchr/testify/mock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/common/v1"
)

type MockWarehouseService struct {
    mock.Mock
}

func (m *MockWarehouseService) PublicWarehouseIDs(ctx context.Context, req *connect.Request[v1.PublicWarehouseIDsRequest]) (*connect.Response[v1.PublicWarehouseIDsResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.PublicWarehouseIDsResponse]), args.Error(1)
}

