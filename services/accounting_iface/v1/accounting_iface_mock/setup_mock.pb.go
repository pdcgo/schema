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

