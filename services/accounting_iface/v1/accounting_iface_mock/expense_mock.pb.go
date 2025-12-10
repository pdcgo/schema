package accounting_iface_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/accounting_iface/v1"
)

type MockExpenseService struct {
    ctrl     *gomock.Controller
    recorder *MockExpenseServiceMockRecorder
}

type MockExpenseServiceMockRecorder struct {
    mock *MockExpenseService
}

func NewMockExpenseService(ctrl *gomock.Controller) *MockExpenseService {
    mock := &MockExpenseService{ctrl: ctrl}
    mock.recorder = &MockExpenseServiceMockRecorder{mock}
    return mock
}

func (m *MockExpenseService) EXPECT() *MockExpenseServiceMockRecorder {
    return m.recorder
}

type MockExpenseServiceClient struct {
    ctrl     *gomock.Controller
    recorder *MockExpenseServiceClientMockRecorder
}

type MockExpenseServiceClientMockRecorder struct {
    mock *MockExpenseServiceClient
}

func NewMockExpenseServiceClient(ctrl *gomock.Controller) *MockExpenseServiceClient {
    mock := &MockExpenseServiceClient{ctrl: ctrl}
    mock.recorder = &MockExpenseServiceClientMockRecorder{mock}
    return mock
}

func (m *MockExpenseServiceClient) EXPECT() *MockExpenseServiceClientMockRecorder {
    return m.recorder
}

func (m *MockExpenseService) ExpenseCreate(ctx context.Context, req *connect.Request[v1.ExpenseCreateRequest]) (*connect.Response[v1.ExpenseCreateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ExpenseCreate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ExpenseCreateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockExpenseServiceMockRecorder) ExpenseCreate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExpenseCreate", reflect.TypeOf((*MockExpenseService)(nil).ExpenseCreate), ctx, req)
}

func (m *MockExpenseServiceClient) ExpenseCreate(ctx context.Context, req *connect.Request[v1.ExpenseCreateRequest]) (*connect.Response[v1.ExpenseCreateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ExpenseCreate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ExpenseCreateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockExpenseServiceClientMockRecorder) ExpenseCreate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExpenseCreate", reflect.TypeOf((*MockExpenseService)(nil).ExpenseCreate), ctx, req)
}

func (m *MockExpenseService) ExpenseList(ctx context.Context, req *connect.Request[v1.ExpenseListRequest]) (*connect.Response[v1.ExpenseListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ExpenseList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ExpenseListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockExpenseServiceMockRecorder) ExpenseList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExpenseList", reflect.TypeOf((*MockExpenseService)(nil).ExpenseList), ctx, req)
}

func (m *MockExpenseServiceClient) ExpenseList(ctx context.Context, req *connect.Request[v1.ExpenseListRequest]) (*connect.Response[v1.ExpenseListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ExpenseList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ExpenseListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockExpenseServiceClientMockRecorder) ExpenseList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExpenseList", reflect.TypeOf((*MockExpenseService)(nil).ExpenseList), ctx, req)
}

func (m *MockExpenseService) ExpenseTypeList(ctx context.Context, req *connect.Request[v1.ExpenseTypeListRequest]) (*connect.Response[v1.ExpenseTypeListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ExpenseTypeList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ExpenseTypeListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockExpenseServiceMockRecorder) ExpenseTypeList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExpenseTypeList", reflect.TypeOf((*MockExpenseService)(nil).ExpenseTypeList), ctx, req)
}

func (m *MockExpenseServiceClient) ExpenseTypeList(ctx context.Context, req *connect.Request[v1.ExpenseTypeListRequest]) (*connect.Response[v1.ExpenseTypeListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ExpenseTypeList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ExpenseTypeListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockExpenseServiceClientMockRecorder) ExpenseTypeList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExpenseTypeList", reflect.TypeOf((*MockExpenseService)(nil).ExpenseTypeList), ctx, req)
}

func (m *MockExpenseService) ExpenseOverviewMetric(ctx context.Context, req *connect.Request[v1.ExpenseOverviewMetricRequest]) (*connect.Response[v1.ExpenseOverviewMetricResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ExpenseOverviewMetric", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ExpenseOverviewMetricResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockExpenseServiceMockRecorder) ExpenseOverviewMetric(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExpenseOverviewMetric", reflect.TypeOf((*MockExpenseService)(nil).ExpenseOverviewMetric), ctx, req)
}

func (m *MockExpenseServiceClient) ExpenseOverviewMetric(ctx context.Context, req *connect.Request[v1.ExpenseOverviewMetricRequest]) (*connect.Response[v1.ExpenseOverviewMetricResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ExpenseOverviewMetric", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ExpenseOverviewMetricResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockExpenseServiceClientMockRecorder) ExpenseOverviewMetric(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExpenseOverviewMetric", reflect.TypeOf((*MockExpenseService)(nil).ExpenseOverviewMetric), ctx, req)
}

func (m *MockExpenseService) ExpenseTimeMetric(ctx context.Context, req *connect.Request[v1.ExpenseTimeMetricRequest]) (*connect.Response[v1.ExpenseTimeMetricResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ExpenseTimeMetric", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ExpenseTimeMetricResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockExpenseServiceMockRecorder) ExpenseTimeMetric(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExpenseTimeMetric", reflect.TypeOf((*MockExpenseService)(nil).ExpenseTimeMetric), ctx, req)
}

func (m *MockExpenseServiceClient) ExpenseTimeMetric(ctx context.Context, req *connect.Request[v1.ExpenseTimeMetricRequest]) (*connect.Response[v1.ExpenseTimeMetricResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ExpenseTimeMetric", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ExpenseTimeMetricResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockExpenseServiceClientMockRecorder) ExpenseTimeMetric(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExpenseTimeMetric", reflect.TypeOf((*MockExpenseService)(nil).ExpenseTimeMetric), ctx, req)
}

