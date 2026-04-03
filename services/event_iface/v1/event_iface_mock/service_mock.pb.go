package event_iface_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/event_iface/v1"
)

type MockEventService struct {
    ctrl     *gomock.Controller
    recorder *MockEventServiceMockRecorder
}

type MockEventServiceMockRecorder struct {
    mock *MockEventService
}

func NewMockEventService(ctrl *gomock.Controller) *MockEventService {
    mock := &MockEventService{ctrl: ctrl}
    mock.recorder = &MockEventServiceMockRecorder{mock}
    return mock
}

func (m *MockEventService) EXPECT() *MockEventServiceMockRecorder {
    return m.recorder
}

type MockEventServiceClient struct {
    ctrl     *gomock.Controller
    recorder *MockEventServiceClientMockRecorder
}

type MockEventServiceClientMockRecorder struct {
    mock *MockEventServiceClient
}

func NewMockEventServiceClient(ctrl *gomock.Controller) *MockEventServiceClient {
    mock := &MockEventServiceClient{ctrl: ctrl}
    mock.recorder = &MockEventServiceClientMockRecorder{mock}
    return mock
}

func (m *MockEventServiceClient) EXPECT() *MockEventServiceClientMockRecorder {
    return m.recorder
}

func (m *MockEventService) Send(ctx context.Context, req *connect.Request[v1.SendRequest]) (*connect.Response[v1.SendResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "Send", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.SendResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockEventServiceMockRecorder) Send(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockEventService)(nil).Send), ctx, req)
}

func (m *MockEventServiceClient) Send(ctx context.Context, req *connect.Request[v1.SendRequest]) (*connect.Response[v1.SendResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "Send", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.SendResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockEventServiceClientMockRecorder) Send(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockEventService)(nil).Send), ctx, req)
}

type MockPushService struct {
    ctrl     *gomock.Controller
    recorder *MockPushServiceMockRecorder
}

type MockPushServiceMockRecorder struct {
    mock *MockPushService
}

func NewMockPushService(ctrl *gomock.Controller) *MockPushService {
    mock := &MockPushService{ctrl: ctrl}
    mock.recorder = &MockPushServiceMockRecorder{mock}
    return mock
}

func (m *MockPushService) EXPECT() *MockPushServiceMockRecorder {
    return m.recorder
}

type MockPushServiceClient struct {
    ctrl     *gomock.Controller
    recorder *MockPushServiceClientMockRecorder
}

type MockPushServiceClientMockRecorder struct {
    mock *MockPushServiceClient
}

func NewMockPushServiceClient(ctrl *gomock.Controller) *MockPushServiceClient {
    mock := &MockPushServiceClient{ctrl: ctrl}
    mock.recorder = &MockPushServiceClientMockRecorder{mock}
    return mock
}

func (m *MockPushServiceClient) EXPECT() *MockPushServiceClientMockRecorder {
    return m.recorder
}

func (m *MockPushService) Push(ctx context.Context, req *connect.Request[v1.PushRequest]) (*connect.Response[v1.PushResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "Push", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.PushResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockPushServiceMockRecorder) Push(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Push", reflect.TypeOf((*MockPushService)(nil).Push), ctx, req)
}

func (m *MockPushServiceClient) Push(ctx context.Context, req *connect.Request[v1.PushRequest]) (*connect.Response[v1.PushResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "Push", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.PushResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockPushServiceClientMockRecorder) Push(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Push", reflect.TypeOf((*MockPushService)(nil).Push), ctx, req)
}

