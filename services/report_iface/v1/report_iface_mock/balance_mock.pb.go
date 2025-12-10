package report_iface_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/report_iface/v1"
)

type MockBalanceService struct {
    ctrl     *gomock.Controller
    recorder *MockBalanceServiceMockRecorder
}

type MockBalanceServiceMockRecorder struct {
    mock *MockBalanceService
}

func NewMockBalanceService(ctrl *gomock.Controller) *MockBalanceService {
    mock := &MockBalanceService{ctrl: ctrl}
    mock.recorder = &MockBalanceServiceMockRecorder{mock}
    return mock
}

func (m *MockBalanceService) EXPECT() *MockBalanceServiceMockRecorder {
    return m.recorder
}

func (m *MockBalanceService) BalanceResync(ctx context.Context, req *connect.Request[v1.BalanceResyncRequest], stream *connect.ServerStream[v1.BalanceResyncResponse]) error {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "BalanceResync", ctx, req, stream)
    ret0, _ := ret[0].(error)
    return ret0
}

func (mr *MockBalanceServiceMockRecorder) BalanceResync(ctx, req, stream interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BalanceResync", reflect.TypeOf((*MockBalanceService)(nil).BalanceResync), ctx, req, stream)
}

