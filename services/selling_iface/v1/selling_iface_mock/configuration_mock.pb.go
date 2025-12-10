package selling_iface_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/selling_iface/v1"
)

type MockConfigurationLimitService struct {
    ctrl     *gomock.Controller
    recorder *MockConfigurationLimitServiceMockRecorder
}

type MockConfigurationLimitServiceMockRecorder struct {
    mock *MockConfigurationLimitService
}

func NewMockConfigurationLimitService(ctrl *gomock.Controller) *MockConfigurationLimitService {
    mock := &MockConfigurationLimitService{ctrl: ctrl}
    mock.recorder = &MockConfigurationLimitServiceMockRecorder{mock}
    return mock
}

func (m *MockConfigurationLimitService) EXPECT() *MockConfigurationLimitServiceMockRecorder {
    return m.recorder
}

func (m *MockConfigurationLimitService) LimitInvoice(ctx context.Context, req *connect.Request[v1.LimitInvoiceRequest]) (*connect.Response[v1.LimitInvoiceResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "LimitInvoice", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.LimitInvoiceResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockConfigurationLimitServiceMockRecorder) LimitInvoice(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LimitInvoice", reflect.TypeOf((*MockConfigurationLimitService)(nil).LimitInvoice), ctx, req)
}

