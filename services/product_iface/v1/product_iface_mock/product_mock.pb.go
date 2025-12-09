package product_iface_mock

import (
    "context"
    "github.com/stretchr/testify/mock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/product_iface/v1"
)

type MockProductService struct {
    mock.Mock
}

func (m *MockProductService) ProductDuplicate(ctx context.Context, req *connect.Request[v1.ProductDuplicateRequest]) (*connect.Response[v1.ProductDuplicateResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.ProductDuplicateResponse]), args.Error(1)
}

func (m *MockProductService) ProductMapGet(ctx context.Context, req *connect.Request[v1.ProductMapGetRequest]) (*connect.Response[v1.ProductMapGetResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.ProductMapGetResponse]), args.Error(1)
}

func (m *MockProductService) ProductMapConnect(ctx context.Context, req *connect.Request[v1.ProductMapConnectRequest]) (*connect.Response[v1.ProductMapConnectResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.ProductMapConnectResponse]), args.Error(1)
}

