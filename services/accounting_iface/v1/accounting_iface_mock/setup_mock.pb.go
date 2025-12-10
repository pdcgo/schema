package accounting_iface_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/accounting_iface/v1"
)

type MockAccountingSetupService struct {
    ctrl     *gomock.Controller
    recorder *MockAccountingSetupServiceMockRecorder
}

type MockAccountingSetupServiceMockRecorder struct {
    mock *MockAccountingSetupService
}

func NewMockAccountingSetupService(ctrl *gomock.Controller) *MockAccountingSetupService {
    mock := &MockAccountingSetupService{ctrl: ctrl}
    mock.recorder = &MockAccountingSetupServiceMockRecorder{mock}
    return mock
}

func (m *MockAccountingSetupService) EXPECT() *MockAccountingSetupServiceMockRecorder {
    return m.recorder
}

type MockAccountingSetupServiceClient struct {
    ctrl     *gomock.Controller
    recorder *MockAccountingSetupServiceClientMockRecorder
}

type MockAccountingSetupServiceClientMockRecorder struct {
    mock *MockAccountingSetupServiceClient
}

func NewMockAccountingSetupServiceClient(ctrl *gomock.Controller) *MockAccountingSetupServiceClient {
    mock := &MockAccountingSetupServiceClient{ctrl: ctrl}
    mock.recorder = &MockAccountingSetupServiceClientMockRecorder{mock}
    return mock
}

func (m *MockAccountingSetupServiceClient) EXPECT() *MockAccountingSetupServiceClientMockRecorder {
    return m.recorder
}

func (m *MockAccountingSetupService) Setup(ctx context.Context, req *connect.Request[v1.SetupRequest], stream *connect.ServerStream[v1.SetupResponse]) error {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "Setup", ctx, req, stream)
    ret0, _ := ret[0].(error)
    return ret0
}

func (mr *MockAccountingSetupServiceMockRecorder) Setup(ctx, req, stream interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Setup", reflect.TypeOf((*MockAccountingSetupService)(nil).Setup), ctx, req, stream)
}

func (m *MockAccountingSetupServiceClient) Setup(ctx context.Context, req *connect.Request[v1.SetupRequest]) (*connect.ServerStreamForClient[v1.SetupResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "Setup", ctx, req)
    ret0, _ := ret[0].(*connect.ServerStreamForClient[v1.SetupResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockAccountingSetupServiceClientMockRecorder) Setup(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Setup", reflect.TypeOf((*MockAccountingSetupService)(nil).Setup), ctx, req)
}

func (m *MockAccountingSetupService) RecalculateDaily(ctx context.Context, req *connect.Request[v1.RecalculateDailyRequest], stream *connect.ServerStream[v1.RecalculateDailyResponse]) error {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "RecalculateDaily", ctx, req, stream)
    ret0, _ := ret[0].(error)
    return ret0
}

func (mr *MockAccountingSetupServiceMockRecorder) RecalculateDaily(ctx, req, stream interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecalculateDaily", reflect.TypeOf((*MockAccountingSetupService)(nil).RecalculateDaily), ctx, req, stream)
}

func (m *MockAccountingSetupServiceClient) RecalculateDaily(ctx context.Context, req *connect.Request[v1.RecalculateDailyRequest]) (*connect.ServerStreamForClient[v1.RecalculateDailyResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "RecalculateDaily", ctx, req)
    ret0, _ := ret[0].(*connect.ServerStreamForClient[v1.RecalculateDailyResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockAccountingSetupServiceClientMockRecorder) RecalculateDaily(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecalculateDaily", reflect.TypeOf((*MockAccountingSetupService)(nil).RecalculateDaily), ctx, req)
}

type MockStreamService struct {
    ctrl     *gomock.Controller
    recorder *MockStreamServiceMockRecorder
}

type MockStreamServiceMockRecorder struct {
    mock *MockStreamService
}

func NewMockStreamService(ctrl *gomock.Controller) *MockStreamService {
    mock := &MockStreamService{ctrl: ctrl}
    mock.recorder = &MockStreamServiceMockRecorder{mock}
    return mock
}

func (m *MockStreamService) EXPECT() *MockStreamServiceMockRecorder {
    return m.recorder
}

type MockStreamServiceClient struct {
    ctrl     *gomock.Controller
    recorder *MockStreamServiceClientMockRecorder
}

type MockStreamServiceClientMockRecorder struct {
    mock *MockStreamServiceClient
}

func NewMockStreamServiceClient(ctrl *gomock.Controller) *MockStreamServiceClient {
    mock := &MockStreamServiceClient{ctrl: ctrl}
    mock.recorder = &MockStreamServiceClientMockRecorder{mock}
    return mock
}

func (m *MockStreamServiceClient) EXPECT() *MockStreamServiceClientMockRecorder {
    return m.recorder
}

func (m *MockStreamService) DummyStream(ctx context.Context, stream *connect.BidiStream[v1.DummyStreamRequest, v1.DummyStreamResponse]) error {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "DummyStream", ctx, stream)
    ret0, _ := ret[0].(error)
    return ret0
}

func (mr *MockStreamServiceMockRecorder) DummyStream(ctx, stream interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DummyStream", reflect.TypeOf((*MockStreamService)(nil).DummyStream), ctx, stream)
}

func (m *MockStreamServiceClient) DummyStream(ctx context.Context) *connect.BidiStreamForClient[v1.DummyStreamRequest, v1.DummyStreamResponse] {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "DummyStream", ctx)
    ret0, _ := ret[0].(*connect.BidiStreamForClient[v1.DummyStreamRequest, v1.DummyStreamResponse])
    return ret0
}

func (mr *MockStreamServiceClientMockRecorder) DummyStream(ctx, stream interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DummyStream", reflect.TypeOf((*MockStreamService)(nil).DummyStream), ctx)
}

