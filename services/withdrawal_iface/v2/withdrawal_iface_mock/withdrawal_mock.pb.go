package withdrawal_iface_mock

import (
    "context"
    "github.com/stretchr/testify/mock"
    connect "connectrpc.com/connect"
    v2 "github.com/pdcgo/schema/services/withdrawal_iface/v2"
)

type MockWithdrawalService struct {
    mock.Mock
}

func (m *MockWithdrawalService) SubmitWithdrawal(ctx context.Context, req *connect.Request[v2.SubmitWithdrawalRequest]) (*connect.Response[v2.SubmitWithdrawalResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v2.SubmitWithdrawalResponse]), args.Error(1)
}

func (m *MockWithdrawalService) SubmitWithdrawalTiktok(ctx context.Context, req *connect.Request[v2.SubmitWithdrawalTiktokRequest]) (*connect.Response[v2.SubmitWithdrawalTiktokResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v2.SubmitWithdrawalTiktokResponse]), args.Error(1)
}

func (m *MockWithdrawalService) SubmitWithdrawalShopee(ctx context.Context, req *connect.Request[v2.SubmitWithdrawalShopeeRequest]) (*connect.Response[v2.SubmitWithdrawalShopeeResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v2.SubmitWithdrawalShopeeResponse]), args.Error(1)
}

