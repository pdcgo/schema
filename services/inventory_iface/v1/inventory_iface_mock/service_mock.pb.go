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

func (m *MockInventoryService) TransactionCreate(ctx context.Context, req *connect.Request[v1.TransactionCreateRequest]) (*connect.Response[v1.TransactionCreateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TransactionCreate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.TransactionCreateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceMockRecorder) TransactionCreate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionCreate", reflect.TypeOf((*MockInventoryService)(nil).TransactionCreate), ctx, req)
}

func (m *MockInventoryServiceClient) TransactionCreate(ctx context.Context, req *connect.Request[v1.TransactionCreateRequest]) (*connect.Response[v1.TransactionCreateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TransactionCreate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.TransactionCreateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceClientMockRecorder) TransactionCreate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionCreate", reflect.TypeOf((*MockInventoryService)(nil).TransactionCreate), ctx, req)
}

func (m *MockInventoryService) TransactionCancel(ctx context.Context, req *connect.Request[v1.TransactionCancelRequest]) (*connect.Response[v1.TransactionCancelResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TransactionCancel", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.TransactionCancelResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceMockRecorder) TransactionCancel(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionCancel", reflect.TypeOf((*MockInventoryService)(nil).TransactionCancel), ctx, req)
}

func (m *MockInventoryServiceClient) TransactionCancel(ctx context.Context, req *connect.Request[v1.TransactionCancelRequest]) (*connect.Response[v1.TransactionCancelResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TransactionCancel", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.TransactionCancelResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceClientMockRecorder) TransactionCancel(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionCancel", reflect.TypeOf((*MockInventoryService)(nil).TransactionCancel), ctx, req)
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

func (m *MockInventoryService) ProductReconcile(ctx context.Context, req *connect.Request[v1.ProductReconcileRequest], stream *connect.ServerStream[v1.ProductReconcileResponse]) error {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ProductReconcile", ctx, req, stream)
    ret0, _ := ret[0].(error)
    return ret0
}

func (mr *MockInventoryServiceMockRecorder) ProductReconcile(ctx, req, stream interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductReconcile", reflect.TypeOf((*MockInventoryService)(nil).ProductReconcile), ctx, req, stream)
}

func (m *MockInventoryServiceClient) ProductReconcile(ctx context.Context, req *connect.Request[v1.ProductReconcileRequest]) (*connect.ServerStreamForClient[v1.ProductReconcileResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ProductReconcile", ctx, req)
    ret0, _ := ret[0].(*connect.ServerStreamForClient[v1.ProductReconcileResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceClientMockRecorder) ProductReconcile(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductReconcile", reflect.TypeOf((*MockInventoryService)(nil).ProductReconcile), ctx, req)
}

func (m *MockInventoryService) RackCreate(ctx context.Context, req *connect.Request[v1.RackCreateRequest]) (*connect.Response[v1.RackCreateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "RackCreate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.RackCreateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceMockRecorder) RackCreate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RackCreate", reflect.TypeOf((*MockInventoryService)(nil).RackCreate), ctx, req)
}

func (m *MockInventoryServiceClient) RackCreate(ctx context.Context, req *connect.Request[v1.RackCreateRequest]) (*connect.Response[v1.RackCreateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "RackCreate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.RackCreateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceClientMockRecorder) RackCreate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RackCreate", reflect.TypeOf((*MockInventoryService)(nil).RackCreate), ctx, req)
}

func (m *MockInventoryService) RackUpdate(ctx context.Context, req *connect.Request[v1.RackUpdateRequest]) (*connect.Response[v1.RackUpdateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "RackUpdate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.RackUpdateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceMockRecorder) RackUpdate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RackUpdate", reflect.TypeOf((*MockInventoryService)(nil).RackUpdate), ctx, req)
}

func (m *MockInventoryServiceClient) RackUpdate(ctx context.Context, req *connect.Request[v1.RackUpdateRequest]) (*connect.Response[v1.RackUpdateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "RackUpdate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.RackUpdateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceClientMockRecorder) RackUpdate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RackUpdate", reflect.TypeOf((*MockInventoryService)(nil).RackUpdate), ctx, req)
}

func (m *MockInventoryService) RackDelete(ctx context.Context, req *connect.Request[v1.RackDeleteRequest]) (*connect.Response[v1.RackDeleteResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "RackDelete", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.RackDeleteResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceMockRecorder) RackDelete(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RackDelete", reflect.TypeOf((*MockInventoryService)(nil).RackDelete), ctx, req)
}

func (m *MockInventoryServiceClient) RackDelete(ctx context.Context, req *connect.Request[v1.RackDeleteRequest]) (*connect.Response[v1.RackDeleteResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "RackDelete", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.RackDeleteResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceClientMockRecorder) RackDelete(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RackDelete", reflect.TypeOf((*MockInventoryService)(nil).RackDelete), ctx, req)
}

func (m *MockInventoryService) RackDetail(ctx context.Context, req *connect.Request[v1.RackDetailRequest]) (*connect.Response[v1.RackDetailResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "RackDetail", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.RackDetailResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceMockRecorder) RackDetail(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RackDetail", reflect.TypeOf((*MockInventoryService)(nil).RackDetail), ctx, req)
}

func (m *MockInventoryServiceClient) RackDetail(ctx context.Context, req *connect.Request[v1.RackDetailRequest]) (*connect.Response[v1.RackDetailResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "RackDetail", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.RackDetailResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceClientMockRecorder) RackDetail(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RackDetail", reflect.TypeOf((*MockInventoryService)(nil).RackDetail), ctx, req)
}

func (m *MockInventoryService) RackList(ctx context.Context, req *connect.Request[v1.RackListRequest]) (*connect.Response[v1.RackListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "RackList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.RackListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceMockRecorder) RackList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RackList", reflect.TypeOf((*MockInventoryService)(nil).RackList), ctx, req)
}

func (m *MockInventoryServiceClient) RackList(ctx context.Context, req *connect.Request[v1.RackListRequest]) (*connect.Response[v1.RackListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "RackList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.RackListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceClientMockRecorder) RackList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RackList", reflect.TypeOf((*MockInventoryService)(nil).RackList), ctx, req)
}

func (m *MockInventoryService) ProductConfig(ctx context.Context, req *connect.Request[v1.ProductConfigRequest]) (*connect.Response[v1.ProductConfigResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ProductConfig", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ProductConfigResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceMockRecorder) ProductConfig(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductConfig", reflect.TypeOf((*MockInventoryService)(nil).ProductConfig), ctx, req)
}

func (m *MockInventoryServiceClient) ProductConfig(ctx context.Context, req *connect.Request[v1.ProductConfigRequest]) (*connect.Response[v1.ProductConfigResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ProductConfig", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ProductConfigResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceClientMockRecorder) ProductConfig(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductConfig", reflect.TypeOf((*MockInventoryService)(nil).ProductConfig), ctx, req)
}

func (m *MockInventoryService) ProductConfigUpdate(ctx context.Context, req *connect.Request[v1.ProductConfigUpdateRequest]) (*connect.Response[v1.ProductConfigUpdateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ProductConfigUpdate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ProductConfigUpdateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceMockRecorder) ProductConfigUpdate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductConfigUpdate", reflect.TypeOf((*MockInventoryService)(nil).ProductConfigUpdate), ctx, req)
}

func (m *MockInventoryServiceClient) ProductConfigUpdate(ctx context.Context, req *connect.Request[v1.ProductConfigUpdateRequest]) (*connect.Response[v1.ProductConfigUpdateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ProductConfigUpdate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ProductConfigUpdateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceClientMockRecorder) ProductConfigUpdate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductConfigUpdate", reflect.TypeOf((*MockInventoryService)(nil).ProductConfigUpdate), ctx, req)
}

func (m *MockInventoryService) ProductList(ctx context.Context, req *connect.Request[v1.ProductListRequest]) (*connect.Response[v1.ProductListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ProductList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ProductListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceMockRecorder) ProductList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductList", reflect.TypeOf((*MockInventoryService)(nil).ProductList), ctx, req)
}

func (m *MockInventoryServiceClient) ProductList(ctx context.Context, req *connect.Request[v1.ProductListRequest]) (*connect.Response[v1.ProductListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ProductList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ProductListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceClientMockRecorder) ProductList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductList", reflect.TypeOf((*MockInventoryService)(nil).ProductList), ctx, req)
}

func (m *MockInventoryService) ProductDetail(ctx context.Context, req *connect.Request[v1.ProductDetailRequest]) (*connect.Response[v1.ProductDetailResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ProductDetail", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ProductDetailResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceMockRecorder) ProductDetail(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductDetail", reflect.TypeOf((*MockInventoryService)(nil).ProductDetail), ctx, req)
}

func (m *MockInventoryServiceClient) ProductDetail(ctx context.Context, req *connect.Request[v1.ProductDetailRequest]) (*connect.Response[v1.ProductDetailResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ProductDetail", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ProductDetailResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceClientMockRecorder) ProductDetail(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductDetail", reflect.TypeOf((*MockInventoryService)(nil).ProductDetail), ctx, req)
}

