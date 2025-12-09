package common_mock

import (
    "context"
    "github.com/stretchr/testify/mock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/common/v1"
)

type MockShopService struct {
    mock.Mock
}

func (m *MockShopService) PublicShopIDs(ctx context.Context, req *connect.Request[v1.PublicShopIDsRequest]) (*connect.Response[v1.PublicShopIDsResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.PublicShopIDsResponse]), args.Error(1)
}

func (m *MockShopService) PublicShopList(ctx context.Context, req *connect.Request[v1.PublicShopListRequest]) (*connect.Response[v1.PublicShopListResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.PublicShopListResponse]), args.Error(1)
}

