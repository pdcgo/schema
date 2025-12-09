package warehouse_iface_mock

import (
    "context"
    "github.com/stretchr/testify/mock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/warehouse_iface/v1"
)

type MockInboundService struct {
    mock.Mock
}

func (m *MockInboundService) InboundAccept(ctx context.Context, req *connect.Request[v1.InboundAcceptRequest]) (*connect.Response[v1.InboundAcceptResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.InboundAcceptResponse]), args.Error(1)
}

