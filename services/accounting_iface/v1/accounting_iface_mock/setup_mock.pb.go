package accounting_iface_mock

import (
    "context"
    "github.com/stretchr/testify/mock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/accounting_iface/v1"
)

type MockAccountingSetupService struct {
    mock.Mock
}

func (m *MockAccountingSetupService) Setup(ctx context.Context, req *connect.Request[v1.SetupRequest]) (*connect.Response[v1.SetupResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.SetupResponse]), args.Error(1)
}

func (m *MockAccountingSetupService) RecalculateDaily(ctx context.Context, req *connect.Request[v1.RecalculateDailyRequest]) (*connect.Response[v1.RecalculateDailyResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.RecalculateDailyResponse]), args.Error(1)
}

type MockStreamService struct {
    mock.Mock
}

func (m *MockStreamService) DummyStream(ctx context.Context, req *connect.Request[v1.DummyStreamRequest]) (*connect.Response[v1.DummyStreamResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.DummyStreamResponse]), args.Error(1)
}

