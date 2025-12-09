package common_mock

import (
    "context"
    "github.com/stretchr/testify/mock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/common/v1"
)

type MockSupplierService struct {
    mock.Mock
}

func (m *MockSupplierService) PublicSupplierIDs(ctx context.Context, req *connect.Request[v1.PublicSupplierIDsRequest]) (*connect.Response[v1.PublicSupplierIDsResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.PublicSupplierIDsResponse]), args.Error(1)
}

