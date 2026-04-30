package selling_iface_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/selling_iface/v1"
)

type MockProductDiscoveryService struct {
    ctrl     *gomock.Controller
    recorder *MockProductDiscoveryServiceMockRecorder
}

type MockProductDiscoveryServiceMockRecorder struct {
    mock *MockProductDiscoveryService
}

func NewMockProductDiscoveryService(ctrl *gomock.Controller) *MockProductDiscoveryService {
    mock := &MockProductDiscoveryService{ctrl: ctrl}
    mock.recorder = &MockProductDiscoveryServiceMockRecorder{mock}
    return mock
}

func (m *MockProductDiscoveryService) EXPECT() *MockProductDiscoveryServiceMockRecorder {
    return m.recorder
}

type MockProductDiscoveryServiceClient struct {
    ctrl     *gomock.Controller
    recorder *MockProductDiscoveryServiceClientMockRecorder
}

type MockProductDiscoveryServiceClientMockRecorder struct {
    mock *MockProductDiscoveryServiceClient
}

func NewMockProductDiscoveryServiceClient(ctrl *gomock.Controller) *MockProductDiscoveryServiceClient {
    mock := &MockProductDiscoveryServiceClient{ctrl: ctrl}
    mock.recorder = &MockProductDiscoveryServiceClientMockRecorder{mock}
    return mock
}

func (m *MockProductDiscoveryServiceClient) EXPECT() *MockProductDiscoveryServiceClientMockRecorder {
    return m.recorder
}

func (m *MockProductDiscoveryService) SearchProduct(ctx context.Context, req *connect.Request[v1.SearchProductRequest]) (*connect.Response[v1.SearchProductResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "SearchProduct", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.SearchProductResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockProductDiscoveryServiceMockRecorder) SearchProduct(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchProduct", reflect.TypeOf((*MockProductDiscoveryService)(nil).SearchProduct), ctx, req)
}

func (m *MockProductDiscoveryServiceClient) SearchProduct(ctx context.Context, req *connect.Request[v1.SearchProductRequest]) (*connect.Response[v1.SearchProductResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "SearchProduct", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.SearchProductResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockProductDiscoveryServiceClientMockRecorder) SearchProduct(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchProduct", reflect.TypeOf((*MockProductDiscoveryService)(nil).SearchProduct), ctx, req)
}

func (m *MockProductDiscoveryService) ProductDetail(ctx context.Context, req *connect.Request[v1.ProductDetailRequest]) (*connect.Response[v1.ProductDetailResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ProductDetail", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ProductDetailResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockProductDiscoveryServiceMockRecorder) ProductDetail(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductDetail", reflect.TypeOf((*MockProductDiscoveryService)(nil).ProductDetail), ctx, req)
}

func (m *MockProductDiscoveryServiceClient) ProductDetail(ctx context.Context, req *connect.Request[v1.ProductDetailRequest]) (*connect.Response[v1.ProductDetailResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ProductDetail", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ProductDetailResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockProductDiscoveryServiceClientMockRecorder) ProductDetail(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductDetail", reflect.TypeOf((*MockProductDiscoveryService)(nil).ProductDetail), ctx, req)
}

