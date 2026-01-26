package common_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/common/v1"
)

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

type MockUserServiceClient struct {
    ctrl     *gomock.Controller
    recorder *MockUserServiceClientMockRecorder
}

type MockUserServiceClientMockRecorder struct {
    mock *MockUserServiceClient
}

func NewMockUserServiceClient(ctrl *gomock.Controller) *MockUserServiceClient {
    mock := &MockUserServiceClient{ctrl: ctrl}
    mock.recorder = &MockUserServiceClientMockRecorder{mock}
    return mock
}

func (m *MockUserServiceClient) EXPECT() *MockUserServiceClientMockRecorder {
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

func (m *MockUserServiceClient) PublicUserIDs(ctx context.Context, req *connect.Request[v1.PublicUserIDsRequest]) (*connect.Response[v1.PublicUserIDsResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "PublicUserIDs", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.PublicUserIDsResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockUserServiceClientMockRecorder) PublicUserIDs(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublicUserIDs", reflect.TypeOf((*MockUserService)(nil).PublicUserIDs), ctx, req)
}

