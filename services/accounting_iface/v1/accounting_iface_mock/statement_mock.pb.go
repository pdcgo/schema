package accounting_iface_mock

import (
    "context"
    "github.com/stretchr/testify/mock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/accounting_iface/v1"
)

type MockStatementService struct {
    mock.Mock
}

func (m *MockStatementService) StatementBalance(ctx context.Context, req *connect.Request[v1.StatementBalanceRequest]) (*connect.Response[v1.StatementBalanceResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.StatementBalanceResponse]), args.Error(1)
}

func (m *MockStatementService) StatementIncome(ctx context.Context, req *connect.Request[v1.StatementIncomeRequest]) (*connect.Response[v1.StatementIncomeResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.StatementIncomeResponse]), args.Error(1)
}

func (m *MockStatementService) StatementCashFlow(ctx context.Context, req *connect.Request[v1.StatementCashFlowRequest]) (*connect.Response[v1.StatementCashFlowResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.StatementCashFlowResponse]), args.Error(1)
}

