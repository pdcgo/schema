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

func (m *MockInventoryService) TransactionByIds(ctx context.Context, req *connect.Request[v1.TransactionByIdsRequest]) (*connect.Response[v1.TransactionByIdsResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TransactionByIds", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.TransactionByIdsResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceMockRecorder) TransactionByIds(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionByIds", reflect.TypeOf((*MockInventoryService)(nil).TransactionByIds), ctx, req)
}

func (m *MockInventoryServiceClient) TransactionByIds(ctx context.Context, req *connect.Request[v1.TransactionByIdsRequest]) (*connect.Response[v1.TransactionByIdsResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TransactionByIds", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.TransactionByIdsResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceClientMockRecorder) TransactionByIds(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionByIds", reflect.TypeOf((*MockInventoryService)(nil).TransactionByIds), ctx, req)
}

func (m *MockInventoryService) TransactionItemByIds(ctx context.Context, req *connect.Request[v1.TransactionItemByIdsRequest]) (*connect.Response[v1.TransactionItemByIdsResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TransactionItemByIds", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.TransactionItemByIdsResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceMockRecorder) TransactionItemByIds(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionItemByIds", reflect.TypeOf((*MockInventoryService)(nil).TransactionItemByIds), ctx, req)
}

func (m *MockInventoryServiceClient) TransactionItemByIds(ctx context.Context, req *connect.Request[v1.TransactionItemByIdsRequest]) (*connect.Response[v1.TransactionItemByIdsResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TransactionItemByIds", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.TransactionItemByIdsResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceClientMockRecorder) TransactionItemByIds(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionItemByIds", reflect.TypeOf((*MockInventoryService)(nil).TransactionItemByIds), ctx, req)
}

func (m *MockInventoryService) ProductBySkuIds(ctx context.Context, req *connect.Request[v1.ProductBySkuIdsRequest]) (*connect.Response[v1.ProductBySkuIdsResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ProductBySkuIds", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ProductBySkuIdsResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceMockRecorder) ProductBySkuIds(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductBySkuIds", reflect.TypeOf((*MockInventoryService)(nil).ProductBySkuIds), ctx, req)
}

func (m *MockInventoryServiceClient) ProductBySkuIds(ctx context.Context, req *connect.Request[v1.ProductBySkuIdsRequest]) (*connect.Response[v1.ProductBySkuIdsResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ProductBySkuIds", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ProductBySkuIdsResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceClientMockRecorder) ProductBySkuIds(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductBySkuIds", reflect.TypeOf((*MockInventoryService)(nil).ProductBySkuIds), ctx, req)
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

func (m *MockInventoryService) PlacementMove(ctx context.Context, req *connect.Request[v1.PlacementMoveRequest]) (*connect.Response[v1.PlacementMoveResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "PlacementMove", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.PlacementMoveResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceMockRecorder) PlacementMove(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PlacementMove", reflect.TypeOf((*MockInventoryService)(nil).PlacementMove), ctx, req)
}

func (m *MockInventoryServiceClient) PlacementMove(ctx context.Context, req *connect.Request[v1.PlacementMoveRequest]) (*connect.Response[v1.PlacementMoveResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "PlacementMove", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.PlacementMoveResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceClientMockRecorder) PlacementMove(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PlacementMove", reflect.TypeOf((*MockInventoryService)(nil).PlacementMove), ctx, req)
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

func (m *MockInventoryService) RackByIds(ctx context.Context, req *connect.Request[v1.RackByIdsRequest]) (*connect.Response[v1.RackByIdsResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "RackByIds", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.RackByIdsResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceMockRecorder) RackByIds(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RackByIds", reflect.TypeOf((*MockInventoryService)(nil).RackByIds), ctx, req)
}

func (m *MockInventoryServiceClient) RackByIds(ctx context.Context, req *connect.Request[v1.RackByIdsRequest]) (*connect.Response[v1.RackByIdsResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "RackByIds", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.RackByIdsResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceClientMockRecorder) RackByIds(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RackByIds", reflect.TypeOf((*MockInventoryService)(nil).RackByIds), ctx, req)
}

func (m *MockInventoryService) RackProductList(ctx context.Context, req *connect.Request[v1.RackProductListRequest]) (*connect.Response[v1.RackProductListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "RackProductList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.RackProductListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceMockRecorder) RackProductList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RackProductList", reflect.TypeOf((*MockInventoryService)(nil).RackProductList), ctx, req)
}

func (m *MockInventoryServiceClient) RackProductList(ctx context.Context, req *connect.Request[v1.RackProductListRequest]) (*connect.Response[v1.RackProductListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "RackProductList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.RackProductListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceClientMockRecorder) RackProductList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RackProductList", reflect.TypeOf((*MockInventoryService)(nil).RackProductList), ctx, req)
}

func (m *MockInventoryService) RackHistory(ctx context.Context, req *connect.Request[v1.RackHistoryRequest]) (*connect.Response[v1.RackHistoryResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "RackHistory", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.RackHistoryResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceMockRecorder) RackHistory(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RackHistory", reflect.TypeOf((*MockInventoryService)(nil).RackHistory), ctx, req)
}

func (m *MockInventoryServiceClient) RackHistory(ctx context.Context, req *connect.Request[v1.RackHistoryRequest]) (*connect.Response[v1.RackHistoryResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "RackHistory", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.RackHistoryResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceClientMockRecorder) RackHistory(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RackHistory", reflect.TypeOf((*MockInventoryService)(nil).RackHistory), ctx, req)
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

func (m *MockInventoryService) RestockCreate(ctx context.Context, req *connect.Request[v1.RestockCreateRequest]) (*connect.Response[v1.RestockCreateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "RestockCreate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.RestockCreateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceMockRecorder) RestockCreate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestockCreate", reflect.TypeOf((*MockInventoryService)(nil).RestockCreate), ctx, req)
}

func (m *MockInventoryServiceClient) RestockCreate(ctx context.Context, req *connect.Request[v1.RestockCreateRequest]) (*connect.Response[v1.RestockCreateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "RestockCreate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.RestockCreateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceClientMockRecorder) RestockCreate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestockCreate", reflect.TypeOf((*MockInventoryService)(nil).RestockCreate), ctx, req)
}

func (m *MockInventoryService) RestockUpdate(ctx context.Context, req *connect.Request[v1.RestockUpdateRequest]) (*connect.Response[v1.RestockUpdateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "RestockUpdate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.RestockUpdateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceMockRecorder) RestockUpdate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestockUpdate", reflect.TypeOf((*MockInventoryService)(nil).RestockUpdate), ctx, req)
}

func (m *MockInventoryServiceClient) RestockUpdate(ctx context.Context, req *connect.Request[v1.RestockUpdateRequest]) (*connect.Response[v1.RestockUpdateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "RestockUpdate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.RestockUpdateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceClientMockRecorder) RestockUpdate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestockUpdate", reflect.TypeOf((*MockInventoryService)(nil).RestockUpdate), ctx, req)
}

func (m *MockInventoryService) RestockDetail(ctx context.Context, req *connect.Request[v1.RestockDetailRequest]) (*connect.Response[v1.RestockDetailResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "RestockDetail", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.RestockDetailResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceMockRecorder) RestockDetail(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestockDetail", reflect.TypeOf((*MockInventoryService)(nil).RestockDetail), ctx, req)
}

func (m *MockInventoryServiceClient) RestockDetail(ctx context.Context, req *connect.Request[v1.RestockDetailRequest]) (*connect.Response[v1.RestockDetailResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "RestockDetail", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.RestockDetailResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceClientMockRecorder) RestockDetail(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestockDetail", reflect.TypeOf((*MockInventoryService)(nil).RestockDetail), ctx, req)
}

func (m *MockInventoryService) RestockList(ctx context.Context, req *connect.Request[v1.RestockListRequest]) (*connect.Response[v1.RestockListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "RestockList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.RestockListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceMockRecorder) RestockList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestockList", reflect.TypeOf((*MockInventoryService)(nil).RestockList), ctx, req)
}

func (m *MockInventoryServiceClient) RestockList(ctx context.Context, req *connect.Request[v1.RestockListRequest]) (*connect.Response[v1.RestockListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "RestockList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.RestockListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceClientMockRecorder) RestockList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestockList", reflect.TypeOf((*MockInventoryService)(nil).RestockList), ctx, req)
}

func (m *MockInventoryService) RestockLogList(ctx context.Context, req *connect.Request[v1.RestockLogListRequest]) (*connect.Response[v1.RestockLogListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "RestockLogList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.RestockLogListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceMockRecorder) RestockLogList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestockLogList", reflect.TypeOf((*MockInventoryService)(nil).RestockLogList), ctx, req)
}

func (m *MockInventoryServiceClient) RestockLogList(ctx context.Context, req *connect.Request[v1.RestockLogListRequest]) (*connect.Response[v1.RestockLogListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "RestockLogList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.RestockLogListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceClientMockRecorder) RestockLogList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestockLogList", reflect.TypeOf((*MockInventoryService)(nil).RestockLogList), ctx, req)
}

func (m *MockInventoryService) TransferCreate(ctx context.Context, req *connect.Request[v1.TransferCreateRequest]) (*connect.Response[v1.TransferCreateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TransferCreate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.TransferCreateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceMockRecorder) TransferCreate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferCreate", reflect.TypeOf((*MockInventoryService)(nil).TransferCreate), ctx, req)
}

func (m *MockInventoryServiceClient) TransferCreate(ctx context.Context, req *connect.Request[v1.TransferCreateRequest]) (*connect.Response[v1.TransferCreateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TransferCreate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.TransferCreateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceClientMockRecorder) TransferCreate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferCreate", reflect.TypeOf((*MockInventoryService)(nil).TransferCreate), ctx, req)
}

func (m *MockInventoryService) TransferCancel(ctx context.Context, req *connect.Request[v1.TransferCancelRequest]) (*connect.Response[v1.TransferCancelResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TransferCancel", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.TransferCancelResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceMockRecorder) TransferCancel(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferCancel", reflect.TypeOf((*MockInventoryService)(nil).TransferCancel), ctx, req)
}

func (m *MockInventoryServiceClient) TransferCancel(ctx context.Context, req *connect.Request[v1.TransferCancelRequest]) (*connect.Response[v1.TransferCancelResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TransferCancel", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.TransferCancelResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceClientMockRecorder) TransferCancel(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferCancel", reflect.TypeOf((*MockInventoryService)(nil).TransferCancel), ctx, req)
}

func (m *MockInventoryService) TransferAccept(ctx context.Context, req *connect.Request[v1.TransferAcceptRequest]) (*connect.Response[v1.TransferAcceptResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TransferAccept", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.TransferAcceptResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceMockRecorder) TransferAccept(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferAccept", reflect.TypeOf((*MockInventoryService)(nil).TransferAccept), ctx, req)
}

func (m *MockInventoryServiceClient) TransferAccept(ctx context.Context, req *connect.Request[v1.TransferAcceptRequest]) (*connect.Response[v1.TransferAcceptResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TransferAccept", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.TransferAcceptResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceClientMockRecorder) TransferAccept(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferAccept", reflect.TypeOf((*MockInventoryService)(nil).TransferAccept), ctx, req)
}

func (m *MockInventoryService) TransferDetail(ctx context.Context, req *connect.Request[v1.TransferDetailRequest]) (*connect.Response[v1.TransferDetailResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TransferDetail", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.TransferDetailResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceMockRecorder) TransferDetail(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferDetail", reflect.TypeOf((*MockInventoryService)(nil).TransferDetail), ctx, req)
}

func (m *MockInventoryServiceClient) TransferDetail(ctx context.Context, req *connect.Request[v1.TransferDetailRequest]) (*connect.Response[v1.TransferDetailResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TransferDetail", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.TransferDetailResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceClientMockRecorder) TransferDetail(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferDetail", reflect.TypeOf((*MockInventoryService)(nil).TransferDetail), ctx, req)
}

func (m *MockInventoryService) TransferList(ctx context.Context, req *connect.Request[v1.TransferListRequest]) (*connect.Response[v1.TransferListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TransferList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.TransferListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceMockRecorder) TransferList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferList", reflect.TypeOf((*MockInventoryService)(nil).TransferList), ctx, req)
}

func (m *MockInventoryServiceClient) TransferList(ctx context.Context, req *connect.Request[v1.TransferListRequest]) (*connect.Response[v1.TransferListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TransferList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.TransferListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceClientMockRecorder) TransferList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferList", reflect.TypeOf((*MockInventoryService)(nil).TransferList), ctx, req)
}

func (m *MockInventoryService) OpnameCreate(ctx context.Context, req *connect.Request[v1.OpnameCreateRequest]) (*connect.Response[v1.OpnameCreateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OpnameCreate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OpnameCreateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceMockRecorder) OpnameCreate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpnameCreate", reflect.TypeOf((*MockInventoryService)(nil).OpnameCreate), ctx, req)
}

func (m *MockInventoryServiceClient) OpnameCreate(ctx context.Context, req *connect.Request[v1.OpnameCreateRequest]) (*connect.Response[v1.OpnameCreateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OpnameCreate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OpnameCreateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceClientMockRecorder) OpnameCreate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpnameCreate", reflect.TypeOf((*MockInventoryService)(nil).OpnameCreate), ctx, req)
}

func (m *MockInventoryService) OpnameList(ctx context.Context, req *connect.Request[v1.OpnameListRequest]) (*connect.Response[v1.OpnameListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OpnameList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OpnameListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceMockRecorder) OpnameList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpnameList", reflect.TypeOf((*MockInventoryService)(nil).OpnameList), ctx, req)
}

func (m *MockInventoryServiceClient) OpnameList(ctx context.Context, req *connect.Request[v1.OpnameListRequest]) (*connect.Response[v1.OpnameListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OpnameList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OpnameListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceClientMockRecorder) OpnameList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpnameList", reflect.TypeOf((*MockInventoryService)(nil).OpnameList), ctx, req)
}

func (m *MockInventoryService) OpnameDetail(ctx context.Context, req *connect.Request[v1.OpnameDetailRequest]) (*connect.Response[v1.OpnameDetailResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OpnameDetail", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OpnameDetailResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceMockRecorder) OpnameDetail(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpnameDetail", reflect.TypeOf((*MockInventoryService)(nil).OpnameDetail), ctx, req)
}

func (m *MockInventoryServiceClient) OpnameDetail(ctx context.Context, req *connect.Request[v1.OpnameDetailRequest]) (*connect.Response[v1.OpnameDetailResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OpnameDetail", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OpnameDetailResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceClientMockRecorder) OpnameDetail(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpnameDetail", reflect.TypeOf((*MockInventoryService)(nil).OpnameDetail), ctx, req)
}

func (m *MockInventoryService) OpnameLineCount(ctx context.Context, req *connect.Request[v1.OpnameLineCountRequest]) (*connect.Response[v1.OpnameLineCountResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OpnameLineCount", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OpnameLineCountResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceMockRecorder) OpnameLineCount(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpnameLineCount", reflect.TypeOf((*MockInventoryService)(nil).OpnameLineCount), ctx, req)
}

func (m *MockInventoryServiceClient) OpnameLineCount(ctx context.Context, req *connect.Request[v1.OpnameLineCountRequest]) (*connect.Response[v1.OpnameLineCountResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OpnameLineCount", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OpnameLineCountResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceClientMockRecorder) OpnameLineCount(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpnameLineCount", reflect.TypeOf((*MockInventoryService)(nil).OpnameLineCount), ctx, req)
}

func (m *MockInventoryService) OpnameComplete(ctx context.Context, req *connect.Request[v1.OpnameCompleteRequest]) (*connect.Response[v1.OpnameCompleteResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OpnameComplete", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OpnameCompleteResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceMockRecorder) OpnameComplete(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpnameComplete", reflect.TypeOf((*MockInventoryService)(nil).OpnameComplete), ctx, req)
}

func (m *MockInventoryServiceClient) OpnameComplete(ctx context.Context, req *connect.Request[v1.OpnameCompleteRequest]) (*connect.Response[v1.OpnameCompleteResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OpnameComplete", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OpnameCompleteResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceClientMockRecorder) OpnameComplete(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpnameComplete", reflect.TypeOf((*MockInventoryService)(nil).OpnameComplete), ctx, req)
}

func (m *MockInventoryService) OpnameCancel(ctx context.Context, req *connect.Request[v1.OpnameCancelRequest]) (*connect.Response[v1.OpnameCancelResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OpnameCancel", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OpnameCancelResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceMockRecorder) OpnameCancel(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpnameCancel", reflect.TypeOf((*MockInventoryService)(nil).OpnameCancel), ctx, req)
}

func (m *MockInventoryServiceClient) OpnameCancel(ctx context.Context, req *connect.Request[v1.OpnameCancelRequest]) (*connect.Response[v1.OpnameCancelResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OpnameCancel", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OpnameCancelResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceClientMockRecorder) OpnameCancel(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpnameCancel", reflect.TypeOf((*MockInventoryService)(nil).OpnameCancel), ctx, req)
}

