package withdrawal_iface_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/withdrawal_iface/v1"
)

type MockWithdrawalService struct {
    ctrl     *gomock.Controller
    recorder *MockWithdrawalServiceMockRecorder
}

type MockWithdrawalServiceMockRecorder struct {
    mock *MockWithdrawalService
}

func NewMockWithdrawalService(ctrl *gomock.Controller) *MockWithdrawalService {
    mock := &MockWithdrawalService{ctrl: ctrl}
    mock.recorder = &MockWithdrawalServiceMockRecorder{mock}
    return mock
}

func (m *MockWithdrawalService) EXPECT() *MockWithdrawalServiceMockRecorder {
    return m.recorder
}

type MockWithdrawalServiceClient struct {
    ctrl     *gomock.Controller
    recorder *MockWithdrawalServiceClientMockRecorder
}

type MockWithdrawalServiceClientMockRecorder struct {
    mock *MockWithdrawalServiceClient
}

func NewMockWithdrawalServiceClient(ctrl *gomock.Controller) *MockWithdrawalServiceClient {
    mock := &MockWithdrawalServiceClient{ctrl: ctrl}
    mock.recorder = &MockWithdrawalServiceClientMockRecorder{mock}
    return mock
}

func (m *MockWithdrawalServiceClient) EXPECT() *MockWithdrawalServiceClientMockRecorder {
    return m.recorder
}

func (m *MockWithdrawalService) SubmitWithdrawal(ctx context.Context, req *connect.Request[v1.SubmitWithdrawalRequest]) (*connect.Response[v1.SubmitWithdrawalResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "SubmitWithdrawal", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.SubmitWithdrawalResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockWithdrawalServiceMockRecorder) SubmitWithdrawal(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitWithdrawal", reflect.TypeOf((*MockWithdrawalService)(nil).SubmitWithdrawal), ctx, req)
}

func (m *MockWithdrawalServiceClient) SubmitWithdrawal(ctx context.Context, req *connect.Request[v1.SubmitWithdrawalRequest]) (*connect.Response[v1.SubmitWithdrawalResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "SubmitWithdrawal", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.SubmitWithdrawalResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockWithdrawalServiceClientMockRecorder) SubmitWithdrawal(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitWithdrawal", reflect.TypeOf((*MockWithdrawalService)(nil).SubmitWithdrawal), ctx, req)
}

func (m *MockWithdrawalService) GetTaskList(ctx context.Context, req *connect.Request[v1.GetTaskListRequest]) (*connect.Response[v1.GetTaskListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "GetTaskList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.GetTaskListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockWithdrawalServiceMockRecorder) GetTaskList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTaskList", reflect.TypeOf((*MockWithdrawalService)(nil).GetTaskList), ctx, req)
}

func (m *MockWithdrawalServiceClient) GetTaskList(ctx context.Context, req *connect.Request[v1.GetTaskListRequest]) (*connect.Response[v1.GetTaskListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "GetTaskList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.GetTaskListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockWithdrawalServiceClientMockRecorder) GetTaskList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTaskList", reflect.TypeOf((*MockWithdrawalService)(nil).GetTaskList), ctx, req)
}

func (m *MockWithdrawalService) Run(ctx context.Context, req *connect.Request[v1.RunRequest]) (*connect.Response[v1.RunResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "Run", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.RunResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockWithdrawalServiceMockRecorder) Run(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockWithdrawalService)(nil).Run), ctx, req)
}

func (m *MockWithdrawalServiceClient) Run(ctx context.Context, req *connect.Request[v1.RunRequest]) (*connect.Response[v1.RunResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "Run", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.RunResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockWithdrawalServiceClientMockRecorder) Run(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockWithdrawalService)(nil).Run), ctx, req)
}

func (m *MockWithdrawalService) Stop(ctx context.Context, req *connect.Request[v1.StopRequest]) (*connect.Response[v1.StopResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "Stop", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.StopResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockWithdrawalServiceMockRecorder) Stop(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockWithdrawalService)(nil).Stop), ctx, req)
}

func (m *MockWithdrawalServiceClient) Stop(ctx context.Context, req *connect.Request[v1.StopRequest]) (*connect.Response[v1.StopResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "Stop", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.StopResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockWithdrawalServiceClientMockRecorder) Stop(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockWithdrawalService)(nil).Stop), ctx, req)
}

func (m *MockWithdrawalService) HealthCheck(ctx context.Context, req *connect.Request[v1.HealthCheckRequest]) (*connect.Response[v1.HealthCheckResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "HealthCheck", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.HealthCheckResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockWithdrawalServiceMockRecorder) HealthCheck(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HealthCheck", reflect.TypeOf((*MockWithdrawalService)(nil).HealthCheck), ctx, req)
}

func (m *MockWithdrawalServiceClient) HealthCheck(ctx context.Context, req *connect.Request[v1.HealthCheckRequest]) (*connect.Response[v1.HealthCheckResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "HealthCheck", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.HealthCheckResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockWithdrawalServiceClientMockRecorder) HealthCheck(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HealthCheck", reflect.TypeOf((*MockWithdrawalService)(nil).HealthCheck), ctx, req)
}

