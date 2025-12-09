package accounting_iface_mock

import (
    "context"
    "github.com/stretchr/testify/mock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/accounting_iface/v1"
)

type MockAdsExpenseService struct {
    mock.Mock
}

func (m *MockAdsExpenseService) AdsExCreate(ctx context.Context, req *connect.Request[v1.AdsExCreateRequest]) (*connect.Response[v1.AdsExCreateResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.AdsExCreateResponse]), args.Error(1)
}

func (m *MockAdsExpenseService) AdsExList(ctx context.Context, req *connect.Request[v1.AdsExListRequest]) (*connect.Response[v1.AdsExListResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.AdsExListResponse]), args.Error(1)
}

func (m *MockAdsExpenseService) AdsExEdit(ctx context.Context, req *connect.Request[v1.AdsExEditRequest]) (*connect.Response[v1.AdsExEditResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.AdsExEditResponse]), args.Error(1)
}

func (m *MockAdsExpenseService) AdsExShopMetric(ctx context.Context, req *connect.Request[v1.AdsExShopMetricRequest]) (*connect.Response[v1.AdsExShopMetricResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.AdsExShopMetricResponse]), args.Error(1)
}

func (m *MockAdsExpenseService) AdsExOverviewMetric(ctx context.Context, req *connect.Request[v1.AdsExOverviewMetricRequest]) (*connect.Response[v1.AdsExOverviewMetricResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.AdsExOverviewMetricResponse]), args.Error(1)
}

func (m *MockAdsExpenseService) AdsExTimeMetric(ctx context.Context, req *connect.Request[v1.AdsExTimeMetricRequest]) (*connect.Response[v1.AdsExTimeMetricResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.AdsExTimeMetricResponse]), args.Error(1)
}

