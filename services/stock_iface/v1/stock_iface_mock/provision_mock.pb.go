package stock_iface_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/stock_iface/v1"
)

type MockStockProvisionService struct {
    ctrl     *gomock.Controller
    recorder *MockStockProvisionServiceMockRecorder
}

type MockStockProvisionServiceMockRecorder struct {
    mock *MockStockProvisionService
}

func NewMockStockProvisionService(ctrl *gomock.Controller) *MockStockProvisionService {
    mock := &MockStockProvisionService{ctrl: ctrl}
    mock.recorder = &MockStockProvisionServiceMockRecorder{mock}
    return mock
}

func (m *MockStockProvisionService) EXPECT() *MockStockProvisionServiceMockRecorder {
    return m.recorder
}

type MockStockProvisionServiceClient struct {
    ctrl     *gomock.Controller
    recorder *MockStockProvisionServiceClientMockRecorder
}

type MockStockProvisionServiceClientMockRecorder struct {
    mock *MockStockProvisionServiceClient
}

func NewMockStockProvisionServiceClient(ctrl *gomock.Controller) *MockStockProvisionServiceClient {
    mock := &MockStockProvisionServiceClient{ctrl: ctrl}
    mock.recorder = &MockStockProvisionServiceClientMockRecorder{mock}
    return mock
}

func (m *MockStockProvisionServiceClient) EXPECT() *MockStockProvisionServiceClientMockRecorder {
    return m.recorder
}

func (m *MockStockProvisionService) StockProvision(ctx context.Context, req *connect.Request[v1.StockProvisionRequest]) (*connect.Response[v1.StockProvisionResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "StockProvision", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.StockProvisionResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockStockProvisionServiceMockRecorder) StockProvision(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StockProvision", reflect.TypeOf((*MockStockProvisionService)(nil).StockProvision), ctx, req)
}

func (m *MockStockProvisionServiceClient) StockProvision(ctx context.Context, req *connect.Request[v1.StockProvisionRequest]) (*connect.Response[v1.StockProvisionResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "StockProvision", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.StockProvisionResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockStockProvisionServiceClientMockRecorder) StockProvision(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StockProvision", reflect.TypeOf((*MockStockProvisionService)(nil).StockProvision), ctx, req)
}

