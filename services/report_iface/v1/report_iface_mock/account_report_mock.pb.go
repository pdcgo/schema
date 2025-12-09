package report_iface_mock

import (
    "context"
    "github.com/stretchr/testify/mock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/report_iface/v1"
)

type MockAccountReportService struct {
    mock.Mock
}

func (m *MockAccountReportService) Balance(ctx context.Context, req *connect.Request[v1.BalanceRequest]) (*connect.Response[v1.BalanceResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.BalanceResponse]), args.Error(1)
}

func (m *MockAccountReportService) DailyBalance(ctx context.Context, req *connect.Request[v1.DailyBalanceRequest]) (*connect.Response[v1.DailyBalanceResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.DailyBalanceResponse]), args.Error(1)
}

func (m *MockAccountReportService) BalanceDetail(ctx context.Context, req *connect.Request[v1.BalanceDetailRequest]) (*connect.Response[v1.BalanceDetailResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.BalanceDetailResponse]), args.Error(1)
}

func (m *MockAccountReportService) DailyBalanceDetail(ctx context.Context, req *connect.Request[v1.DailyBalanceDetailRequest]) (*connect.Response[v1.DailyBalanceDetailResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.DailyBalanceDetailResponse]), args.Error(1)
}

func (m *MockAccountReportService) MonthlyBalance(ctx context.Context, req *connect.Request[v1.MonthlyBalanceRequest]) (*connect.Response[v1.MonthlyBalanceResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.MonthlyBalanceResponse]), args.Error(1)
}

func (m *MockAccountReportService) MonthlyBalanceDetail(ctx context.Context, req *connect.Request[v1.MonthlyBalanceDetailRequest]) (*connect.Response[v1.MonthlyBalanceDetailResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.MonthlyBalanceDetailResponse]), args.Error(1)
}

func (m *MockAccountReportService) DailyUpdateBalance(ctx context.Context, req *connect.Request[v1.DailyUpdateBalanceRequest]) (*connect.Response[v1.DailyUpdateBalanceResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.DailyUpdateBalanceResponse]), args.Error(1)
}

func (m *MockAccountReportService) DailyUpdateBalanceAsync(ctx context.Context, req *connect.Request[v1.DailyUpdateBalanceAsyncRequest]) (*connect.Response[v1.DailyUpdateBalanceAsyncResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.DailyUpdateBalanceAsyncResponse]), args.Error(1)
}

