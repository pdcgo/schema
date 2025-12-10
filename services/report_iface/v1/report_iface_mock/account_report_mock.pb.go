package report_iface_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/report_iface/v1"
)

type MockAccountReportService struct {
    ctrl     *gomock.Controller
    recorder *MockAccountReportServiceMockRecorder
}

type MockAccountReportServiceMockRecorder struct {
    mock *MockAccountReportService
}

func NewMockAccountReportService(ctrl *gomock.Controller) *MockAccountReportService {
    mock := &MockAccountReportService{ctrl: ctrl}
    mock.recorder = &MockAccountReportServiceMockRecorder{mock}
    return mock
}

func (m *MockAccountReportService) EXPECT() *MockAccountReportServiceMockRecorder {
    return m.recorder
}

type MockAccountReportServiceClient struct {
    ctrl     *gomock.Controller
    recorder *MockAccountReportServiceClientMockRecorder
}

type MockAccountReportServiceClientMockRecorder struct {
    mock *MockAccountReportServiceClient
}

func NewMockAccountReportServiceClient(ctrl *gomock.Controller) *MockAccountReportServiceClient {
    mock := &MockAccountReportServiceClient{ctrl: ctrl}
    mock.recorder = &MockAccountReportServiceClientMockRecorder{mock}
    return mock
}

func (m *MockAccountReportServiceClient) EXPECT() *MockAccountReportServiceClientMockRecorder {
    return m.recorder
}

func (m *MockAccountReportService) Balance(ctx context.Context, req *connect.Request[v1.BalanceRequest]) (*connect.Response[v1.BalanceResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "Balance", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.BalanceResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockAccountReportServiceMockRecorder) Balance(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Balance", reflect.TypeOf((*MockAccountReportService)(nil).Balance), ctx, req)
}

func (m *MockAccountReportServiceClient) Balance(ctx context.Context, req *connect.Request[v1.BalanceRequest]) (*connect.Response[v1.BalanceResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "Balance", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.BalanceResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockAccountReportServiceClientMockRecorder) Balance(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Balance", reflect.TypeOf((*MockAccountReportService)(nil).Balance), ctx, req)
}

func (m *MockAccountReportService) DailyBalance(ctx context.Context, req *connect.Request[v1.DailyBalanceRequest]) (*connect.Response[v1.DailyBalanceResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "DailyBalance", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.DailyBalanceResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockAccountReportServiceMockRecorder) DailyBalance(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DailyBalance", reflect.TypeOf((*MockAccountReportService)(nil).DailyBalance), ctx, req)
}

func (m *MockAccountReportServiceClient) DailyBalance(ctx context.Context, req *connect.Request[v1.DailyBalanceRequest]) (*connect.Response[v1.DailyBalanceResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "DailyBalance", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.DailyBalanceResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockAccountReportServiceClientMockRecorder) DailyBalance(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DailyBalance", reflect.TypeOf((*MockAccountReportService)(nil).DailyBalance), ctx, req)
}

func (m *MockAccountReportService) BalanceDetail(ctx context.Context, req *connect.Request[v1.BalanceDetailRequest]) (*connect.Response[v1.BalanceDetailResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "BalanceDetail", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.BalanceDetailResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockAccountReportServiceMockRecorder) BalanceDetail(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BalanceDetail", reflect.TypeOf((*MockAccountReportService)(nil).BalanceDetail), ctx, req)
}

func (m *MockAccountReportServiceClient) BalanceDetail(ctx context.Context, req *connect.Request[v1.BalanceDetailRequest]) (*connect.Response[v1.BalanceDetailResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "BalanceDetail", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.BalanceDetailResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockAccountReportServiceClientMockRecorder) BalanceDetail(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BalanceDetail", reflect.TypeOf((*MockAccountReportService)(nil).BalanceDetail), ctx, req)
}

func (m *MockAccountReportService) DailyBalanceDetail(ctx context.Context, req *connect.Request[v1.DailyBalanceDetailRequest]) (*connect.Response[v1.DailyBalanceDetailResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "DailyBalanceDetail", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.DailyBalanceDetailResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockAccountReportServiceMockRecorder) DailyBalanceDetail(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DailyBalanceDetail", reflect.TypeOf((*MockAccountReportService)(nil).DailyBalanceDetail), ctx, req)
}

func (m *MockAccountReportServiceClient) DailyBalanceDetail(ctx context.Context, req *connect.Request[v1.DailyBalanceDetailRequest]) (*connect.Response[v1.DailyBalanceDetailResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "DailyBalanceDetail", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.DailyBalanceDetailResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockAccountReportServiceClientMockRecorder) DailyBalanceDetail(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DailyBalanceDetail", reflect.TypeOf((*MockAccountReportService)(nil).DailyBalanceDetail), ctx, req)
}

func (m *MockAccountReportService) MonthlyBalance(ctx context.Context, req *connect.Request[v1.MonthlyBalanceRequest]) (*connect.Response[v1.MonthlyBalanceResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "MonthlyBalance", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.MonthlyBalanceResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockAccountReportServiceMockRecorder) MonthlyBalance(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MonthlyBalance", reflect.TypeOf((*MockAccountReportService)(nil).MonthlyBalance), ctx, req)
}

func (m *MockAccountReportServiceClient) MonthlyBalance(ctx context.Context, req *connect.Request[v1.MonthlyBalanceRequest]) (*connect.Response[v1.MonthlyBalanceResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "MonthlyBalance", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.MonthlyBalanceResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockAccountReportServiceClientMockRecorder) MonthlyBalance(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MonthlyBalance", reflect.TypeOf((*MockAccountReportService)(nil).MonthlyBalance), ctx, req)
}

func (m *MockAccountReportService) MonthlyBalanceDetail(ctx context.Context, req *connect.Request[v1.MonthlyBalanceDetailRequest]) (*connect.Response[v1.MonthlyBalanceDetailResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "MonthlyBalanceDetail", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.MonthlyBalanceDetailResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockAccountReportServiceMockRecorder) MonthlyBalanceDetail(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MonthlyBalanceDetail", reflect.TypeOf((*MockAccountReportService)(nil).MonthlyBalanceDetail), ctx, req)
}

func (m *MockAccountReportServiceClient) MonthlyBalanceDetail(ctx context.Context, req *connect.Request[v1.MonthlyBalanceDetailRequest]) (*connect.Response[v1.MonthlyBalanceDetailResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "MonthlyBalanceDetail", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.MonthlyBalanceDetailResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockAccountReportServiceClientMockRecorder) MonthlyBalanceDetail(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MonthlyBalanceDetail", reflect.TypeOf((*MockAccountReportService)(nil).MonthlyBalanceDetail), ctx, req)
}

func (m *MockAccountReportService) DailyUpdateBalance(ctx context.Context, req *connect.Request[v1.DailyUpdateBalanceRequest]) (*connect.Response[v1.DailyUpdateBalanceResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "DailyUpdateBalance", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.DailyUpdateBalanceResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockAccountReportServiceMockRecorder) DailyUpdateBalance(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DailyUpdateBalance", reflect.TypeOf((*MockAccountReportService)(nil).DailyUpdateBalance), ctx, req)
}

func (m *MockAccountReportServiceClient) DailyUpdateBalance(ctx context.Context, req *connect.Request[v1.DailyUpdateBalanceRequest]) (*connect.Response[v1.DailyUpdateBalanceResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "DailyUpdateBalance", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.DailyUpdateBalanceResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockAccountReportServiceClientMockRecorder) DailyUpdateBalance(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DailyUpdateBalance", reflect.TypeOf((*MockAccountReportService)(nil).DailyUpdateBalance), ctx, req)
}

func (m *MockAccountReportService) DailyUpdateBalanceAsync(ctx context.Context, req *connect.Request[v1.DailyUpdateBalanceAsyncRequest]) (*connect.Response[v1.DailyUpdateBalanceAsyncResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "DailyUpdateBalanceAsync", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.DailyUpdateBalanceAsyncResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockAccountReportServiceMockRecorder) DailyUpdateBalanceAsync(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DailyUpdateBalanceAsync", reflect.TypeOf((*MockAccountReportService)(nil).DailyUpdateBalanceAsync), ctx, req)
}

func (m *MockAccountReportServiceClient) DailyUpdateBalanceAsync(ctx context.Context, req *connect.Request[v1.DailyUpdateBalanceAsyncRequest]) (*connect.Response[v1.DailyUpdateBalanceAsyncResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "DailyUpdateBalanceAsync", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.DailyUpdateBalanceAsyncResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockAccountReportServiceClientMockRecorder) DailyUpdateBalanceAsync(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DailyUpdateBalanceAsync", reflect.TypeOf((*MockAccountReportService)(nil).DailyUpdateBalanceAsync), ctx, req)
}

