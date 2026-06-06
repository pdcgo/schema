package user_iface_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v2 "github.com/pdcgo/schema/services/user_iface/v2"
)

type MockV2AuthService struct {
    ctrl     *gomock.Controller
    recorder *MockV2AuthServiceMockRecorder
}

type MockV2AuthServiceMockRecorder struct {
    mock *MockV2AuthService
}

func NewMockV2AuthService(ctrl *gomock.Controller) *MockV2AuthService {
    mock := &MockV2AuthService{ctrl: ctrl}
    mock.recorder = &MockV2AuthServiceMockRecorder{mock}
    return mock
}

func (m *MockV2AuthService) EXPECT() *MockV2AuthServiceMockRecorder {
    return m.recorder
}

type MockV2AuthServiceClient struct {
    ctrl     *gomock.Controller
    recorder *MockV2AuthServiceClientMockRecorder
}

type MockV2AuthServiceClientMockRecorder struct {
    mock *MockV2AuthServiceClient
}

func NewMockV2AuthServiceClient(ctrl *gomock.Controller) *MockV2AuthServiceClient {
    mock := &MockV2AuthServiceClient{ctrl: ctrl}
    mock.recorder = &MockV2AuthServiceClientMockRecorder{mock}
    return mock
}

func (m *MockV2AuthServiceClient) EXPECT() *MockV2AuthServiceClientMockRecorder {
    return m.recorder
}

func (m *MockV2AuthService) Login(ctx context.Context, req *connect.Request[v2.LoginRequest]) (*connect.Response[v2.LoginResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "Login", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v2.LoginResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockV2AuthServiceMockRecorder) Login(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockV2AuthService)(nil).Login), ctx, req)
}

func (m *MockV2AuthServiceClient) Login(ctx context.Context, req *connect.Request[v2.LoginRequest]) (*connect.Response[v2.LoginResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "Login", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v2.LoginResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockV2AuthServiceClientMockRecorder) Login(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockV2AuthService)(nil).Login), ctx, req)
}

func (m *MockV2AuthService) Logout(ctx context.Context, req *connect.Request[v2.LogoutRequest]) (*connect.Response[v2.LogoutResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "Logout", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v2.LogoutResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockV2AuthServiceMockRecorder) Logout(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logout", reflect.TypeOf((*MockV2AuthService)(nil).Logout), ctx, req)
}

func (m *MockV2AuthServiceClient) Logout(ctx context.Context, req *connect.Request[v2.LogoutRequest]) (*connect.Response[v2.LogoutResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "Logout", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v2.LogoutResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockV2AuthServiceClientMockRecorder) Logout(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logout", reflect.TypeOf((*MockV2AuthService)(nil).Logout), ctx, req)
}

func (m *MockV2AuthService) CheckAccess(ctx context.Context, req *connect.Request[v2.CheckAccessRequest]) (*connect.Response[v2.CheckAccessResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "CheckAccess", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v2.CheckAccessResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockV2AuthServiceMockRecorder) CheckAccess(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckAccess", reflect.TypeOf((*MockV2AuthService)(nil).CheckAccess), ctx, req)
}

func (m *MockV2AuthServiceClient) CheckAccess(ctx context.Context, req *connect.Request[v2.CheckAccessRequest]) (*connect.Response[v2.CheckAccessResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "CheckAccess", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v2.CheckAccessResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockV2AuthServiceClientMockRecorder) CheckAccess(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckAccess", reflect.TypeOf((*MockV2AuthService)(nil).CheckAccess), ctx, req)
}

