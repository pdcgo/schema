package expense_iface_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/expense_iface/v1"
)

type MockExpenseService struct {
    ctrl     *gomock.Controller
    recorder *MockExpenseServiceMockRecorder
}

type MockExpenseServiceMockRecorder struct {
    mock *MockExpenseService
}

func NewMockExpenseService(ctrl *gomock.Controller) *MockExpenseService {
    mock := &MockExpenseService{ctrl: ctrl}
    mock.recorder = &MockExpenseServiceMockRecorder{mock}
    return mock
}

func (m *MockExpenseService) EXPECT() *MockExpenseServiceMockRecorder {
    return m.recorder
}

type MockExpenseServiceClient struct {
    ctrl     *gomock.Controller
    recorder *MockExpenseServiceClientMockRecorder
}

type MockExpenseServiceClientMockRecorder struct {
    mock *MockExpenseServiceClient
}

func NewMockExpenseServiceClient(ctrl *gomock.Controller) *MockExpenseServiceClient {
    mock := &MockExpenseServiceClient{ctrl: ctrl}
    mock.recorder = &MockExpenseServiceClientMockRecorder{mock}
    return mock
}

func (m *MockExpenseServiceClient) EXPECT() *MockExpenseServiceClientMockRecorder {
    return m.recorder
}

func (m *MockExpenseService) Hello(ctx context.Context, req *connect.Request[v1.HelloRequest]) (*connect.Response[v1.HelloResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "Hello", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.HelloResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockExpenseServiceMockRecorder) Hello(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hello", reflect.TypeOf((*MockExpenseService)(nil).Hello), ctx, req)
}

func (m *MockExpenseServiceClient) Hello(ctx context.Context, req *connect.Request[v1.HelloRequest]) (*connect.Response[v1.HelloResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "Hello", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.HelloResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockExpenseServiceClientMockRecorder) Hello(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hello", reflect.TypeOf((*MockExpenseService)(nil).Hello), ctx, req)
}

