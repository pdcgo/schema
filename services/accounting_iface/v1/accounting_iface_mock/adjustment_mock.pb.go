package accounting_iface_mock

import (
    "context"
    "github.com/stretchr/testify/mock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/accounting_iface/v1"
)

type MockAdjustmentService struct {
    mock.Mock
}

func (m *MockAdjustmentService) AccountAdjustment(ctx context.Context, req *connect.Request[v1.AccountAdjustmentRequest]) (*connect.Response[v1.AccountAdjustmentResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.AccountAdjustmentResponse]), args.Error(1)
}

func (m *MockAdjustmentService) AdjCreate(ctx context.Context, req *connect.Request[v1.AdjCreateRequest]) (*connect.Response[v1.AdjCreateResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.AdjCreateResponse]), args.Error(1)
}

