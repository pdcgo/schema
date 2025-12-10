package accounting_iface_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/accounting_iface/v1"
)

type MockStatementService struct {
    ctrl     *gomock.Controller
    recorder *MockStatementServiceMockRecorder
}

type MockStatementServiceMockRecorder struct {
    mock *MockStatementService
}

func NewMockStatementService(ctrl *gomock.Controller) *MockStatementService {
    mock := &MockStatementService{ctrl: ctrl}
    mock.recorder = &MockStatementServiceMockRecorder{mock}
    return mock
}

func (m *MockStatementService) EXPECT() *MockStatementServiceMockRecorder {
    return m.recorder
}

type MockStatementServiceClient struct {
    ctrl     *gomock.Controller
    recorder *MockStatementServiceClientMockRecorder
}

type MockStatementServiceClientMockRecorder struct {
    mock *MockStatementServiceClient
}

func NewMockStatementServiceClient(ctrl *gomock.Controller) *MockStatementServiceClient {
    mock := &MockStatementServiceClient{ctrl: ctrl}
    mock.recorder = &MockStatementServiceClientMockRecorder{mock}
    return mock
}

func (m *MockStatementServiceClient) EXPECT() *MockStatementServiceClientMockRecorder {
    return m.recorder
}

func (m *MockStatementService) StatementBalance(ctx context.Context, req *connect.Request[v1.StatementBalanceRequest], stream *connect.ServerStream[v1.StatementBalanceResponse]) error {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "StatementBalance", ctx, req, stream)
    ret0, _ := ret[0].(error)
    return ret0
}

func (mr *MockStatementServiceMockRecorder) StatementBalance(ctx, req, stream interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatementBalance", reflect.TypeOf((*MockStatementService)(nil).StatementBalance), ctx, req, stream)
}

func (m *MockStatementServiceClient) StatementBalance(ctx context.Context, req *connect.Request[v1.StatementBalanceRequest]) (*connect.ServerStreamForClient[v1.StatementBalanceResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "StatementBalance", ctx, req)
    ret0, _ := ret[0].(*connect.ServerStreamForClient[v1.StatementBalanceResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockStatementServiceClientMockRecorder) StatementBalance(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatementBalance", reflect.TypeOf((*MockStatementService)(nil).StatementBalance), ctx, req)
}

func (m *MockStatementService) StatementIncome(ctx context.Context, req *connect.Request[v1.StatementIncomeRequest], stream *connect.ServerStream[v1.StatementIncomeResponse]) error {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "StatementIncome", ctx, req, stream)
    ret0, _ := ret[0].(error)
    return ret0
}

func (mr *MockStatementServiceMockRecorder) StatementIncome(ctx, req, stream interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatementIncome", reflect.TypeOf((*MockStatementService)(nil).StatementIncome), ctx, req, stream)
}

func (m *MockStatementServiceClient) StatementIncome(ctx context.Context, req *connect.Request[v1.StatementIncomeRequest]) (*connect.ServerStreamForClient[v1.StatementIncomeResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "StatementIncome", ctx, req)
    ret0, _ := ret[0].(*connect.ServerStreamForClient[v1.StatementIncomeResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockStatementServiceClientMockRecorder) StatementIncome(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatementIncome", reflect.TypeOf((*MockStatementService)(nil).StatementIncome), ctx, req)
}

func (m *MockStatementService) StatementCashFlow(ctx context.Context, req *connect.Request[v1.StatementCashFlowRequest]) (*connect.Response[v1.StatementCashFlowResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "StatementCashFlow", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.StatementCashFlowResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockStatementServiceMockRecorder) StatementCashFlow(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatementCashFlow", reflect.TypeOf((*MockStatementService)(nil).StatementCashFlow), ctx, req)
}

func (m *MockStatementServiceClient) StatementCashFlow(ctx context.Context, req *connect.Request[v1.StatementCashFlowRequest]) (*connect.Response[v1.StatementCashFlowResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "StatementCashFlow", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.StatementCashFlowResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockStatementServiceClientMockRecorder) StatementCashFlow(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatementCashFlow", reflect.TypeOf((*MockStatementService)(nil).StatementCashFlow), ctx, req)
}

