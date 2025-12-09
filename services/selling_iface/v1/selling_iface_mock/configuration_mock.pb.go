package selling_iface_mock

import (
    "context"
    "github.com/stretchr/testify/mock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/selling_iface/v1"
)

type MockConfigurationLimitService struct {
    mock.Mock
}

func (m *MockConfigurationLimitService) LimitInvoice(ctx context.Context, req *connect.Request[v1.LimitInvoiceRequest]) (*connect.Response[v1.LimitInvoiceResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.LimitInvoiceResponse]), args.Error(1)
}

