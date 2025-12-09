package accounting_iface_mock

import (
    "context"
    "github.com/stretchr/testify/mock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/accounting_iface/v1"
)

type MockTransferService struct {
    mock.Mock
}

func (m *MockTransferService) TransferTeam(ctx context.Context, req *connect.Request[v1.TransferTeamRequest]) (*connect.Response[v1.TransferTeamResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.TransferTeamResponse]), args.Error(1)
}

func (m *MockTransferService) TransferList(ctx context.Context, req *connect.Request[v1.TransferListRequest]) (*connect.Response[v1.TransferListResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.TransferListResponse]), args.Error(1)
}

