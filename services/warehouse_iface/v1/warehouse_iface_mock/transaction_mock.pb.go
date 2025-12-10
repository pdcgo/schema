package warehouse_iface_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/warehouse_iface/v1"
)

type MockTransactionService struct {
    ctrl     *gomock.Controller
    recorder *MockTransactionServiceMockRecorder
}

type MockTransactionServiceMockRecorder struct {
    mock *MockTransactionService
}

func NewMockTransactionService(ctrl *gomock.Controller) *MockTransactionService {
    mock := &MockTransactionService{ctrl: ctrl}
    mock.recorder = &MockTransactionServiceMockRecorder{mock}
    return mock
}

func (m *MockTransactionService) EXPECT() *MockTransactionServiceMockRecorder {
    return m.recorder
}

type MockTransactionServiceClient struct {
    ctrl     *gomock.Controller
    recorder *MockTransactionServiceClientMockRecorder
}

type MockTransactionServiceClientMockRecorder struct {
    mock *MockTransactionServiceClient
}

func NewMockTransactionServiceClient(ctrl *gomock.Controller) *MockTransactionServiceClient {
    mock := &MockTransactionServiceClient{ctrl: ctrl}
    mock.recorder = &MockTransactionServiceClientMockRecorder{mock}
    return mock
}

func (m *MockTransactionServiceClient) EXPECT() *MockTransactionServiceClientMockRecorder {
    return m.recorder
}

func (m *MockTransactionService) TransactionItemSearch(ctx context.Context, req *connect.Request[v1.TransactionItemSearchRequest]) (*connect.Response[v1.TransactionItemSearchResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TransactionItemSearch", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.TransactionItemSearchResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockTransactionServiceMockRecorder) TransactionItemSearch(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionItemSearch", reflect.TypeOf((*MockTransactionService)(nil).TransactionItemSearch), ctx, req)
}

func (m *MockTransactionServiceClient) TransactionItemSearch(ctx context.Context, req *connect.Request[v1.TransactionItemSearchRequest]) (*connect.Response[v1.TransactionItemSearchResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TransactionItemSearch", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.TransactionItemSearchResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockTransactionServiceClientMockRecorder) TransactionItemSearch(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionItemSearch", reflect.TypeOf((*MockTransactionService)(nil).TransactionItemSearch), ctx, req)
}

