package warehouse_iface_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/warehouse_iface/v1"
)

type MockInboundService struct {
    ctrl     *gomock.Controller
    recorder *MockInboundServiceMockRecorder
}

type MockInboundServiceMockRecorder struct {
    mock *MockInboundService
}

func NewMockInboundService(ctrl *gomock.Controller) *MockInboundService {
    mock := &MockInboundService{ctrl: ctrl}
    mock.recorder = &MockInboundServiceMockRecorder{mock}
    return mock
}

func (m *MockInboundService) EXPECT() *MockInboundServiceMockRecorder {
    return m.recorder
}

type MockInboundServiceClient struct {
    ctrl     *gomock.Controller
    recorder *MockInboundServiceClientMockRecorder
}

type MockInboundServiceClientMockRecorder struct {
    mock *MockInboundServiceClient
}

func NewMockInboundServiceClient(ctrl *gomock.Controller) *MockInboundServiceClient {
    mock := &MockInboundServiceClient{ctrl: ctrl}
    mock.recorder = &MockInboundServiceClientMockRecorder{mock}
    return mock
}

func (m *MockInboundServiceClient) EXPECT() *MockInboundServiceClientMockRecorder {
    return m.recorder
}

func (m *MockInboundService) InboundAccept(ctx context.Context, req *connect.Request[v1.InboundAcceptRequest]) (*connect.Response[v1.InboundAcceptResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "InboundAccept", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.InboundAcceptResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInboundServiceMockRecorder) InboundAccept(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InboundAccept", reflect.TypeOf((*MockInboundService)(nil).InboundAccept), ctx, req)
}

func (m *MockInboundServiceClient) InboundAccept(ctx context.Context, req *connect.Request[v1.InboundAcceptRequest]) (*connect.Response[v1.InboundAcceptResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "InboundAccept", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.InboundAcceptResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInboundServiceClientMockRecorder) InboundAccept(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InboundAccept", reflect.TypeOf((*MockInboundService)(nil).InboundAccept), ctx, req)
}

