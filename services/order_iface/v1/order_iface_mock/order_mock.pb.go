package order_iface_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/order_iface/v1"
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

func (m *MockOrderService) OrderFundSet(ctx context.Context, stream *connect.ClientStream[v1.OrderFundSetRequest]) (*connect.Response[v1.OrderFundSetResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OrderFundSet", ctx, stream)
    ret0, _ := ret[0].(*connect.Response[v1.OrderFundSetResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockOrderServiceMockRecorder) OrderFundSet(ctx, stream interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderFundSet", reflect.TypeOf((*MockOrderService)(nil).OrderFundSet), ctx, stream)
}

func (m *MockOrderService) OrderTagRemove(ctx context.Context, req *connect.Request[v1.OrderTagRemoveRequest]) (*connect.Response[v1.OrderTagRemoveResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OrderTagRemove", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OrderTagRemoveResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockOrderServiceMockRecorder) OrderTagRemove(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderTagRemove", reflect.TypeOf((*MockOrderService)(nil).OrderTagRemove), ctx, req)
}

func (m *MockOrderService) OrderTagAdd(ctx context.Context, req *connect.Request[v1.OrderTagAddRequest]) (*connect.Response[v1.OrderTagAddResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OrderTagAdd", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OrderTagAddResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockOrderServiceMockRecorder) OrderTagAdd(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderTagAdd", reflect.TypeOf((*MockOrderService)(nil).OrderTagAdd), ctx, req)
}

func (m *MockOrderService) OrderList(ctx context.Context, req *connect.Request[v1.OrderListRequest], stream *connect.ServerStream[v1.OrderListResponse]) error {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OrderList", ctx, req, stream)
    ret0, _ := ret[0].(error)
    return ret0
}

func (mr *MockOrderServiceMockRecorder) OrderList(ctx, req, stream interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderList", reflect.TypeOf((*MockOrderService)(nil).OrderList), ctx, req, stream)
}

func (m *MockOrderService) OrderOverview(ctx context.Context, req *connect.Request[v1.OrderOverviewRequest]) (*connect.Response[v1.OrderOverviewResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OrderOverview", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OrderOverviewResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockOrderServiceMockRecorder) OrderOverview(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderOverview", reflect.TypeOf((*MockOrderService)(nil).OrderOverview), ctx, req)
}

