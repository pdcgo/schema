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

