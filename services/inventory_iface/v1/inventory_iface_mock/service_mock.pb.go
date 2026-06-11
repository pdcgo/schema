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

func (m *MockInventoryService) Hello(ctx context.Context, req *connect.Request[v1.HelloRequest]) (*connect.Response[v1.HelloResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "Hello", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.HelloResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceMockRecorder) Hello(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hello", reflect.TypeOf((*MockInventoryService)(nil).Hello), ctx, req)
}

func (m *MockInventoryServiceClient) Hello(ctx context.Context, req *connect.Request[v1.HelloRequest]) (*connect.Response[v1.HelloResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "Hello", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.HelloResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInventoryServiceClientMockRecorder) Hello(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hello", reflect.TypeOf((*MockInventoryService)(nil).Hello), ctx, req)
}

