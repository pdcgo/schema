package common_mock

import (
    "context"
    "github.com/stretchr/testify/mock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/common/v1"
)

type MockCustomerDataService struct {
    mock.Mock
}

func (m *MockCustomerDataService) CustomerIDs(ctx context.Context, req *connect.Request[v1.CustomerIDsRequest]) (*connect.Response[v1.CustomerIDsResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.CustomerIDsResponse]), args.Error(1)
}

