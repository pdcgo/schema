package common_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/common/v1"
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

func (m *MockTeamService) PublicTeamIDs(ctx context.Context, req *connect.Request[v1.PublicTeamIDsRequest]) (*connect.Response[v1.PublicTeamIDsResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "PublicTeamIDs", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.PublicTeamIDsResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockTeamServiceMockRecorder) PublicTeamIDs(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublicTeamIDs", reflect.TypeOf((*MockTeamService)(nil).PublicTeamIDs), ctx, req)
}

func (m *MockTeamService) PublicTeamList(ctx context.Context, req *connect.Request[v1.PublicTeamListRequest]) (*connect.Response[v1.PublicTeamListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "PublicTeamList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.PublicTeamListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockTeamServiceMockRecorder) PublicTeamList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublicTeamList", reflect.TypeOf((*MockTeamService)(nil).PublicTeamList), ctx, req)
}

