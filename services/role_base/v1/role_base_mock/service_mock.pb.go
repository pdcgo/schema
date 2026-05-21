package role_base_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/role_base/v1"
)

type MockRoleBaseService struct {
    ctrl     *gomock.Controller
    recorder *MockRoleBaseServiceMockRecorder
}

type MockRoleBaseServiceMockRecorder struct {
    mock *MockRoleBaseService
}

func NewMockRoleBaseService(ctrl *gomock.Controller) *MockRoleBaseService {
    mock := &MockRoleBaseService{ctrl: ctrl}
    mock.recorder = &MockRoleBaseServiceMockRecorder{mock}
    return mock
}

func (m *MockRoleBaseService) EXPECT() *MockRoleBaseServiceMockRecorder {
    return m.recorder
}

type MockRoleBaseServiceClient struct {
    ctrl     *gomock.Controller
    recorder *MockRoleBaseServiceClientMockRecorder
}

type MockRoleBaseServiceClientMockRecorder struct {
    mock *MockRoleBaseServiceClient
}

func NewMockRoleBaseServiceClient(ctrl *gomock.Controller) *MockRoleBaseServiceClient {
    mock := &MockRoleBaseServiceClient{ctrl: ctrl}
    mock.recorder = &MockRoleBaseServiceClientMockRecorder{mock}
    return mock
}

func (m *MockRoleBaseServiceClient) EXPECT() *MockRoleBaseServiceClientMockRecorder {
    return m.recorder
}

func (m *MockRoleBaseService) Login(ctx context.Context, req *connect.Request[v1.LoginRequest]) (*connect.Response[v1.LoginResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "Login", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.LoginResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockRoleBaseServiceMockRecorder) Login(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockRoleBaseService)(nil).Login), ctx, req)
}

func (m *MockRoleBaseServiceClient) Login(ctx context.Context, req *connect.Request[v1.LoginRequest]) (*connect.Response[v1.LoginResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "Login", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.LoginResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockRoleBaseServiceClientMockRecorder) Login(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockRoleBaseService)(nil).Login), ctx, req)
}

func (m *MockRoleBaseService) RequestAccess(ctx context.Context, req *connect.Request[v1.RequestAccessRequest]) (*connect.Response[v1.RequestAccessResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "RequestAccess", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.RequestAccessResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockRoleBaseServiceMockRecorder) RequestAccess(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestAccess", reflect.TypeOf((*MockRoleBaseService)(nil).RequestAccess), ctx, req)
}

func (m *MockRoleBaseServiceClient) RequestAccess(ctx context.Context, req *connect.Request[v1.RequestAccessRequest]) (*connect.Response[v1.RequestAccessResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "RequestAccess", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.RequestAccessResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockRoleBaseServiceClientMockRecorder) RequestAccess(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestAccess", reflect.TypeOf((*MockRoleBaseService)(nil).RequestAccess), ctx, req)
}

func (m *MockRoleBaseService) RoleList(ctx context.Context, req *connect.Request[v1.RoleListRequest]) (*connect.Response[v1.RoleListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "RoleList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.RoleListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockRoleBaseServiceMockRecorder) RoleList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RoleList", reflect.TypeOf((*MockRoleBaseService)(nil).RoleList), ctx, req)
}

func (m *MockRoleBaseServiceClient) RoleList(ctx context.Context, req *connect.Request[v1.RoleListRequest]) (*connect.Response[v1.RoleListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "RoleList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.RoleListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockRoleBaseServiceClientMockRecorder) RoleList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RoleList", reflect.TypeOf((*MockRoleBaseService)(nil).RoleList), ctx, req)
}

func (m *MockRoleBaseService) UserRoleList(ctx context.Context, req *connect.Request[v1.UserRoleListRequest]) (*connect.Response[v1.UserRoleListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "UserRoleList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.UserRoleListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockRoleBaseServiceMockRecorder) UserRoleList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserRoleList", reflect.TypeOf((*MockRoleBaseService)(nil).UserRoleList), ctx, req)
}

func (m *MockRoleBaseServiceClient) UserRoleList(ctx context.Context, req *connect.Request[v1.UserRoleListRequest]) (*connect.Response[v1.UserRoleListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "UserRoleList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.UserRoleListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockRoleBaseServiceClientMockRecorder) UserRoleList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserRoleList", reflect.TypeOf((*MockRoleBaseService)(nil).UserRoleList), ctx, req)
}

func (m *MockRoleBaseService) UserList(ctx context.Context, req *connect.Request[v1.UserListRequest]) (*connect.Response[v1.UserListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "UserList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.UserListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockRoleBaseServiceMockRecorder) UserList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserList", reflect.TypeOf((*MockRoleBaseService)(nil).UserList), ctx, req)
}

func (m *MockRoleBaseServiceClient) UserList(ctx context.Context, req *connect.Request[v1.UserListRequest]) (*connect.Response[v1.UserListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "UserList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.UserListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockRoleBaseServiceClientMockRecorder) UserList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserList", reflect.TypeOf((*MockRoleBaseService)(nil).UserList), ctx, req)
}

func (m *MockRoleBaseService) UserRoleUpdate(ctx context.Context, req *connect.Request[v1.UserRoleUpdateRequest]) (*connect.Response[v1.UserRoleUpdateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "UserRoleUpdate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.UserRoleUpdateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockRoleBaseServiceMockRecorder) UserRoleUpdate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserRoleUpdate", reflect.TypeOf((*MockRoleBaseService)(nil).UserRoleUpdate), ctx, req)
}

func (m *MockRoleBaseServiceClient) UserRoleUpdate(ctx context.Context, req *connect.Request[v1.UserRoleUpdateRequest]) (*connect.Response[v1.UserRoleUpdateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "UserRoleUpdate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.UserRoleUpdateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockRoleBaseServiceClientMockRecorder) UserRoleUpdate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserRoleUpdate", reflect.TypeOf((*MockRoleBaseService)(nil).UserRoleUpdate), ctx, req)
}

