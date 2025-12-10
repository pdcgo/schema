package access_iface_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/access_iface/v1"
)

type MockFrontendAccessService struct {
    ctrl     *gomock.Controller
    recorder *MockFrontendAccessServiceMockRecorder
}

type MockFrontendAccessServiceMockRecorder struct {
    mock *MockFrontendAccessService
}

func NewMockFrontendAccessService(ctrl *gomock.Controller) *MockFrontendAccessService {
    mock := &MockFrontendAccessService{ctrl: ctrl}
    mock.recorder = &MockFrontendAccessServiceMockRecorder{mock}
    return mock
}

func (m *MockFrontendAccessService) EXPECT() *MockFrontendAccessServiceMockRecorder {
    return m.recorder
}

func (m *MockFrontendAccessService) SetupAccess(ctx context.Context, req *connect.Request[v1.SetupAccessRequest], stream *connect.ServerStream[v1.SetupAccessResponse]) error {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "SetupAccess", ctx, req, stream)
    ret0, _ := ret[0].(error)
    return ret0
}

func (mr *MockFrontendAccessServiceMockRecorder) SetupAccess(ctx, req, stream interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetupAccess", reflect.TypeOf((*MockFrontendAccessService)(nil).SetupAccess), ctx, req, stream)
}

func (m *MockFrontendAccessService) MenuAccess(ctx context.Context, req *connect.Request[v1.MenuAccessRequest]) (*connect.Response[v1.MenuAccessResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "MenuAccess", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.MenuAccessResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockFrontendAccessServiceMockRecorder) MenuAccess(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MenuAccess", reflect.TypeOf((*MockFrontendAccessService)(nil).MenuAccess), ctx, req)
}

