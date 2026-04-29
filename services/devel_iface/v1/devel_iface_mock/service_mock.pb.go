package devel_iface_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/devel_iface/v1"
)

type MockDevelToolService struct {
    ctrl     *gomock.Controller
    recorder *MockDevelToolServiceMockRecorder
}

type MockDevelToolServiceMockRecorder struct {
    mock *MockDevelToolService
}

func NewMockDevelToolService(ctrl *gomock.Controller) *MockDevelToolService {
    mock := &MockDevelToolService{ctrl: ctrl}
    mock.recorder = &MockDevelToolServiceMockRecorder{mock}
    return mock
}

func (m *MockDevelToolService) EXPECT() *MockDevelToolServiceMockRecorder {
    return m.recorder
}

type MockDevelToolServiceClient struct {
    ctrl     *gomock.Controller
    recorder *MockDevelToolServiceClientMockRecorder
}

type MockDevelToolServiceClientMockRecorder struct {
    mock *MockDevelToolServiceClient
}

func NewMockDevelToolServiceClient(ctrl *gomock.Controller) *MockDevelToolServiceClient {
    mock := &MockDevelToolServiceClient{ctrl: ctrl}
    mock.recorder = &MockDevelToolServiceClientMockRecorder{mock}
    return mock
}

func (m *MockDevelToolServiceClient) EXPECT() *MockDevelToolServiceClientMockRecorder {
    return m.recorder
}

func (m *MockDevelToolService) SendEvent(ctx context.Context, req *connect.Request[v1.SendEventRequest]) (*connect.Response[v1.SendEventResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "SendEvent", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.SendEventResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockDevelToolServiceMockRecorder) SendEvent(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendEvent", reflect.TypeOf((*MockDevelToolService)(nil).SendEvent), ctx, req)
}

func (m *MockDevelToolServiceClient) SendEvent(ctx context.Context, req *connect.Request[v1.SendEventRequest]) (*connect.Response[v1.SendEventResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "SendEvent", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.SendEventResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockDevelToolServiceClientMockRecorder) SendEvent(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendEvent", reflect.TypeOf((*MockDevelToolService)(nil).SendEvent), ctx, req)
}

