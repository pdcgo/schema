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

func (m *MockInboundService) InboundReject(ctx context.Context, req *connect.Request[v1.InboundRejectRequest]) (*connect.Response[v1.InboundRejectResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "InboundReject", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.InboundRejectResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInboundServiceMockRecorder) InboundReject(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InboundReject", reflect.TypeOf((*MockInboundService)(nil).InboundReject), ctx, req)
}

func (m *MockInboundServiceClient) InboundReject(ctx context.Context, req *connect.Request[v1.InboundRejectRequest]) (*connect.Response[v1.InboundRejectResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "InboundReject", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.InboundRejectResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInboundServiceClientMockRecorder) InboundReject(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InboundReject", reflect.TypeOf((*MockInboundService)(nil).InboundReject), ctx, req)
}

func (m *MockInboundService) InboundDetailSearch(ctx context.Context, req *connect.Request[v1.InboundDetailSearchRequest]) (*connect.Response[v1.InboundDetailSearchResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "InboundDetailSearch", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.InboundDetailSearchResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInboundServiceMockRecorder) InboundDetailSearch(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InboundDetailSearch", reflect.TypeOf((*MockInboundService)(nil).InboundDetailSearch), ctx, req)
}

func (m *MockInboundServiceClient) InboundDetailSearch(ctx context.Context, req *connect.Request[v1.InboundDetailSearchRequest]) (*connect.Response[v1.InboundDetailSearchResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "InboundDetailSearch", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.InboundDetailSearchResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInboundServiceClientMockRecorder) InboundDetailSearch(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InboundDetailSearch", reflect.TypeOf((*MockInboundService)(nil).InboundDetailSearch), ctx, req)
}

func (m *MockInboundService) InboundCancel(ctx context.Context, req *connect.Request[v1.InboundCancelRequest]) (*connect.Response[v1.InboundCancelResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "InboundCancel", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.InboundCancelResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInboundServiceMockRecorder) InboundCancel(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InboundCancel", reflect.TypeOf((*MockInboundService)(nil).InboundCancel), ctx, req)
}

func (m *MockInboundServiceClient) InboundCancel(ctx context.Context, req *connect.Request[v1.InboundCancelRequest]) (*connect.Response[v1.InboundCancelResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "InboundCancel", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.InboundCancelResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInboundServiceClientMockRecorder) InboundCancel(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InboundCancel", reflect.TypeOf((*MockInboundService)(nil).InboundCancel), ctx, req)
}

