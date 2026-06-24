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

func (m *MockV2UserService) UserByIDs(ctx context.Context, req *connect.Request[v2.UserByIDsRequest]) (*connect.Response[v2.UserByIDsResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "UserByIDs", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v2.UserByIDsResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockV2UserServiceMockRecorder) UserByIDs(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserByIDs", reflect.TypeOf((*MockV2UserService)(nil).UserByIDs), ctx, req)
}

func (m *MockV2UserServiceClient) UserByIDs(ctx context.Context, req *connect.Request[v2.UserByIDsRequest]) (*connect.Response[v2.UserByIDsResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "UserByIDs", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v2.UserByIDsResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockV2UserServiceClientMockRecorder) UserByIDs(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserByIDs", reflect.TypeOf((*MockV2UserService)(nil).UserByIDs), ctx, req)
}

func (m *MockV2UserService) UserList(ctx context.Context, req *connect.Request[v2.UserListRequest]) (*connect.Response[v2.UserListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "UserList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v2.UserListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockV2UserServiceMockRecorder) UserList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserList", reflect.TypeOf((*MockV2UserService)(nil).UserList), ctx, req)
}

func (m *MockV2UserServiceClient) UserList(ctx context.Context, req *connect.Request[v2.UserListRequest]) (*connect.Response[v2.UserListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "UserList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v2.UserListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockV2UserServiceClientMockRecorder) UserList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserList", reflect.TypeOf((*MockV2UserService)(nil).UserList), ctx, req)
}

func (m *MockV2UserService) UserChangePhoneNumber(ctx context.Context, req *connect.Request[v2.UserChangePhoneNumberRequest]) (*connect.Response[v2.UserChangePhoneNumberResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "UserChangePhoneNumber", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v2.UserChangePhoneNumberResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockV2UserServiceMockRecorder) UserChangePhoneNumber(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserChangePhoneNumber", reflect.TypeOf((*MockV2UserService)(nil).UserChangePhoneNumber), ctx, req)
}

func (m *MockV2UserServiceClient) UserChangePhoneNumber(ctx context.Context, req *connect.Request[v2.UserChangePhoneNumberRequest]) (*connect.Response[v2.UserChangePhoneNumberResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "UserChangePhoneNumber", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v2.UserChangePhoneNumberResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockV2UserServiceClientMockRecorder) UserChangePhoneNumber(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserChangePhoneNumber", reflect.TypeOf((*MockV2UserService)(nil).UserChangePhoneNumber), ctx, req)
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

func (m *MockV2UserService) TeamAccessList(ctx context.Context, req *connect.Request[v2.TeamAccessListRequest]) (*connect.Response[v2.TeamAccessListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TeamAccessList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v2.TeamAccessListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockV2UserServiceMockRecorder) TeamAccessList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeamAccessList", reflect.TypeOf((*MockV2UserService)(nil).TeamAccessList), ctx, req)
}

func (m *MockV2UserServiceClient) TeamAccessList(ctx context.Context, req *connect.Request[v2.TeamAccessListRequest]) (*connect.Response[v2.TeamAccessListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TeamAccessList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v2.TeamAccessListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockV2UserServiceClientMockRecorder) TeamAccessList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeamAccessList", reflect.TypeOf((*MockV2UserService)(nil).TeamAccessList), ctx, req)
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

