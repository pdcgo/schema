package invoice_iface_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/invoice_iface/v1"
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

func (m *MockInvoiceService) Balance(ctx context.Context, req *connect.Request[v1.BalanceRequest]) (*connect.Response[v1.BalanceResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "Balance", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.BalanceResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInvoiceServiceMockRecorder) Balance(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Balance", reflect.TypeOf((*MockInvoiceService)(nil).Balance), ctx, req)
}

func (m *MockInvoiceServiceClient) Balance(ctx context.Context, req *connect.Request[v1.BalanceRequest]) (*connect.Response[v1.BalanceResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "Balance", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.BalanceResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInvoiceServiceClientMockRecorder) Balance(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Balance", reflect.TypeOf((*MockInvoiceService)(nil).Balance), ctx, req)
}

func (m *MockInvoiceService) TransactionList(ctx context.Context, req *connect.Request[v1.TransactionListRequest]) (*connect.Response[v1.TransactionListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TransactionList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.TransactionListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInvoiceServiceMockRecorder) TransactionList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionList", reflect.TypeOf((*MockInvoiceService)(nil).TransactionList), ctx, req)
}

func (m *MockInvoiceServiceClient) TransactionList(ctx context.Context, req *connect.Request[v1.TransactionListRequest]) (*connect.Response[v1.TransactionListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TransactionList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.TransactionListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInvoiceServiceClientMockRecorder) TransactionList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionList", reflect.TypeOf((*MockInvoiceService)(nil).TransactionList), ctx, req)
}

func (m *MockInvoiceService) CreateInvoice(ctx context.Context, req *connect.Request[v1.CreateInvoiceRequest]) (*connect.Response[v1.CreateInvoiceResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "CreateInvoice", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.CreateInvoiceResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInvoiceServiceMockRecorder) CreateInvoice(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInvoice", reflect.TypeOf((*MockInvoiceService)(nil).CreateInvoice), ctx, req)
}

func (m *MockInvoiceServiceClient) CreateInvoice(ctx context.Context, req *connect.Request[v1.CreateInvoiceRequest]) (*connect.Response[v1.CreateInvoiceResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "CreateInvoice", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.CreateInvoiceResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInvoiceServiceClientMockRecorder) CreateInvoice(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInvoice", reflect.TypeOf((*MockInvoiceService)(nil).CreateInvoice), ctx, req)
}

func (m *MockInvoiceService) InvoiceDetail(ctx context.Context, req *connect.Request[v1.InvoiceDetailRequest]) (*connect.Response[v1.InvoiceDetailResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "InvoiceDetail", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.InvoiceDetailResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInvoiceServiceMockRecorder) InvoiceDetail(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvoiceDetail", reflect.TypeOf((*MockInvoiceService)(nil).InvoiceDetail), ctx, req)
}

func (m *MockInvoiceServiceClient) InvoiceDetail(ctx context.Context, req *connect.Request[v1.InvoiceDetailRequest]) (*connect.Response[v1.InvoiceDetailResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "InvoiceDetail", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.InvoiceDetailResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInvoiceServiceClientMockRecorder) InvoiceDetail(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvoiceDetail", reflect.TypeOf((*MockInvoiceService)(nil).InvoiceDetail), ctx, req)
}

func (m *MockInvoiceService) ListInvoice(ctx context.Context, req *connect.Request[v1.ListInvoiceRequest]) (*connect.Response[v1.ListInvoiceResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ListInvoice", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ListInvoiceResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInvoiceServiceMockRecorder) ListInvoice(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListInvoice", reflect.TypeOf((*MockInvoiceService)(nil).ListInvoice), ctx, req)
}

func (m *MockInvoiceServiceClient) ListInvoice(ctx context.Context, req *connect.Request[v1.ListInvoiceRequest]) (*connect.Response[v1.ListInvoiceResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ListInvoice", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ListInvoiceResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInvoiceServiceClientMockRecorder) ListInvoice(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListInvoice", reflect.TypeOf((*MockInvoiceService)(nil).ListInvoice), ctx, req)
}

func (m *MockInvoiceService) ConfirmInvoice(ctx context.Context, req *connect.Request[v1.ConfirmInvoiceRequest]) (*connect.Response[v1.ConfirmInvoiceResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ConfirmInvoice", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ConfirmInvoiceResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInvoiceServiceMockRecorder) ConfirmInvoice(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfirmInvoice", reflect.TypeOf((*MockInvoiceService)(nil).ConfirmInvoice), ctx, req)
}

func (m *MockInvoiceServiceClient) ConfirmInvoice(ctx context.Context, req *connect.Request[v1.ConfirmInvoiceRequest]) (*connect.Response[v1.ConfirmInvoiceResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ConfirmInvoice", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ConfirmInvoiceResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInvoiceServiceClientMockRecorder) ConfirmInvoice(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfirmInvoice", reflect.TypeOf((*MockInvoiceService)(nil).ConfirmInvoice), ctx, req)
}

func (m *MockInvoiceService) RejectInvoice(ctx context.Context, req *connect.Request[v1.RejectInvoiceRequest]) (*connect.Response[v1.RejectInvoiceResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "RejectInvoice", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.RejectInvoiceResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInvoiceServiceMockRecorder) RejectInvoice(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RejectInvoice", reflect.TypeOf((*MockInvoiceService)(nil).RejectInvoice), ctx, req)
}

func (m *MockInvoiceServiceClient) RejectInvoice(ctx context.Context, req *connect.Request[v1.RejectInvoiceRequest]) (*connect.Response[v1.RejectInvoiceResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "RejectInvoice", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.RejectInvoiceResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInvoiceServiceClientMockRecorder) RejectInvoice(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RejectInvoice", reflect.TypeOf((*MockInvoiceService)(nil).RejectInvoice), ctx, req)
}

func (m *MockInvoiceService) DailySummary(ctx context.Context, req *connect.Request[v1.DailySummaryRequest]) (*connect.Response[v1.DailySummaryResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "DailySummary", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.DailySummaryResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInvoiceServiceMockRecorder) DailySummary(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DailySummary", reflect.TypeOf((*MockInvoiceService)(nil).DailySummary), ctx, req)
}

func (m *MockInvoiceServiceClient) DailySummary(ctx context.Context, req *connect.Request[v1.DailySummaryRequest]) (*connect.Response[v1.DailySummaryResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "DailySummary", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.DailySummaryResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockInvoiceServiceClientMockRecorder) DailySummary(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DailySummary", reflect.TypeOf((*MockInvoiceService)(nil).DailySummary), ctx, req)
}

