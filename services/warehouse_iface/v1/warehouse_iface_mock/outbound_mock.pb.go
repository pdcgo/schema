package warehouse_iface_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/warehouse_iface/v1"
)

type MockOutboundService struct {
    ctrl     *gomock.Controller
    recorder *MockOutboundServiceMockRecorder
}

type MockOutboundServiceMockRecorder struct {
    mock *MockOutboundService
}

func NewMockOutboundService(ctrl *gomock.Controller) *MockOutboundService {
    mock := &MockOutboundService{ctrl: ctrl}
    mock.recorder = &MockOutboundServiceMockRecorder{mock}
    return mock
}

func (m *MockOutboundService) EXPECT() *MockOutboundServiceMockRecorder {
    return m.recorder
}

type MockOutboundServiceClient struct {
    ctrl     *gomock.Controller
    recorder *MockOutboundServiceClientMockRecorder
}

type MockOutboundServiceClientMockRecorder struct {
    mock *MockOutboundServiceClient
}

func NewMockOutboundServiceClient(ctrl *gomock.Controller) *MockOutboundServiceClient {
    mock := &MockOutboundServiceClient{ctrl: ctrl}
    mock.recorder = &MockOutboundServiceClientMockRecorder{mock}
    return mock
}

func (m *MockOutboundServiceClient) EXPECT() *MockOutboundServiceClientMockRecorder {
    return m.recorder
}

func (m *MockOutboundService) OutboundList(ctx context.Context, req *connect.Request[v1.OutboundListRequest]) (*connect.Response[v1.OutboundListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OutboundList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OutboundListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockOutboundServiceMockRecorder) OutboundList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OutboundList", reflect.TypeOf((*MockOutboundService)(nil).OutboundList), ctx, req)
}

func (m *MockOutboundServiceClient) OutboundList(ctx context.Context, req *connect.Request[v1.OutboundListRequest]) (*connect.Response[v1.OutboundListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OutboundList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OutboundListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockOutboundServiceClientMockRecorder) OutboundList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OutboundList", reflect.TypeOf((*MockOutboundService)(nil).OutboundList), ctx, req)
}

func (m *MockOutboundService) OutboundDetail(ctx context.Context, req *connect.Request[v1.OutboundDetailRequest]) (*connect.Response[v1.OutboundDetailResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OutboundDetail", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OutboundDetailResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockOutboundServiceMockRecorder) OutboundDetail(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OutboundDetail", reflect.TypeOf((*MockOutboundService)(nil).OutboundDetail), ctx, req)
}

func (m *MockOutboundServiceClient) OutboundDetail(ctx context.Context, req *connect.Request[v1.OutboundDetailRequest]) (*connect.Response[v1.OutboundDetailResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OutboundDetail", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OutboundDetailResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockOutboundServiceClientMockRecorder) OutboundDetail(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OutboundDetail", reflect.TypeOf((*MockOutboundService)(nil).OutboundDetail), ctx, req)
}

func (m *MockOutboundService) OrderDetailSearch(ctx context.Context, req *connect.Request[v1.OrderDetailSearchRequest]) (*connect.Response[v1.OrderDetailSearchResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OrderDetailSearch", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OrderDetailSearchResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockOutboundServiceMockRecorder) OrderDetailSearch(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderDetailSearch", reflect.TypeOf((*MockOutboundService)(nil).OrderDetailSearch), ctx, req)
}

func (m *MockOutboundServiceClient) OrderDetailSearch(ctx context.Context, req *connect.Request[v1.OrderDetailSearchRequest]) (*connect.Response[v1.OrderDetailSearchResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OrderDetailSearch", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OrderDetailSearchResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockOutboundServiceClientMockRecorder) OrderDetailSearch(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderDetailSearch", reflect.TypeOf((*MockOutboundService)(nil).OrderDetailSearch), ctx, req)
}

