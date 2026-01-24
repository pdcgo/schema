package warehouse_iface_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/warehouse_iface/v1"
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

func (m *MockInventoryService) Placements(ctx context.Context, req *connect.Request[v1.PlacementsRequest]) (*connect.Response[v1.PlacementsResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "Placements", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.PlacementsResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceMockRecorder) Placements(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Placements", reflect.TypeOf((*MockInventoryService)(nil).Placements), ctx, req)
}

func (m *MockInventoryServiceClient) Placements(ctx context.Context, req *connect.Request[v1.PlacementsRequest]) (*connect.Response[v1.PlacementsResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "Placements", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.PlacementsResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceClientMockRecorder) Placements(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Placements", reflect.TypeOf((*MockInventoryService)(nil).Placements), ctx, req)
}

func (m *MockInventoryService) PlacementsIDs(ctx context.Context, req *connect.Request[v1.PlacementsIDsRequest]) (*connect.Response[v1.PlacementsIDsResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "PlacementsIDs", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.PlacementsIDsResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceMockRecorder) PlacementsIDs(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PlacementsIDs", reflect.TypeOf((*MockInventoryService)(nil).PlacementsIDs), ctx, req)
}

func (m *MockInventoryServiceClient) PlacementsIDs(ctx context.Context, req *connect.Request[v1.PlacementsIDsRequest]) (*connect.Response[v1.PlacementsIDsResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "PlacementsIDs", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.PlacementsIDsResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceClientMockRecorder) PlacementsIDs(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PlacementsIDs", reflect.TypeOf((*MockInventoryService)(nil).PlacementsIDs), ctx, req)
}

func (m *MockInventoryService) BlacklistedSku(ctx context.Context, req *connect.Request[v1.BlacklistedSkuRequest]) (*connect.Response[v1.BlacklistedSkuResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "BlacklistedSku", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.BlacklistedSkuResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceMockRecorder) BlacklistedSku(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlacklistedSku", reflect.TypeOf((*MockInventoryService)(nil).BlacklistedSku), ctx, req)
}

func (m *MockInventoryServiceClient) BlacklistedSku(ctx context.Context, req *connect.Request[v1.BlacklistedSkuRequest]) (*connect.Response[v1.BlacklistedSkuResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "BlacklistedSku", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.BlacklistedSkuResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceClientMockRecorder) BlacklistedSku(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlacklistedSku", reflect.TypeOf((*MockInventoryService)(nil).BlacklistedSku), ctx, req)
}

func (m *MockInventoryService) BlacklistedSkuAdd(ctx context.Context, req *connect.Request[v1.BlacklistedSkuAddRequest]) (*connect.Response[v1.BlacklistedSkuAddResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "BlacklistedSkuAdd", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.BlacklistedSkuAddResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceMockRecorder) BlacklistedSkuAdd(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlacklistedSkuAdd", reflect.TypeOf((*MockInventoryService)(nil).BlacklistedSkuAdd), ctx, req)
}

func (m *MockInventoryServiceClient) BlacklistedSkuAdd(ctx context.Context, req *connect.Request[v1.BlacklistedSkuAddRequest]) (*connect.Response[v1.BlacklistedSkuAddResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "BlacklistedSkuAdd", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.BlacklistedSkuAddResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceClientMockRecorder) BlacklistedSkuAdd(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlacklistedSkuAdd", reflect.TypeOf((*MockInventoryService)(nil).BlacklistedSkuAdd), ctx, req)
}

func (m *MockInventoryService) BlacklistedSkuRemove(ctx context.Context, req *connect.Request[v1.BlacklistedSkuRemoveRequest]) (*connect.Response[v1.BlacklistedSkuRemoveResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "BlacklistedSkuRemove", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.BlacklistedSkuRemoveResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceMockRecorder) BlacklistedSkuRemove(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlacklistedSkuRemove", reflect.TypeOf((*MockInventoryService)(nil).BlacklistedSkuRemove), ctx, req)
}

func (m *MockInventoryServiceClient) BlacklistedSkuRemove(ctx context.Context, req *connect.Request[v1.BlacklistedSkuRemoveRequest]) (*connect.Response[v1.BlacklistedSkuRemoveResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "BlacklistedSkuRemove", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.BlacklistedSkuRemoveResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceClientMockRecorder) BlacklistedSkuRemove(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlacklistedSkuRemove", reflect.TypeOf((*MockInventoryService)(nil).BlacklistedSkuRemove), ctx, req)
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

func (m *MockInventoryService) ProductHistory(ctx context.Context, req *connect.Request[v1.ProductHistoryRequest]) (*connect.Response[v1.ProductHistoryResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ProductHistory", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ProductHistoryResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceMockRecorder) ProductHistory(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductHistory", reflect.TypeOf((*MockInventoryService)(nil).ProductHistory), ctx, req)
}

func (m *MockInventoryServiceClient) ProductHistory(ctx context.Context, req *connect.Request[v1.ProductHistoryRequest]) (*connect.Response[v1.ProductHistoryResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ProductHistory", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ProductHistoryResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceClientMockRecorder) ProductHistory(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductHistory", reflect.TypeOf((*MockInventoryService)(nil).ProductHistory), ctx, req)
}

