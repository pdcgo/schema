package team_iface_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/team_iface/v1"
)

type MockTeamService struct {
    ctrl     *gomock.Controller
    recorder *MockTeamServiceMockRecorder
}

type MockTeamServiceMockRecorder struct {
    mock *MockTeamService
}

func NewMockTeamService(ctrl *gomock.Controller) *MockTeamService {
    mock := &MockTeamService{ctrl: ctrl}
    mock.recorder = &MockTeamServiceMockRecorder{mock}
    return mock
}

func (m *MockTeamService) EXPECT() *MockTeamServiceMockRecorder {
    return m.recorder
}

type MockTeamServiceClient struct {
    ctrl     *gomock.Controller
    recorder *MockTeamServiceClientMockRecorder
}

type MockTeamServiceClientMockRecorder struct {
    mock *MockTeamServiceClient
}

func NewMockTeamServiceClient(ctrl *gomock.Controller) *MockTeamServiceClient {
    mock := &MockTeamServiceClient{ctrl: ctrl}
    mock.recorder = &MockTeamServiceClientMockRecorder{mock}
    return mock
}

func (m *MockTeamServiceClient) EXPECT() *MockTeamServiceClientMockRecorder {
    return m.recorder
}

func (m *MockTeamService) TeamCreate(ctx context.Context, req *connect.Request[v1.TeamCreateRequest], stream *connect.ServerStream[v1.TeamCreateResponse]) error {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TeamCreate", ctx, req, stream)
    ret0, _ := ret[0].(error)
    return ret0
}

func (mr *MockTeamServiceMockRecorder) TeamCreate(ctx, req, stream interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeamCreate", reflect.TypeOf((*MockTeamService)(nil).TeamCreate), ctx, req, stream)
}

func (m *MockTeamServiceClient) TeamCreate(ctx context.Context, req *connect.Request[v1.TeamCreateRequest]) (*connect.ServerStreamForClient[v1.TeamCreateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TeamCreate", ctx, req)
    ret0, _ := ret[0].(*connect.ServerStreamForClient[v1.TeamCreateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockTeamServiceClientMockRecorder) TeamCreate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeamCreate", reflect.TypeOf((*MockTeamService)(nil).TeamCreate), ctx, req)
}

func (m *MockTeamService) TeamUpdate(ctx context.Context, req *connect.Request[v1.TeamUpdateRequest]) (*connect.Response[v1.TeamUpdateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TeamUpdate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.TeamUpdateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockTeamServiceMockRecorder) TeamUpdate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeamUpdate", reflect.TypeOf((*MockTeamService)(nil).TeamUpdate), ctx, req)
}

func (m *MockTeamServiceClient) TeamUpdate(ctx context.Context, req *connect.Request[v1.TeamUpdateRequest]) (*connect.Response[v1.TeamUpdateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TeamUpdate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.TeamUpdateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockTeamServiceClientMockRecorder) TeamUpdate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeamUpdate", reflect.TypeOf((*MockTeamService)(nil).TeamUpdate), ctx, req)
}

func (m *MockTeamService) TeamDelete(ctx context.Context, req *connect.Request[v1.TeamDeleteRequest]) (*connect.Response[v1.TeamDeleteResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TeamDelete", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.TeamDeleteResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockTeamServiceMockRecorder) TeamDelete(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeamDelete", reflect.TypeOf((*MockTeamService)(nil).TeamDelete), ctx, req)
}

func (m *MockTeamServiceClient) TeamDelete(ctx context.Context, req *connect.Request[v1.TeamDeleteRequest]) (*connect.Response[v1.TeamDeleteResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TeamDelete", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.TeamDeleteResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockTeamServiceClientMockRecorder) TeamDelete(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeamDelete", reflect.TypeOf((*MockTeamService)(nil).TeamDelete), ctx, req)
}

func (m *MockTeamService) TeamList(ctx context.Context, req *connect.Request[v1.TeamListRequest]) (*connect.Response[v1.TeamListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TeamList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.TeamListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockTeamServiceMockRecorder) TeamList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeamList", reflect.TypeOf((*MockTeamService)(nil).TeamList), ctx, req)
}

func (m *MockTeamServiceClient) TeamList(ctx context.Context, req *connect.Request[v1.TeamListRequest]) (*connect.Response[v1.TeamListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TeamList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.TeamListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockTeamServiceClientMockRecorder) TeamList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeamList", reflect.TypeOf((*MockTeamService)(nil).TeamList), ctx, req)
}

func (m *MockTeamService) TeamDetail(ctx context.Context, req *connect.Request[v1.TeamDetailRequest]) (*connect.Response[v1.TeamDetailResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TeamDetail", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.TeamDetailResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockTeamServiceMockRecorder) TeamDetail(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeamDetail", reflect.TypeOf((*MockTeamService)(nil).TeamDetail), ctx, req)
}

func (m *MockTeamServiceClient) TeamDetail(ctx context.Context, req *connect.Request[v1.TeamDetailRequest]) (*connect.Response[v1.TeamDetailResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TeamDetail", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.TeamDetailResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockTeamServiceClientMockRecorder) TeamDetail(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeamDetail", reflect.TypeOf((*MockTeamService)(nil).TeamDetail), ctx, req)
}

func (m *MockTeamService) TeamByIds(ctx context.Context, req *connect.Request[v1.TeamByIdsRequest]) (*connect.Response[v1.TeamByIdsResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TeamByIds", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.TeamByIdsResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockTeamServiceMockRecorder) TeamByIds(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeamByIds", reflect.TypeOf((*MockTeamService)(nil).TeamByIds), ctx, req)
}

func (m *MockTeamServiceClient) TeamByIds(ctx context.Context, req *connect.Request[v1.TeamByIdsRequest]) (*connect.Response[v1.TeamByIdsResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TeamByIds", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.TeamByIdsResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockTeamServiceClientMockRecorder) TeamByIds(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeamByIds", reflect.TypeOf((*MockTeamService)(nil).TeamByIds), ctx, req)
}

func (m *MockTeamService) TeamInfoUpdate(ctx context.Context, req *connect.Request[v1.TeamInfoUpdateRequest]) (*connect.Response[v1.TeamInfoUpdateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TeamInfoUpdate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.TeamInfoUpdateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockTeamServiceMockRecorder) TeamInfoUpdate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeamInfoUpdate", reflect.TypeOf((*MockTeamService)(nil).TeamInfoUpdate), ctx, req)
}

func (m *MockTeamServiceClient) TeamInfoUpdate(ctx context.Context, req *connect.Request[v1.TeamInfoUpdateRequest]) (*connect.Response[v1.TeamInfoUpdateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TeamInfoUpdate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.TeamInfoUpdateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockTeamServiceClientMockRecorder) TeamInfoUpdate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeamInfoUpdate", reflect.TypeOf((*MockTeamService)(nil).TeamInfoUpdate), ctx, req)
}

