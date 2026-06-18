package user_iface_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v2 "github.com/pdcgo/schema/services/user_iface/v2"
)

type MockV2UserService struct {
    ctrl     *gomock.Controller
    recorder *MockV2UserServiceMockRecorder
}

type MockV2UserServiceMockRecorder struct {
    mock *MockV2UserService
}

func NewMockV2UserService(ctrl *gomock.Controller) *MockV2UserService {
    mock := &MockV2UserService{ctrl: ctrl}
    mock.recorder = &MockV2UserServiceMockRecorder{mock}
    return mock
}

func (m *MockV2UserService) EXPECT() *MockV2UserServiceMockRecorder {
    return m.recorder
}

type MockV2UserServiceClient struct {
    ctrl     *gomock.Controller
    recorder *MockV2UserServiceClientMockRecorder
}

type MockV2UserServiceClientMockRecorder struct {
    mock *MockV2UserServiceClient
}

func NewMockV2UserServiceClient(ctrl *gomock.Controller) *MockV2UserServiceClient {
    mock := &MockV2UserServiceClient{ctrl: ctrl}
    mock.recorder = &MockV2UserServiceClientMockRecorder{mock}
    return mock
}

func (m *MockV2UserServiceClient) EXPECT() *MockV2UserServiceClientMockRecorder {
    return m.recorder
}

func (m *MockV2UserService) CreateUser(ctx context.Context, req *connect.Request[v2.CreateUserRequest]) (*connect.Response[v2.CreateUserResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "CreateUser", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v2.CreateUserResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockV2UserServiceMockRecorder) CreateUser(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockV2UserService)(nil).CreateUser), ctx, req)
}

func (m *MockV2UserServiceClient) CreateUser(ctx context.Context, req *connect.Request[v2.CreateUserRequest]) (*connect.Response[v2.CreateUserResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "CreateUser", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v2.CreateUserResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockV2UserServiceClientMockRecorder) CreateUser(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockV2UserService)(nil).CreateUser), ctx, req)
}

func (m *MockV2UserService) UpdateUser(ctx context.Context, req *connect.Request[v2.UpdateUserRequest]) (*connect.Response[v2.UpdateUserResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "UpdateUser", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v2.UpdateUserResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockV2UserServiceMockRecorder) UpdateUser(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockV2UserService)(nil).UpdateUser), ctx, req)
}

func (m *MockV2UserServiceClient) UpdateUser(ctx context.Context, req *connect.Request[v2.UpdateUserRequest]) (*connect.Response[v2.UpdateUserResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "UpdateUser", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v2.UpdateUserResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockV2UserServiceClientMockRecorder) UpdateUser(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockV2UserService)(nil).UpdateUser), ctx, req)
}

func (m *MockV2UserService) DeleteUser(ctx context.Context, req *connect.Request[v2.DeleteUserRequest]) (*connect.Response[v2.DeleteUserResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "DeleteUser", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v2.DeleteUserResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockV2UserServiceMockRecorder) DeleteUser(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockV2UserService)(nil).DeleteUser), ctx, req)
}

func (m *MockV2UserServiceClient) DeleteUser(ctx context.Context, req *connect.Request[v2.DeleteUserRequest]) (*connect.Response[v2.DeleteUserResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "DeleteUser", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v2.DeleteUserResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockV2UserServiceClientMockRecorder) DeleteUser(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockV2UserService)(nil).DeleteUser), ctx, req)
}

func (m *MockV2UserService) SuspendUser(ctx context.Context, req *connect.Request[v2.SuspendUserRequest]) (*connect.Response[v2.SuspendUserResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "SuspendUser", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v2.SuspendUserResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockV2UserServiceMockRecorder) SuspendUser(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SuspendUser", reflect.TypeOf((*MockV2UserService)(nil).SuspendUser), ctx, req)
}

func (m *MockV2UserServiceClient) SuspendUser(ctx context.Context, req *connect.Request[v2.SuspendUserRequest]) (*connect.Response[v2.SuspendUserResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "SuspendUser", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v2.SuspendUserResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockV2UserServiceClientMockRecorder) SuspendUser(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SuspendUser", reflect.TypeOf((*MockV2UserService)(nil).SuspendUser), ctx, req)
}

func (m *MockV2UserService) SearchUser(ctx context.Context, req *connect.Request[v2.SearchUserRequest]) (*connect.Response[v2.SearchUserResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "SearchUser", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v2.SearchUserResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockV2UserServiceMockRecorder) SearchUser(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchUser", reflect.TypeOf((*MockV2UserService)(nil).SearchUser), ctx, req)
}

func (m *MockV2UserServiceClient) SearchUser(ctx context.Context, req *connect.Request[v2.SearchUserRequest]) (*connect.Response[v2.SearchUserResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "SearchUser", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v2.SearchUserResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockV2UserServiceClientMockRecorder) SearchUser(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchUser", reflect.TypeOf((*MockV2UserService)(nil).SearchUser), ctx, req)
}

func (m *MockV2UserService) ResetPassword(ctx context.Context, req *connect.Request[v2.ResetPasswordRequest]) (*connect.Response[v2.ResetPasswordResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ResetPassword", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v2.ResetPasswordResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockV2UserServiceMockRecorder) ResetPassword(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetPassword", reflect.TypeOf((*MockV2UserService)(nil).ResetPassword), ctx, req)
}

func (m *MockV2UserServiceClient) ResetPassword(ctx context.Context, req *connect.Request[v2.ResetPasswordRequest]) (*connect.Response[v2.ResetPasswordResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ResetPassword", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v2.ResetPasswordResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockV2UserServiceClientMockRecorder) ResetPassword(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetPassword", reflect.TypeOf((*MockV2UserService)(nil).ResetPassword), ctx, req)
}

func (m *MockV2UserService) TeamUserList(ctx context.Context, req *connect.Request[v2.TeamUserListRequest]) (*connect.Response[v2.TeamUserListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TeamUserList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v2.TeamUserListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockV2UserServiceMockRecorder) TeamUserList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeamUserList", reflect.TypeOf((*MockV2UserService)(nil).TeamUserList), ctx, req)
}

func (m *MockV2UserServiceClient) TeamUserList(ctx context.Context, req *connect.Request[v2.TeamUserListRequest]) (*connect.Response[v2.TeamUserListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TeamUserList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v2.TeamUserListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockV2UserServiceClientMockRecorder) TeamUserList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeamUserList", reflect.TypeOf((*MockV2UserService)(nil).TeamUserList), ctx, req)
}

func (m *MockV2UserService) TeamUserUpdate(ctx context.Context, req *connect.Request[v2.TeamUserUpdateRequest]) (*connect.Response[v2.TeamUserUpdateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TeamUserUpdate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v2.TeamUserUpdateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockV2UserServiceMockRecorder) TeamUserUpdate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeamUserUpdate", reflect.TypeOf((*MockV2UserService)(nil).TeamUserUpdate), ctx, req)
}

func (m *MockV2UserServiceClient) TeamUserUpdate(ctx context.Context, req *connect.Request[v2.TeamUserUpdateRequest]) (*connect.Response[v2.TeamUserUpdateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TeamUserUpdate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v2.TeamUserUpdateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockV2UserServiceClientMockRecorder) TeamUserUpdate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeamUserUpdate", reflect.TypeOf((*MockV2UserService)(nil).TeamUserUpdate), ctx, req)
}

func (m *MockV2UserService) TeamSynclegacy(ctx context.Context, req *connect.Request[v2.TeamSynclegacyRequest], stream *connect.ServerStream[v2.TeamSynclegacyResponse]) error {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TeamSynclegacy", ctx, req, stream)
    ret0, _ := ret[0].(error)
    return ret0
}

func (mr *MockV2UserServiceMockRecorder) TeamSynclegacy(ctx, req, stream interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeamSynclegacy", reflect.TypeOf((*MockV2UserService)(nil).TeamSynclegacy), ctx, req, stream)
}

func (m *MockV2UserServiceClient) TeamSynclegacy(ctx context.Context, req *connect.Request[v2.TeamSynclegacyRequest]) (*connect.ServerStreamForClient[v2.TeamSynclegacyResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TeamSynclegacy", ctx, req)
    ret0, _ := ret[0].(*connect.ServerStreamForClient[v2.TeamSynclegacyResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockV2UserServiceClientMockRecorder) TeamSynclegacy(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeamSynclegacy", reflect.TypeOf((*MockV2UserService)(nil).TeamSynclegacy), ctx, req)
}

