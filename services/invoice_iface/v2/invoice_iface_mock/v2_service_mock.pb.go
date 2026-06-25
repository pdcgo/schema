package invoice_iface_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v2 "github.com/pdcgo/schema/services/invoice_iface/v2"
)

type MockInvoiceService struct {
    ctrl     *gomock.Controller
    recorder *MockInvoiceServiceMockRecorder
}

type MockInvoiceServiceMockRecorder struct {
    mock *MockInvoiceService
}

func NewMockInvoiceService(ctrl *gomock.Controller) *MockInvoiceService {
    mock := &MockInvoiceService{ctrl: ctrl}
    mock.recorder = &MockInvoiceServiceMockRecorder{mock}
    return mock
}

func (m *MockInvoiceService) EXPECT() *MockInvoiceServiceMockRecorder {
    return m.recorder
}

type MockInvoiceServiceClient struct {
    ctrl     *gomock.Controller
    recorder *MockInvoiceServiceClientMockRecorder
}

type MockInvoiceServiceClientMockRecorder struct {
    mock *MockInvoiceServiceClient
}

func NewMockInvoiceServiceClient(ctrl *gomock.Controller) *MockInvoiceServiceClient {
    mock := &MockInvoiceServiceClient{ctrl: ctrl}
    mock.recorder = &MockInvoiceServiceClientMockRecorder{mock}
    return mock
}

func (m *MockInvoiceServiceClient) EXPECT() *MockInvoiceServiceClientMockRecorder {
    return m.recorder
}

func (m *MockInvoiceService) Overview(ctx context.Context, req *connect.Request[v2.OverviewRequest]) (*connect.Response[v2.OverviewResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "Overview", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v2.OverviewResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInvoiceServiceMockRecorder) Overview(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Overview", reflect.TypeOf((*MockInvoiceService)(nil).Overview), ctx, req)
}

func (m *MockInvoiceServiceClient) Overview(ctx context.Context, req *connect.Request[v2.OverviewRequest]) (*connect.Response[v2.OverviewResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "Overview", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v2.OverviewResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInvoiceServiceClientMockRecorder) Overview(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Overview", reflect.TypeOf((*MockInvoiceService)(nil).Overview), ctx, req)
}

func (m *MockInvoiceService) TeamBalanceList(ctx context.Context, req *connect.Request[v2.TeamBalanceListRequest]) (*connect.Response[v2.TeamBalanceListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TeamBalanceList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v2.TeamBalanceListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInvoiceServiceMockRecorder) TeamBalanceList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeamBalanceList", reflect.TypeOf((*MockInvoiceService)(nil).TeamBalanceList), ctx, req)
}

func (m *MockInvoiceServiceClient) TeamBalanceList(ctx context.Context, req *connect.Request[v2.TeamBalanceListRequest]) (*connect.Response[v2.TeamBalanceListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TeamBalanceList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v2.TeamBalanceListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInvoiceServiceClientMockRecorder) TeamBalanceList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeamBalanceList", reflect.TypeOf((*MockInvoiceService)(nil).TeamBalanceList), ctx, req)
}

func (m *MockInvoiceService) CreatePayment(ctx context.Context, req *connect.Request[v2.CreatePaymentRequest]) (*connect.Response[v2.CreatePaymentResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "CreatePayment", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v2.CreatePaymentResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInvoiceServiceMockRecorder) CreatePayment(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePayment", reflect.TypeOf((*MockInvoiceService)(nil).CreatePayment), ctx, req)
}

func (m *MockInvoiceServiceClient) CreatePayment(ctx context.Context, req *connect.Request[v2.CreatePaymentRequest]) (*connect.Response[v2.CreatePaymentResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "CreatePayment", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v2.CreatePaymentResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInvoiceServiceClientMockRecorder) CreatePayment(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePayment", reflect.TypeOf((*MockInvoiceService)(nil).CreatePayment), ctx, req)
}

func (m *MockInvoiceService) AcceptPayment(ctx context.Context, req *connect.Request[v2.AcceptPaymentRequest]) (*connect.Response[v2.AcceptPaymentResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "AcceptPayment", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v2.AcceptPaymentResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInvoiceServiceMockRecorder) AcceptPayment(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptPayment", reflect.TypeOf((*MockInvoiceService)(nil).AcceptPayment), ctx, req)
}

func (m *MockInvoiceServiceClient) AcceptPayment(ctx context.Context, req *connect.Request[v2.AcceptPaymentRequest]) (*connect.Response[v2.AcceptPaymentResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "AcceptPayment", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v2.AcceptPaymentResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInvoiceServiceClientMockRecorder) AcceptPayment(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptPayment", reflect.TypeOf((*MockInvoiceService)(nil).AcceptPayment), ctx, req)
}

func (m *MockInvoiceService) RejectPayment(ctx context.Context, req *connect.Request[v2.RejectPaymentRequest]) (*connect.Response[v2.RejectPaymentResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "RejectPayment", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v2.RejectPaymentResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInvoiceServiceMockRecorder) RejectPayment(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RejectPayment", reflect.TypeOf((*MockInvoiceService)(nil).RejectPayment), ctx, req)
}

func (m *MockInvoiceServiceClient) RejectPayment(ctx context.Context, req *connect.Request[v2.RejectPaymentRequest]) (*connect.Response[v2.RejectPaymentResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "RejectPayment", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v2.RejectPaymentResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInvoiceServiceClientMockRecorder) RejectPayment(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RejectPayment", reflect.TypeOf((*MockInvoiceService)(nil).RejectPayment), ctx, req)
}

func (m *MockInvoiceService) ListPayment(ctx context.Context, req *connect.Request[v2.ListPaymentRequest]) (*connect.Response[v2.ListPaymentResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ListPayment", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v2.ListPaymentResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInvoiceServiceMockRecorder) ListPayment(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPayment", reflect.TypeOf((*MockInvoiceService)(nil).ListPayment), ctx, req)
}

func (m *MockInvoiceServiceClient) ListPayment(ctx context.Context, req *connect.Request[v2.ListPaymentRequest]) (*connect.Response[v2.ListPaymentResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ListPayment", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v2.ListPaymentResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInvoiceServiceClientMockRecorder) ListPayment(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPayment", reflect.TypeOf((*MockInvoiceService)(nil).ListPayment), ctx, req)
}

func (m *MockInvoiceService) ListIncomingPayment(ctx context.Context, req *connect.Request[v2.ListIncomingPaymentRequest]) (*connect.Response[v2.ListIncomingPaymentResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ListIncomingPayment", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v2.ListIncomingPaymentResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInvoiceServiceMockRecorder) ListIncomingPayment(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListIncomingPayment", reflect.TypeOf((*MockInvoiceService)(nil).ListIncomingPayment), ctx, req)
}

func (m *MockInvoiceServiceClient) ListIncomingPayment(ctx context.Context, req *connect.Request[v2.ListIncomingPaymentRequest]) (*connect.Response[v2.ListIncomingPaymentResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ListIncomingPayment", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v2.ListIncomingPaymentResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInvoiceServiceClientMockRecorder) ListIncomingPayment(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListIncomingPayment", reflect.TypeOf((*MockInvoiceService)(nil).ListIncomingPayment), ctx, req)
}

func (m *MockInvoiceService) ListTeamBalanceLog(ctx context.Context, req *connect.Request[v2.ListTeamBalanceLogRequest]) (*connect.Response[v2.ListTeamBalanceLogResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ListTeamBalanceLog", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v2.ListTeamBalanceLogResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInvoiceServiceMockRecorder) ListTeamBalanceLog(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTeamBalanceLog", reflect.TypeOf((*MockInvoiceService)(nil).ListTeamBalanceLog), ctx, req)
}

func (m *MockInvoiceServiceClient) ListTeamBalanceLog(ctx context.Context, req *connect.Request[v2.ListTeamBalanceLogRequest]) (*connect.Response[v2.ListTeamBalanceLogResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ListTeamBalanceLog", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v2.ListTeamBalanceLogResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInvoiceServiceClientMockRecorder) ListTeamBalanceLog(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTeamBalanceLog", reflect.TypeOf((*MockInvoiceService)(nil).ListTeamBalanceLog), ctx, req)
}

func (m *MockInvoiceService) CreateBalanceLog(ctx context.Context, req *connect.Request[v2.CreateBalanceLogRequest]) (*connect.Response[v2.CreateBalanceLogResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "CreateBalanceLog", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v2.CreateBalanceLogResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInvoiceServiceMockRecorder) CreateBalanceLog(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBalanceLog", reflect.TypeOf((*MockInvoiceService)(nil).CreateBalanceLog), ctx, req)
}

func (m *MockInvoiceServiceClient) CreateBalanceLog(ctx context.Context, req *connect.Request[v2.CreateBalanceLogRequest]) (*connect.Response[v2.CreateBalanceLogResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "CreateBalanceLog", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v2.CreateBalanceLogResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInvoiceServiceClientMockRecorder) CreateBalanceLog(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBalanceLog", reflect.TypeOf((*MockInvoiceService)(nil).CreateBalanceLog), ctx, req)
}

func (m *MockInvoiceService) TeamReconcile(ctx context.Context, req *connect.Request[v2.TeamReconcileRequest]) (*connect.Response[v2.TeamReconcileResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TeamReconcile", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v2.TeamReconcileResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInvoiceServiceMockRecorder) TeamReconcile(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeamReconcile", reflect.TypeOf((*MockInvoiceService)(nil).TeamReconcile), ctx, req)
}

func (m *MockInvoiceServiceClient) TeamReconcile(ctx context.Context, req *connect.Request[v2.TeamReconcileRequest]) (*connect.Response[v2.TeamReconcileResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TeamReconcile", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v2.TeamReconcileResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInvoiceServiceClientMockRecorder) TeamReconcile(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeamReconcile", reflect.TypeOf((*MockInvoiceService)(nil).TeamReconcile), ctx, req)
}

