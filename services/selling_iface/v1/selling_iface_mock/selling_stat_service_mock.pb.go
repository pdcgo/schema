package selling_iface_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/selling_iface/v1"
)

type MockSellingStatService struct {
    ctrl     *gomock.Controller
    recorder *MockSellingStatServiceMockRecorder
}

type MockSellingStatServiceMockRecorder struct {
    mock *MockSellingStatService
}

func NewMockSellingStatService(ctrl *gomock.Controller) *MockSellingStatService {
    mock := &MockSellingStatService{ctrl: ctrl}
    mock.recorder = &MockSellingStatServiceMockRecorder{mock}
    return mock
}

func (m *MockSellingStatService) EXPECT() *MockSellingStatServiceMockRecorder {
    return m.recorder
}

type MockSellingStatServiceClient struct {
    ctrl     *gomock.Controller
    recorder *MockSellingStatServiceClientMockRecorder
}

type MockSellingStatServiceClientMockRecorder struct {
    mock *MockSellingStatServiceClient
}

func NewMockSellingStatServiceClient(ctrl *gomock.Controller) *MockSellingStatServiceClient {
    mock := &MockSellingStatServiceClient{ctrl: ctrl}
    mock.recorder = &MockSellingStatServiceClientMockRecorder{mock}
    return mock
}

func (m *MockSellingStatServiceClient) EXPECT() *MockSellingStatServiceClientMockRecorder {
    return m.recorder
}

func (m *MockSellingStatService) Stat(ctx context.Context, req *connect.Request[v1.StatRequest]) (*connect.Response[v1.StatResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "Stat", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.StatResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockSellingStatServiceMockRecorder) Stat(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stat", reflect.TypeOf((*MockSellingStatService)(nil).Stat), ctx, req)
}

func (m *MockSellingStatServiceClient) Stat(ctx context.Context, req *connect.Request[v1.StatRequest]) (*connect.Response[v1.StatResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "Stat", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.StatResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockSellingStatServiceClientMockRecorder) Stat(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stat", reflect.TypeOf((*MockSellingStatService)(nil).Stat), ctx, req)
}

func (m *MockSellingStatService) ProductStatMetric(ctx context.Context, req *connect.Request[v1.ProductStatMetricRequest]) (*connect.Response[v1.ProductStatMetricResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ProductStatMetric", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ProductStatMetricResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockSellingStatServiceMockRecorder) ProductStatMetric(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductStatMetric", reflect.TypeOf((*MockSellingStatService)(nil).ProductStatMetric), ctx, req)
}

func (m *MockSellingStatServiceClient) ProductStatMetric(ctx context.Context, req *connect.Request[v1.ProductStatMetricRequest]) (*connect.Response[v1.ProductStatMetricResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ProductStatMetric", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ProductStatMetricResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockSellingStatServiceClientMockRecorder) ProductStatMetric(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductStatMetric", reflect.TypeOf((*MockSellingStatService)(nil).ProductStatMetric), ctx, req)
}

func (m *MockSellingStatService) ShopStatMetric(ctx context.Context, req *connect.Request[v1.ShopStatMetricRequest]) (*connect.Response[v1.ShopStatMetricResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ShopStatMetric", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ShopStatMetricResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockSellingStatServiceMockRecorder) ShopStatMetric(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShopStatMetric", reflect.TypeOf((*MockSellingStatService)(nil).ShopStatMetric), ctx, req)
}

func (m *MockSellingStatServiceClient) ShopStatMetric(ctx context.Context, req *connect.Request[v1.ShopStatMetricRequest]) (*connect.Response[v1.ShopStatMetricResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ShopStatMetric", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ShopStatMetricResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockSellingStatServiceClientMockRecorder) ShopStatMetric(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShopStatMetric", reflect.TypeOf((*MockSellingStatService)(nil).ShopStatMetric), ctx, req)
}

func (m *MockSellingStatService) UserStatMetric(ctx context.Context, req *connect.Request[v1.UserStatMetricRequest]) (*connect.Response[v1.UserStatMetricResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "UserStatMetric", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.UserStatMetricResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockSellingStatServiceMockRecorder) UserStatMetric(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserStatMetric", reflect.TypeOf((*MockSellingStatService)(nil).UserStatMetric), ctx, req)
}

func (m *MockSellingStatServiceClient) UserStatMetric(ctx context.Context, req *connect.Request[v1.UserStatMetricRequest]) (*connect.Response[v1.UserStatMetricResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "UserStatMetric", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.UserStatMetricResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockSellingStatServiceClientMockRecorder) UserStatMetric(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserStatMetric", reflect.TypeOf((*MockSellingStatService)(nil).UserStatMetric), ctx, req)
}

func (m *MockSellingStatService) CrossProductList(ctx context.Context, req *connect.Request[v1.CrossProductListRequest]) (*connect.Response[v1.CrossProductListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "CrossProductList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.CrossProductListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockSellingStatServiceMockRecorder) CrossProductList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CrossProductList", reflect.TypeOf((*MockSellingStatService)(nil).CrossProductList), ctx, req)
}

func (m *MockSellingStatServiceClient) CrossProductList(ctx context.Context, req *connect.Request[v1.CrossProductListRequest]) (*connect.Response[v1.CrossProductListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "CrossProductList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.CrossProductListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockSellingStatServiceClientMockRecorder) CrossProductList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CrossProductList", reflect.TypeOf((*MockSellingStatService)(nil).CrossProductList), ctx, req)
}

