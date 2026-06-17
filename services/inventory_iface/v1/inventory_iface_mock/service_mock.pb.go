package inventory_iface_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/inventory_iface/v1"
)

type MockInventoryService struct {
    ctrl     *gomock.Controller
    recorder *MockInventoryServiceMockRecorder
}

type MockInventoryServiceMockRecorder struct {
    mock *MockInventoryService
}

func NewMockInventoryService(ctrl *gomock.Controller) *MockInventoryService {
    mock := &MockInventoryService{ctrl: ctrl}
    mock.recorder = &MockInventoryServiceMockRecorder{mock}
    return mock
}

func (m *MockInventoryService) EXPECT() *MockInventoryServiceMockRecorder {
    return m.recorder
}

type MockInventoryServiceClient struct {
    ctrl     *gomock.Controller
    recorder *MockInventoryServiceClientMockRecorder
}

type MockInventoryServiceClientMockRecorder struct {
    mock *MockInventoryServiceClient
}

func NewMockInventoryServiceClient(ctrl *gomock.Controller) *MockInventoryServiceClient {
    mock := &MockInventoryServiceClient{ctrl: ctrl}
    mock.recorder = &MockInventoryServiceClientMockRecorder{mock}
    return mock
}

func (m *MockInventoryServiceClient) EXPECT() *MockInventoryServiceClientMockRecorder {
    return m.recorder
}

func (m *MockInventoryService) Order(ctx context.Context, req *connect.Request[v1.OrderRequest]) (*connect.Response[v1.OrderResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "Order", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OrderResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceMockRecorder) Order(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Order", reflect.TypeOf((*MockInventoryService)(nil).Order), ctx, req)
}

func (m *MockInventoryServiceClient) Order(ctx context.Context, req *connect.Request[v1.OrderRequest]) (*connect.Response[v1.OrderResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "Order", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OrderResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceClientMockRecorder) Order(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Order", reflect.TypeOf((*MockInventoryService)(nil).Order), ctx, req)
}

func (m *MockInventoryService) StockMovement(ctx context.Context, req *connect.Request[v1.StockMovementRequest]) (*connect.Response[v1.StockMovementResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "StockMovement", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.StockMovementResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceMockRecorder) StockMovement(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StockMovement", reflect.TypeOf((*MockInventoryService)(nil).StockMovement), ctx, req)
}

func (m *MockInventoryServiceClient) StockMovement(ctx context.Context, req *connect.Request[v1.StockMovementRequest]) (*connect.Response[v1.StockMovementResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "StockMovement", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.StockMovementResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceClientMockRecorder) StockMovement(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StockMovement", reflect.TypeOf((*MockInventoryService)(nil).StockMovement), ctx, req)
}

func (m *MockInventoryService) PushStockEvent(ctx context.Context, req *connect.Request[v1.PushStockEventRequest]) (*connect.Response[v1.PushStockEventResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "PushStockEvent", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.PushStockEventResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceMockRecorder) PushStockEvent(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PushStockEvent", reflect.TypeOf((*MockInventoryService)(nil).PushStockEvent), ctx, req)
}

func (m *MockInventoryServiceClient) PushStockEvent(ctx context.Context, req *connect.Request[v1.PushStockEventRequest]) (*connect.Response[v1.PushStockEventResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "PushStockEvent", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.PushStockEventResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceClientMockRecorder) PushStockEvent(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PushStockEvent", reflect.TypeOf((*MockInventoryService)(nil).PushStockEvent), ctx, req)
}

func (m *MockInventoryService) ProductPlacementList(ctx context.Context, req *connect.Request[v1.ProductPlacementListRequest]) (*connect.Response[v1.ProductPlacementListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ProductPlacementList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ProductPlacementListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceMockRecorder) ProductPlacementList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductPlacementList", reflect.TypeOf((*MockInventoryService)(nil).ProductPlacementList), ctx, req)
}

func (m *MockInventoryServiceClient) ProductPlacementList(ctx context.Context, req *connect.Request[v1.ProductPlacementListRequest]) (*connect.Response[v1.ProductPlacementListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ProductPlacementList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ProductPlacementListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceClientMockRecorder) ProductPlacementList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductPlacementList", reflect.TypeOf((*MockInventoryService)(nil).ProductPlacementList), ctx, req)
}

func (m *MockInventoryService) ProductPlacementLog(ctx context.Context, req *connect.Request[v1.ProductPlacementLogRequest]) (*connect.Response[v1.ProductPlacementLogResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ProductPlacementLog", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ProductPlacementLogResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceMockRecorder) ProductPlacementLog(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductPlacementLog", reflect.TypeOf((*MockInventoryService)(nil).ProductPlacementLog), ctx, req)
}

func (m *MockInventoryServiceClient) ProductPlacementLog(ctx context.Context, req *connect.Request[v1.ProductPlacementLogRequest]) (*connect.Response[v1.ProductPlacementLogResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ProductPlacementLog", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ProductPlacementLogResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceClientMockRecorder) ProductPlacementLog(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductPlacementLog", reflect.TypeOf((*MockInventoryService)(nil).ProductPlacementLog), ctx, req)
}

func (m *MockInventoryService) ProductBatchList(ctx context.Context, req *connect.Request[v1.ProductBatchListRequest]) (*connect.Response[v1.ProductBatchListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ProductBatchList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ProductBatchListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceMockRecorder) ProductBatchList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductBatchList", reflect.TypeOf((*MockInventoryService)(nil).ProductBatchList), ctx, req)
}

func (m *MockInventoryServiceClient) ProductBatchList(ctx context.Context, req *connect.Request[v1.ProductBatchListRequest]) (*connect.Response[v1.ProductBatchListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ProductBatchList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ProductBatchListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceClientMockRecorder) ProductBatchList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductBatchList", reflect.TypeOf((*MockInventoryService)(nil).ProductBatchList), ctx, req)
}

func (m *MockInventoryService) ProductReconcile(ctx context.Context, req *connect.Request[v1.ProductReconcileRequest]) (*connect.Response[v1.ProductReconcileResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ProductReconcile", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ProductReconcileResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceMockRecorder) ProductReconcile(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductReconcile", reflect.TypeOf((*MockInventoryService)(nil).ProductReconcile), ctx, req)
}

func (m *MockInventoryServiceClient) ProductReconcile(ctx context.Context, req *connect.Request[v1.ProductReconcileRequest]) (*connect.Response[v1.ProductReconcileResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ProductReconcile", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ProductReconcileResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceClientMockRecorder) ProductReconcile(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductReconcile", reflect.TypeOf((*MockInventoryService)(nil).ProductReconcile), ctx, req)
}

