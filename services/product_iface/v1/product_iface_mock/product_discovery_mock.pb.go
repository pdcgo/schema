package product_iface_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/product_iface/v1"
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

func (m *MockProductDiscoveryService) ProductDiscovery(ctx context.Context, req *connect.Request[v1.ProductDiscoveryRequest]) (*connect.Response[v1.ProductDiscoveryResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ProductDiscovery", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ProductDiscoveryResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockProductDiscoveryServiceMockRecorder) ProductDiscovery(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductDiscovery", reflect.TypeOf((*MockProductDiscoveryService)(nil).ProductDiscovery), ctx, req)
}

func (m *MockProductDiscoveryServiceClient) ProductDiscovery(ctx context.Context, req *connect.Request[v1.ProductDiscoveryRequest]) (*connect.Response[v1.ProductDiscoveryResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ProductDiscovery", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ProductDiscoveryResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockProductDiscoveryServiceClientMockRecorder) ProductDiscovery(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductDiscovery", reflect.TypeOf((*MockProductDiscoveryService)(nil).ProductDiscovery), ctx, req)
}

func (m *MockProductDiscoveryService) IndexAdd(ctx context.Context, req *connect.Request[v1.IndexAddRequest]) (*connect.Response[v1.IndexAddResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "IndexAdd", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.IndexAddResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockProductDiscoveryServiceMockRecorder) IndexAdd(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IndexAdd", reflect.TypeOf((*MockProductDiscoveryService)(nil).IndexAdd), ctx, req)
}

func (m *MockProductDiscoveryServiceClient) IndexAdd(ctx context.Context, req *connect.Request[v1.IndexAddRequest]) (*connect.Response[v1.IndexAddResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "IndexAdd", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.IndexAddResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockProductDiscoveryServiceClientMockRecorder) IndexAdd(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IndexAdd", reflect.TypeOf((*MockProductDiscoveryService)(nil).IndexAdd), ctx, req)
}

