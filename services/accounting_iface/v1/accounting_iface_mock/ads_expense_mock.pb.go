package accounting_iface_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/accounting_iface/v1"
)

type MockAdsExpenseService struct {
    ctrl     *gomock.Controller
    recorder *MockAdsExpenseServiceMockRecorder
}

type MockAdsExpenseServiceMockRecorder struct {
    mock *MockAdsExpenseService
}

func NewMockAdsExpenseService(ctrl *gomock.Controller) *MockAdsExpenseService {
    mock := &MockAdsExpenseService{ctrl: ctrl}
    mock.recorder = &MockAdsExpenseServiceMockRecorder{mock}
    return mock
}

func (m *MockAdsExpenseService) EXPECT() *MockAdsExpenseServiceMockRecorder {
    return m.recorder
}

func (m *MockAdsExpenseService) AdsExCreate(ctx context.Context, req *connect.Request[v1.AdsExCreateRequest]) (*connect.Response[v1.AdsExCreateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "AdsExCreate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.AdsExCreateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockAdsExpenseServiceMockRecorder) AdsExCreate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdsExCreate", reflect.TypeOf((*MockAdsExpenseService)(nil).AdsExCreate), ctx, req)
}

func (m *MockAdsExpenseService) AdsExList(ctx context.Context, req *connect.Request[v1.AdsExListRequest]) (*connect.Response[v1.AdsExListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "AdsExList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.AdsExListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockAdsExpenseServiceMockRecorder) AdsExList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdsExList", reflect.TypeOf((*MockAdsExpenseService)(nil).AdsExList), ctx, req)
}

func (m *MockAdsExpenseService) AdsExEdit(ctx context.Context, req *connect.Request[v1.AdsExEditRequest]) (*connect.Response[v1.AdsExEditResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "AdsExEdit", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.AdsExEditResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockAdsExpenseServiceMockRecorder) AdsExEdit(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdsExEdit", reflect.TypeOf((*MockAdsExpenseService)(nil).AdsExEdit), ctx, req)
}

func (m *MockAdsExpenseService) AdsExShopMetric(ctx context.Context, req *connect.Request[v1.AdsExShopMetricRequest]) (*connect.Response[v1.AdsExShopMetricResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "AdsExShopMetric", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.AdsExShopMetricResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockAdsExpenseServiceMockRecorder) AdsExShopMetric(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdsExShopMetric", reflect.TypeOf((*MockAdsExpenseService)(nil).AdsExShopMetric), ctx, req)
}

func (m *MockAdsExpenseService) AdsExOverviewMetric(ctx context.Context, req *connect.Request[v1.AdsExOverviewMetricRequest]) (*connect.Response[v1.AdsExOverviewMetricResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "AdsExOverviewMetric", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.AdsExOverviewMetricResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockAdsExpenseServiceMockRecorder) AdsExOverviewMetric(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdsExOverviewMetric", reflect.TypeOf((*MockAdsExpenseService)(nil).AdsExOverviewMetric), ctx, req)
}

func (m *MockAdsExpenseService) AdsExTimeMetric(ctx context.Context, req *connect.Request[v1.AdsExTimeMetricRequest]) (*connect.Response[v1.AdsExTimeMetricResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "AdsExTimeMetric", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.AdsExTimeMetricResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockAdsExpenseServiceMockRecorder) AdsExTimeMetric(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdsExTimeMetric", reflect.TypeOf((*MockAdsExpenseService)(nil).AdsExTimeMetric), ctx, req)
}

