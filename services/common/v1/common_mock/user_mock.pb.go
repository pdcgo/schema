package common_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/common/v1"
)

type MockLoginService struct {
    ctrl     *gomock.Controller
    recorder *MockLoginServiceMockRecorder
}

type MockLoginServiceMockRecorder struct {
    mock *MockLoginService
}

func NewMockLoginService(ctrl *gomock.Controller) *MockLoginService {
    mock := &MockLoginService{ctrl: ctrl}
    mock.recorder = &MockLoginServiceMockRecorder{mock}
    return mock
}

func (m *MockLoginService) EXPECT() *MockLoginServiceMockRecorder {
    return m.recorder
}

func (m *MockLoginService) Login(ctx context.Context, req *connect.Request[v1.LoginRequest]) (*connect.Response[v1.LoginResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "Login", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.LoginResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockLoginServiceMockRecorder) Login(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockLoginService)(nil).Login), ctx, req)
}

type MockUserService struct {
    ctrl     *gomock.Controller
    recorder *MockUserServiceMockRecorder
}

type MockUserServiceMockRecorder struct {
    mock *MockUserService
}

func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
    mock := &MockUserService{ctrl: ctrl}
    mock.recorder = &MockUserServiceMockRecorder{mock}
    return mock
}

func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
    return m.recorder
}

func (m *MockUserService) PublicUserIDs(ctx context.Context, req *connect.Request[v1.PublicUserIDsRequest]) (*connect.Response[v1.PublicUserIDsResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "PublicUserIDs", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.PublicUserIDsResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockUserServiceMockRecorder) PublicUserIDs(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublicUserIDs", reflect.TypeOf((*MockUserService)(nil).PublicUserIDs), ctx, req)
}

