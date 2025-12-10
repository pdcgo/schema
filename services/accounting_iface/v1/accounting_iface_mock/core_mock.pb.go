package accounting_iface_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/accounting_iface/v1"
)

type MockCoreService struct {
    ctrl     *gomock.Controller
    recorder *MockCoreServiceMockRecorder
}

type MockCoreServiceMockRecorder struct {
    mock *MockCoreService
}

func NewMockCoreService(ctrl *gomock.Controller) *MockCoreService {
    mock := &MockCoreService{ctrl: ctrl}
    mock.recorder = &MockCoreServiceMockRecorder{mock}
    return mock
}

func (m *MockCoreService) EXPECT() *MockCoreServiceMockRecorder {
    return m.recorder
}

func (m *MockCoreService) AccountKeyList(ctx context.Context, req *connect.Request[v1.AccountKeyListRequest]) (*connect.Response[v1.AccountKeyListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "AccountKeyList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.AccountKeyListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockCoreServiceMockRecorder) AccountKeyList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccountKeyList", reflect.TypeOf((*MockCoreService)(nil).AccountKeyList), ctx, req)
}

func (m *MockCoreService) TypeLabelList(ctx context.Context, req *connect.Request[v1.TypeLabelListRequest]) (*connect.Response[v1.TypeLabelListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TypeLabelList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.TypeLabelListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockCoreServiceMockRecorder) TypeLabelList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TypeLabelList", reflect.TypeOf((*MockCoreService)(nil).TypeLabelList), ctx, req)
}

