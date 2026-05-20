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

