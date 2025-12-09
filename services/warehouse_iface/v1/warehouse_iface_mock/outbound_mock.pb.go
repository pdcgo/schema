package warehouse_iface_mock

import (
    "context"
    "github.com/stretchr/testify/mock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/warehouse_iface/v1"
)

type MockOutboundService struct {
    mock.Mock
}

func (m *MockOutboundService) OutboundList(ctx context.Context, req *connect.Request[v1.OutboundListRequest]) (*connect.Response[v1.OutboundListResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.OutboundListResponse]), args.Error(1)
}

