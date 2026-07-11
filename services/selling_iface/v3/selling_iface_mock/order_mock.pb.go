package selling_iface_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v3 "github.com/pdcgo/schema/services/selling_iface/v3"
)

type MockOrderService struct {
    ctrl     *gomock.Controller
    recorder *MockOrderServiceMockRecorder
}

type MockOrderServiceMockRecorder struct {
    mock *MockOrderService
}

func NewMockOrderService(ctrl *gomock.Controller) *MockOrderService {
    mock := &MockOrderService{ctrl: ctrl}
    mock.recorder = &MockOrderServiceMockRecorder{mock}
    return mock
}

func (m *MockOrderService) EXPECT() *MockOrderServiceMockRecorder {
    return m.recorder
}

type MockOrderServiceClient struct {
    ctrl     *gomock.Controller
    recorder *MockOrderServiceClientMockRecorder
}

type MockOrderServiceClientMockRecorder struct {
    mock *MockOrderServiceClient
}

func NewMockOrderServiceClient(ctrl *gomock.Controller) *MockOrderServiceClient {
    mock := &MockOrderServiceClient{ctrl: ctrl}
    mock.recorder = &MockOrderServiceClientMockRecorder{mock}
    return mock
}

func (m *MockOrderServiceClient) EXPECT() *MockOrderServiceClientMockRecorder {
    return m.recorder
}

func (m *MockOrderService) OrderCreate(ctx context.Context, req *connect.Request[v3.OrderCreateRequest]) (*connect.Response[v3.OrderCreateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OrderCreate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v3.OrderCreateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockOrderServiceMockRecorder) OrderCreate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderCreate", reflect.TypeOf((*MockOrderService)(nil).OrderCreate), ctx, req)
}

func (m *MockOrderServiceClient) OrderCreate(ctx context.Context, req *connect.Request[v3.OrderCreateRequest]) (*connect.Response[v3.OrderCreateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OrderCreate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v3.OrderCreateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockOrderServiceClientMockRecorder) OrderCreate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderCreate", reflect.TypeOf((*MockOrderService)(nil).OrderCreate), ctx, req)
}

func (m *MockOrderService) OrderList(ctx context.Context, req *connect.Request[v3.OrderListRequest]) (*connect.Response[v3.OrderListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OrderList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v3.OrderListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockOrderServiceMockRecorder) OrderList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderList", reflect.TypeOf((*MockOrderService)(nil).OrderList), ctx, req)
}

func (m *MockOrderServiceClient) OrderList(ctx context.Context, req *connect.Request[v3.OrderListRequest]) (*connect.Response[v3.OrderListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OrderList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v3.OrderListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockOrderServiceClientMockRecorder) OrderList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderList", reflect.TypeOf((*MockOrderService)(nil).OrderList), ctx, req)
}

func (m *MockOrderService) OrderDetail(ctx context.Context, req *connect.Request[v3.OrderDetailRequest]) (*connect.Response[v3.OrderDetailResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OrderDetail", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v3.OrderDetailResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockOrderServiceMockRecorder) OrderDetail(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderDetail", reflect.TypeOf((*MockOrderService)(nil).OrderDetail), ctx, req)
}

func (m *MockOrderServiceClient) OrderDetail(ctx context.Context, req *connect.Request[v3.OrderDetailRequest]) (*connect.Response[v3.OrderDetailResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OrderDetail", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v3.OrderDetailResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockOrderServiceClientMockRecorder) OrderDetail(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderDetail", reflect.TypeOf((*MockOrderService)(nil).OrderDetail), ctx, req)
}

func (m *MockOrderService) OrderCancel(ctx context.Context, req *connect.Request[v3.OrderCancelRequest]) (*connect.Response[v3.OrderCancelResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OrderCancel", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v3.OrderCancelResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockOrderServiceMockRecorder) OrderCancel(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderCancel", reflect.TypeOf((*MockOrderService)(nil).OrderCancel), ctx, req)
}

func (m *MockOrderServiceClient) OrderCancel(ctx context.Context, req *connect.Request[v3.OrderCancelRequest]) (*connect.Response[v3.OrderCancelResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OrderCancel", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v3.OrderCancelResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockOrderServiceClientMockRecorder) OrderCancel(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderCancel", reflect.TypeOf((*MockOrderService)(nil).OrderCancel), ctx, req)
}

