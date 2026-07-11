package product_iface_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/product_iface/v1"
)

type MockProductService struct {
    ctrl     *gomock.Controller
    recorder *MockProductServiceMockRecorder
}

type MockProductServiceMockRecorder struct {
    mock *MockProductService
}

func NewMockProductService(ctrl *gomock.Controller) *MockProductService {
    mock := &MockProductService{ctrl: ctrl}
    mock.recorder = &MockProductServiceMockRecorder{mock}
    return mock
}

func (m *MockProductService) EXPECT() *MockProductServiceMockRecorder {
    return m.recorder
}

type MockProductServiceClient struct {
    ctrl     *gomock.Controller
    recorder *MockProductServiceClientMockRecorder
}

type MockProductServiceClientMockRecorder struct {
    mock *MockProductServiceClient
}

func NewMockProductServiceClient(ctrl *gomock.Controller) *MockProductServiceClient {
    mock := &MockProductServiceClient{ctrl: ctrl}
    mock.recorder = &MockProductServiceClientMockRecorder{mock}
    return mock
}

func (m *MockProductServiceClient) EXPECT() *MockProductServiceClientMockRecorder {
    return m.recorder
}

func (m *MockProductService) ProductDuplicate(ctx context.Context, req *connect.Request[v1.ProductDuplicateRequest]) (*connect.Response[v1.ProductDuplicateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ProductDuplicate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ProductDuplicateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockProductServiceMockRecorder) ProductDuplicate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductDuplicate", reflect.TypeOf((*MockProductService)(nil).ProductDuplicate), ctx, req)
}

func (m *MockProductServiceClient) ProductDuplicate(ctx context.Context, req *connect.Request[v1.ProductDuplicateRequest]) (*connect.Response[v1.ProductDuplicateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ProductDuplicate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ProductDuplicateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockProductServiceClientMockRecorder) ProductDuplicate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductDuplicate", reflect.TypeOf((*MockProductService)(nil).ProductDuplicate), ctx, req)
}

func (m *MockProductService) ProductMapGet(ctx context.Context, req *connect.Request[v1.ProductMapGetRequest]) (*connect.Response[v1.ProductMapGetResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ProductMapGet", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ProductMapGetResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockProductServiceMockRecorder) ProductMapGet(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductMapGet", reflect.TypeOf((*MockProductService)(nil).ProductMapGet), ctx, req)
}

func (m *MockProductServiceClient) ProductMapGet(ctx context.Context, req *connect.Request[v1.ProductMapGetRequest]) (*connect.Response[v1.ProductMapGetResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ProductMapGet", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ProductMapGetResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockProductServiceClientMockRecorder) ProductMapGet(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductMapGet", reflect.TypeOf((*MockProductService)(nil).ProductMapGet), ctx, req)
}

func (m *MockProductService) ProductMapConnect(ctx context.Context, req *connect.Request[v1.ProductMapConnectRequest]) (*connect.Response[v1.ProductMapConnectResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ProductMapConnect", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ProductMapConnectResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockProductServiceMockRecorder) ProductMapConnect(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductMapConnect", reflect.TypeOf((*MockProductService)(nil).ProductMapConnect), ctx, req)
}

func (m *MockProductServiceClient) ProductMapConnect(ctx context.Context, req *connect.Request[v1.ProductMapConnectRequest]) (*connect.Response[v1.ProductMapConnectResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ProductMapConnect", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ProductMapConnectResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockProductServiceClientMockRecorder) ProductMapConnect(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductMapConnect", reflect.TypeOf((*MockProductService)(nil).ProductMapConnect), ctx, req)
}

func (m *MockProductService) ProductByIDs(ctx context.Context, req *connect.Request[v1.ProductByIDsRequest]) (*connect.Response[v1.ProductByIDsResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ProductByIDs", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ProductByIDsResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockProductServiceMockRecorder) ProductByIDs(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductByIDs", reflect.TypeOf((*MockProductService)(nil).ProductByIDs), ctx, req)
}

func (m *MockProductServiceClient) ProductByIDs(ctx context.Context, req *connect.Request[v1.ProductByIDsRequest]) (*connect.Response[v1.ProductByIDsResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ProductByIDs", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ProductByIDsResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockProductServiceClientMockRecorder) ProductByIDs(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductByIDs", reflect.TypeOf((*MockProductService)(nil).ProductByIDs), ctx, req)
}

func (m *MockProductService) ProductOrderInfo(ctx context.Context, req *connect.Request[v1.ProductOrderInfoRequest]) (*connect.Response[v1.ProductOrderInfoResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ProductOrderInfo", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ProductOrderInfoResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockProductServiceMockRecorder) ProductOrderInfo(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductOrderInfo", reflect.TypeOf((*MockProductService)(nil).ProductOrderInfo), ctx, req)
}

func (m *MockProductServiceClient) ProductOrderInfo(ctx context.Context, req *connect.Request[v1.ProductOrderInfoRequest]) (*connect.Response[v1.ProductOrderInfoResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ProductOrderInfo", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ProductOrderInfoResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockProductServiceClientMockRecorder) ProductOrderInfo(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductOrderInfo", reflect.TypeOf((*MockProductService)(nil).ProductOrderInfo), ctx, req)
}

func (m *MockProductService) ProductListExport(ctx context.Context, req *connect.Request[v1.ProductListExportRequest], stream *connect.ServerStream[v1.ProductListExportResponse]) error {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ProductListExport", ctx, req, stream)
    ret0, _ := ret[0].(error)
    return ret0
}

func (mr *MockProductServiceMockRecorder) ProductListExport(ctx, req, stream interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductListExport", reflect.TypeOf((*MockProductService)(nil).ProductListExport), ctx, req, stream)
}

func (m *MockProductServiceClient) ProductListExport(ctx context.Context, req *connect.Request[v1.ProductListExportRequest]) (*connect.ServerStreamForClient[v1.ProductListExportResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ProductListExport", ctx, req)
    ret0, _ := ret[0].(*connect.ServerStreamForClient[v1.ProductListExportResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockProductServiceClientMockRecorder) ProductListExport(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductListExport", reflect.TypeOf((*MockProductService)(nil).ProductListExport), ctx, req)
}

func (m *MockProductService) ProductSearch(ctx context.Context, req *connect.Request[v1.ProductSearchRequest]) (*connect.Response[v1.ProductSearchResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ProductSearch", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ProductSearchResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockProductServiceMockRecorder) ProductSearch(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductSearch", reflect.TypeOf((*MockProductService)(nil).ProductSearch), ctx, req)
}

func (m *MockProductServiceClient) ProductSearch(ctx context.Context, req *connect.Request[v1.ProductSearchRequest]) (*connect.Response[v1.ProductSearchResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ProductSearch", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ProductSearchResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockProductServiceClientMockRecorder) ProductSearch(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductSearch", reflect.TypeOf((*MockProductService)(nil).ProductSearch), ctx, req)
}

func (m *MockProductService) ProductListSearch(ctx context.Context, req *connect.Request[v1.ProductListSearchRequest]) (*connect.Response[v1.ProductListSearchResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ProductListSearch", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ProductListSearchResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockProductServiceMockRecorder) ProductListSearch(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductListSearch", reflect.TypeOf((*MockProductService)(nil).ProductListSearch), ctx, req)
}

func (m *MockProductServiceClient) ProductListSearch(ctx context.Context, req *connect.Request[v1.ProductListSearchRequest]) (*connect.Response[v1.ProductListSearchResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ProductListSearch", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ProductListSearchResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockProductServiceClientMockRecorder) ProductListSearch(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductListSearch", reflect.TypeOf((*MockProductService)(nil).ProductListSearch), ctx, req)
}

func (m *MockProductService) ProductCreate(ctx context.Context, req *connect.Request[v1.ProductCreateRequest]) (*connect.Response[v1.ProductCreateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ProductCreate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ProductCreateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockProductServiceMockRecorder) ProductCreate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductCreate", reflect.TypeOf((*MockProductService)(nil).ProductCreate), ctx, req)
}

func (m *MockProductServiceClient) ProductCreate(ctx context.Context, req *connect.Request[v1.ProductCreateRequest]) (*connect.Response[v1.ProductCreateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ProductCreate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ProductCreateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockProductServiceClientMockRecorder) ProductCreate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductCreate", reflect.TypeOf((*MockProductService)(nil).ProductCreate), ctx, req)
}

func (m *MockProductService) ProductUpdate(ctx context.Context, req *connect.Request[v1.ProductUpdateRequest]) (*connect.Response[v1.ProductUpdateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ProductUpdate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ProductUpdateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockProductServiceMockRecorder) ProductUpdate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductUpdate", reflect.TypeOf((*MockProductService)(nil).ProductUpdate), ctx, req)
}

func (m *MockProductServiceClient) ProductUpdate(ctx context.Context, req *connect.Request[v1.ProductUpdateRequest]) (*connect.Response[v1.ProductUpdateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ProductUpdate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ProductUpdateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockProductServiceClientMockRecorder) ProductUpdate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductUpdate", reflect.TypeOf((*MockProductService)(nil).ProductUpdate), ctx, req)
}

func (m *MockProductService) ProductDelete(ctx context.Context, req *connect.Request[v1.ProductDeleteRequest]) (*connect.Response[v1.ProductDeleteResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ProductDelete", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ProductDeleteResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockProductServiceMockRecorder) ProductDelete(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductDelete", reflect.TypeOf((*MockProductService)(nil).ProductDelete), ctx, req)
}

func (m *MockProductServiceClient) ProductDelete(ctx context.Context, req *connect.Request[v1.ProductDeleteRequest]) (*connect.Response[v1.ProductDeleteResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ProductDelete", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ProductDeleteResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockProductServiceClientMockRecorder) ProductDelete(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductDelete", reflect.TypeOf((*MockProductService)(nil).ProductDelete), ctx, req)
}

func (m *MockProductService) ProductList(ctx context.Context, req *connect.Request[v1.ProductListRequest]) (*connect.Response[v1.ProductListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ProductList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ProductListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockProductServiceMockRecorder) ProductList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductList", reflect.TypeOf((*MockProductService)(nil).ProductList), ctx, req)
}

func (m *MockProductServiceClient) ProductList(ctx context.Context, req *connect.Request[v1.ProductListRequest]) (*connect.Response[v1.ProductListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ProductList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ProductListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockProductServiceClientMockRecorder) ProductList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductList", reflect.TypeOf((*MockProductService)(nil).ProductList), ctx, req)
}

func (m *MockProductService) ProductDetail(ctx context.Context, req *connect.Request[v1.ProductDetailRequest]) (*connect.Response[v1.ProductDetailResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ProductDetail", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ProductDetailResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockProductServiceMockRecorder) ProductDetail(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductDetail", reflect.TypeOf((*MockProductService)(nil).ProductDetail), ctx, req)
}

func (m *MockProductServiceClient) ProductDetail(ctx context.Context, req *connect.Request[v1.ProductDetailRequest]) (*connect.Response[v1.ProductDetailResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ProductDetail", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ProductDetailResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockProductServiceClientMockRecorder) ProductDetail(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductDetail", reflect.TypeOf((*MockProductService)(nil).ProductDetail), ctx, req)
}

func (m *MockProductService) ProductCodeGenerate(ctx context.Context, req *connect.Request[v1.ProductCodeGenerateRequest]) (*connect.Response[v1.ProductCodeGenerateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ProductCodeGenerate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ProductCodeGenerateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockProductServiceMockRecorder) ProductCodeGenerate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductCodeGenerate", reflect.TypeOf((*MockProductService)(nil).ProductCodeGenerate), ctx, req)
}

func (m *MockProductServiceClient) ProductCodeGenerate(ctx context.Context, req *connect.Request[v1.ProductCodeGenerateRequest]) (*connect.Response[v1.ProductCodeGenerateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ProductCodeGenerate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ProductCodeGenerateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockProductServiceClientMockRecorder) ProductCodeGenerate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductCodeGenerate", reflect.TypeOf((*MockProductService)(nil).ProductCodeGenerate), ctx, req)
}

type MockCategoryService struct {
    ctrl     *gomock.Controller
    recorder *MockCategoryServiceMockRecorder
}

type MockCategoryServiceMockRecorder struct {
    mock *MockCategoryService
}

func NewMockCategoryService(ctrl *gomock.Controller) *MockCategoryService {
    mock := &MockCategoryService{ctrl: ctrl}
    mock.recorder = &MockCategoryServiceMockRecorder{mock}
    return mock
}

func (m *MockCategoryService) EXPECT() *MockCategoryServiceMockRecorder {
    return m.recorder
}

type MockCategoryServiceClient struct {
    ctrl     *gomock.Controller
    recorder *MockCategoryServiceClientMockRecorder
}

type MockCategoryServiceClientMockRecorder struct {
    mock *MockCategoryServiceClient
}

func NewMockCategoryServiceClient(ctrl *gomock.Controller) *MockCategoryServiceClient {
    mock := &MockCategoryServiceClient{ctrl: ctrl}
    mock.recorder = &MockCategoryServiceClientMockRecorder{mock}
    return mock
}

func (m *MockCategoryServiceClient) EXPECT() *MockCategoryServiceClientMockRecorder {
    return m.recorder
}

func (m *MockCategoryService) CategoryCreate(ctx context.Context, req *connect.Request[v1.CategoryCreateRequest]) (*connect.Response[v1.CategoryCreateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "CategoryCreate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.CategoryCreateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockCategoryServiceMockRecorder) CategoryCreate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CategoryCreate", reflect.TypeOf((*MockCategoryService)(nil).CategoryCreate), ctx, req)
}

func (m *MockCategoryServiceClient) CategoryCreate(ctx context.Context, req *connect.Request[v1.CategoryCreateRequest]) (*connect.Response[v1.CategoryCreateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "CategoryCreate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.CategoryCreateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockCategoryServiceClientMockRecorder) CategoryCreate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CategoryCreate", reflect.TypeOf((*MockCategoryService)(nil).CategoryCreate), ctx, req)
}

func (m *MockCategoryService) CategoryUpdate(ctx context.Context, req *connect.Request[v1.CategoryUpdateRequest]) (*connect.Response[v1.CategoryUpdateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "CategoryUpdate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.CategoryUpdateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockCategoryServiceMockRecorder) CategoryUpdate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CategoryUpdate", reflect.TypeOf((*MockCategoryService)(nil).CategoryUpdate), ctx, req)
}

func (m *MockCategoryServiceClient) CategoryUpdate(ctx context.Context, req *connect.Request[v1.CategoryUpdateRequest]) (*connect.Response[v1.CategoryUpdateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "CategoryUpdate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.CategoryUpdateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockCategoryServiceClientMockRecorder) CategoryUpdate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CategoryUpdate", reflect.TypeOf((*MockCategoryService)(nil).CategoryUpdate), ctx, req)
}

func (m *MockCategoryService) CategoryDelete(ctx context.Context, req *connect.Request[v1.CategoryDeleteRequest]) (*connect.Response[v1.CategoryDeleteResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "CategoryDelete", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.CategoryDeleteResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockCategoryServiceMockRecorder) CategoryDelete(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CategoryDelete", reflect.TypeOf((*MockCategoryService)(nil).CategoryDelete), ctx, req)
}

func (m *MockCategoryServiceClient) CategoryDelete(ctx context.Context, req *connect.Request[v1.CategoryDeleteRequest]) (*connect.Response[v1.CategoryDeleteResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "CategoryDelete", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.CategoryDeleteResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockCategoryServiceClientMockRecorder) CategoryDelete(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CategoryDelete", reflect.TypeOf((*MockCategoryService)(nil).CategoryDelete), ctx, req)
}

func (m *MockCategoryService) CategoryList(ctx context.Context, req *connect.Request[v1.CategoryListRequest]) (*connect.Response[v1.CategoryListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "CategoryList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.CategoryListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockCategoryServiceMockRecorder) CategoryList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CategoryList", reflect.TypeOf((*MockCategoryService)(nil).CategoryList), ctx, req)
}

func (m *MockCategoryServiceClient) CategoryList(ctx context.Context, req *connect.Request[v1.CategoryListRequest]) (*connect.Response[v1.CategoryListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "CategoryList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.CategoryListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockCategoryServiceClientMockRecorder) CategoryList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CategoryList", reflect.TypeOf((*MockCategoryService)(nil).CategoryList), ctx, req)
}

