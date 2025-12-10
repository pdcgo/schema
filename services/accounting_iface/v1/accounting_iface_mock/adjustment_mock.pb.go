package accounting_iface_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/accounting_iface/v1"
)

type MockAdjustmentService struct {
    ctrl     *gomock.Controller
    recorder *MockAdjustmentServiceMockRecorder
}

type MockAdjustmentServiceMockRecorder struct {
    mock *MockAdjustmentService
}

func NewMockAdjustmentService(ctrl *gomock.Controller) *MockAdjustmentService {
    mock := &MockAdjustmentService{ctrl: ctrl}
    mock.recorder = &MockAdjustmentServiceMockRecorder{mock}
    return mock
}

func (m *MockAdjustmentService) EXPECT() *MockAdjustmentServiceMockRecorder {
    return m.recorder
}

func (m *MockAdjustmentService) AccountAdjustment(ctx context.Context, req *connect.Request[v1.AccountAdjustmentRequest]) (*connect.Response[v1.AccountAdjustmentResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "AccountAdjustment", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.AccountAdjustmentResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockAdjustmentServiceMockRecorder) AccountAdjustment(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccountAdjustment", reflect.TypeOf((*MockAdjustmentService)(nil).AccountAdjustment), ctx, req)
}

func (m *MockAdjustmentService) AdjCreate(ctx context.Context, req *connect.Request[v1.AdjCreateRequest]) (*connect.Response[v1.AdjCreateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "AdjCreate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.AdjCreateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockAdjustmentServiceMockRecorder) AdjCreate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdjCreate", reflect.TypeOf((*MockAdjustmentService)(nil).AdjCreate), ctx, req)
}

