package accounting_iface_mock

import (
    "context"
    "github.com/stretchr/testify/mock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/accounting_iface/v1"
)

type MockExpenseService struct {
    mock.Mock
}

func (m *MockExpenseService) ExpenseCreate(ctx context.Context, req *connect.Request[v1.ExpenseCreateRequest]) (*connect.Response[v1.ExpenseCreateResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.ExpenseCreateResponse]), args.Error(1)
}

func (m *MockExpenseService) ExpenseList(ctx context.Context, req *connect.Request[v1.ExpenseListRequest]) (*connect.Response[v1.ExpenseListResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.ExpenseListResponse]), args.Error(1)
}

func (m *MockExpenseService) ExpenseTypeList(ctx context.Context, req *connect.Request[v1.ExpenseTypeListRequest]) (*connect.Response[v1.ExpenseTypeListResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.ExpenseTypeListResponse]), args.Error(1)
}

func (m *MockExpenseService) ExpenseOverviewMetric(ctx context.Context, req *connect.Request[v1.ExpenseOverviewMetricRequest]) (*connect.Response[v1.ExpenseOverviewMetricResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.ExpenseOverviewMetricResponse]), args.Error(1)
}

func (m *MockExpenseService) ExpenseTimeMetric(ctx context.Context, req *connect.Request[v1.ExpenseTimeMetricRequest]) (*connect.Response[v1.ExpenseTimeMetricResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.ExpenseTimeMetricResponse]), args.Error(1)
}

