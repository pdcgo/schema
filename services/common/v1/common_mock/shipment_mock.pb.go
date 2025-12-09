package common_mock

import (
    "context"
    "github.com/stretchr/testify/mock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/common/v1"
)

type MockShipmentService struct {
    mock.Mock
}

func (m *MockShipmentService) PublicShipmentIDs(ctx context.Context, req *connect.Request[v1.PublicShipmentIDsRequest]) (*connect.Response[v1.PublicShipmentIDsResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.PublicShipmentIDsResponse]), args.Error(1)
}

