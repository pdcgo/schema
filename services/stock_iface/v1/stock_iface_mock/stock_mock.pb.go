package stock_iface_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/stock_iface/v1"
)

type MockStockService struct {
    ctrl     *gomock.Controller
    recorder *MockStockServiceMockRecorder
}

type MockStockServiceMockRecorder struct {
    mock *MockStockService
}

func NewMockStockService(ctrl *gomock.Controller) *MockStockService {
    mock := &MockStockService{ctrl: ctrl}
    mock.recorder = &MockStockServiceMockRecorder{mock}
    return mock
}

func (m *MockStockService) EXPECT() *MockStockServiceMockRecorder {
    return m.recorder
}

func (m *MockStockService) InboundCreate(ctx context.Context, req *connect.Request[v1.InboundCreateRequest]) (*connect.Response[v1.InboundCreateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "InboundCreate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.InboundCreateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockStockServiceMockRecorder) InboundCreate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InboundCreate", reflect.TypeOf((*MockStockService)(nil).InboundCreate), ctx, req)
}

func (m *MockStockService) InboundUpdate(ctx context.Context, req *connect.Request[v1.InboundUpdateRequest]) (*connect.Response[v1.InboundUpdateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "InboundUpdate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.InboundUpdateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockStockServiceMockRecorder) InboundUpdate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InboundUpdate", reflect.TypeOf((*MockStockService)(nil).InboundUpdate), ctx, req)
}

func (m *MockStockService) InboundAccept(ctx context.Context, req *connect.Request[v1.InboundAcceptRequest]) (*connect.Response[v1.InboundAcceptResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "InboundAccept", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.InboundAcceptResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockStockServiceMockRecorder) InboundAccept(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InboundAccept", reflect.TypeOf((*MockStockService)(nil).InboundAccept), ctx, req)
}

func (m *MockStockService) StockAdjustment(ctx context.Context, req *connect.Request[v1.StockAdjustmentRequest]) (*connect.Response[v1.StockAdjustmentResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "StockAdjustment", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.StockAdjustmentResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockStockServiceMockRecorder) StockAdjustment(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StockAdjustment", reflect.TypeOf((*MockStockService)(nil).StockAdjustment), ctx, req)
}

func (m *MockStockService) TransferToWarehouse(ctx context.Context, req *connect.Request[v1.TransferToWarehouseRequest]) (*connect.Response[v1.TransferToWarehouseResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TransferToWarehouse", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.TransferToWarehouseResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockStockServiceMockRecorder) TransferToWarehouse(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferToWarehouse", reflect.TypeOf((*MockStockService)(nil).TransferToWarehouse), ctx, req)
}

func (m *MockStockService) TransferToWarehouseAccept(ctx context.Context, req *connect.Request[v1.TransferToWarehouseAcceptRequest]) (*connect.Response[v1.TransferToWarehouseAcceptResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TransferToWarehouseAccept", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.TransferToWarehouseAcceptResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockStockServiceMockRecorder) TransferToWarehouseAccept(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferToWarehouseAccept", reflect.TypeOf((*MockStockService)(nil).TransferToWarehouseAccept), ctx, req)
}

func (m *MockStockService) TransferToWarehouseCancel(ctx context.Context, req *connect.Request[v1.TransferToWarehouseCancelRequest]) (*connect.Response[v1.TransferToWarehouseCancelResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TransferToWarehouseCancel", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.TransferToWarehouseCancelResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockStockServiceMockRecorder) TransferToWarehouseCancel(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferToWarehouseCancel", reflect.TypeOf((*MockStockService)(nil).TransferToWarehouseCancel), ctx, req)
}

func (m *MockStockService) StockProblemCreate(ctx context.Context, req *connect.Request[v1.StockProblemCreateRequest]) (*connect.Response[v1.StockProblemCreateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "StockProblemCreate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.StockProblemCreateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockStockServiceMockRecorder) StockProblemCreate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StockProblemCreate", reflect.TypeOf((*MockStockService)(nil).StockProblemCreate), ctx, req)
}

