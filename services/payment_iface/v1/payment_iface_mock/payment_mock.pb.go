package payment_iface_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/payment_iface/v1"
)

type MockPaymentService struct {
    ctrl     *gomock.Controller
    recorder *MockPaymentServiceMockRecorder
}

type MockPaymentServiceMockRecorder struct {
    mock *MockPaymentService
}

func NewMockPaymentService(ctrl *gomock.Controller) *MockPaymentService {
    mock := &MockPaymentService{ctrl: ctrl}
    mock.recorder = &MockPaymentServiceMockRecorder{mock}
    return mock
}

func (m *MockPaymentService) EXPECT() *MockPaymentServiceMockRecorder {
    return m.recorder
}

type MockPaymentServiceClient struct {
    ctrl     *gomock.Controller
    recorder *MockPaymentServiceClientMockRecorder
}

type MockPaymentServiceClientMockRecorder struct {
    mock *MockPaymentServiceClient
}

func NewMockPaymentServiceClient(ctrl *gomock.Controller) *MockPaymentServiceClient {
    mock := &MockPaymentServiceClient{ctrl: ctrl}
    mock.recorder = &MockPaymentServiceClientMockRecorder{mock}
    return mock
}

func (m *MockPaymentServiceClient) EXPECT() *MockPaymentServiceClientMockRecorder {
    return m.recorder
}

func (m *MockPaymentService) PaymentCreate(ctx context.Context, req *connect.Request[v1.PaymentCreateRequest]) (*connect.Response[v1.PaymentCreateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "PaymentCreate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.PaymentCreateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockPaymentServiceMockRecorder) PaymentCreate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PaymentCreate", reflect.TypeOf((*MockPaymentService)(nil).PaymentCreate), ctx, req)
}

func (m *MockPaymentServiceClient) PaymentCreate(ctx context.Context, req *connect.Request[v1.PaymentCreateRequest]) (*connect.Response[v1.PaymentCreateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "PaymentCreate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.PaymentCreateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockPaymentServiceClientMockRecorder) PaymentCreate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PaymentCreate", reflect.TypeOf((*MockPaymentService)(nil).PaymentCreate), ctx, req)
}

func (m *MockPaymentService) PaymentCancel(ctx context.Context, req *connect.Request[v1.PaymentCancelRequest]) (*connect.Response[v1.PaymentCancelResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "PaymentCancel", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.PaymentCancelResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockPaymentServiceMockRecorder) PaymentCancel(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PaymentCancel", reflect.TypeOf((*MockPaymentService)(nil).PaymentCancel), ctx, req)
}

func (m *MockPaymentServiceClient) PaymentCancel(ctx context.Context, req *connect.Request[v1.PaymentCancelRequest]) (*connect.Response[v1.PaymentCancelResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "PaymentCancel", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.PaymentCancelResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockPaymentServiceClientMockRecorder) PaymentCancel(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PaymentCancel", reflect.TypeOf((*MockPaymentService)(nil).PaymentCancel), ctx, req)
}

func (m *MockPaymentService) PaymentAccept(ctx context.Context, req *connect.Request[v1.PaymentAcceptRequest]) (*connect.Response[v1.PaymentAcceptResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "PaymentAccept", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.PaymentAcceptResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockPaymentServiceMockRecorder) PaymentAccept(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PaymentAccept", reflect.TypeOf((*MockPaymentService)(nil).PaymentAccept), ctx, req)
}

func (m *MockPaymentServiceClient) PaymentAccept(ctx context.Context, req *connect.Request[v1.PaymentAcceptRequest]) (*connect.Response[v1.PaymentAcceptResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "PaymentAccept", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.PaymentAcceptResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockPaymentServiceClientMockRecorder) PaymentAccept(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PaymentAccept", reflect.TypeOf((*MockPaymentService)(nil).PaymentAccept), ctx, req)
}

func (m *MockPaymentService) PaymentReject(ctx context.Context, req *connect.Request[v1.PaymentRejectRequest]) (*connect.Response[v1.PaymentRejectResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "PaymentReject", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.PaymentRejectResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockPaymentServiceMockRecorder) PaymentReject(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PaymentReject", reflect.TypeOf((*MockPaymentService)(nil).PaymentReject), ctx, req)
}

func (m *MockPaymentServiceClient) PaymentReject(ctx context.Context, req *connect.Request[v1.PaymentRejectRequest]) (*connect.Response[v1.PaymentRejectResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "PaymentReject", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.PaymentRejectResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockPaymentServiceClientMockRecorder) PaymentReject(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PaymentReject", reflect.TypeOf((*MockPaymentService)(nil).PaymentReject), ctx, req)
}

func (m *MockPaymentService) PaymentList(ctx context.Context, req *connect.Request[v1.PaymentListRequest]) (*connect.Response[v1.PaymentListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "PaymentList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.PaymentListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockPaymentServiceMockRecorder) PaymentList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PaymentList", reflect.TypeOf((*MockPaymentService)(nil).PaymentList), ctx, req)
}

func (m *MockPaymentServiceClient) PaymentList(ctx context.Context, req *connect.Request[v1.PaymentListRequest]) (*connect.Response[v1.PaymentListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "PaymentList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.PaymentListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockPaymentServiceClientMockRecorder) PaymentList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PaymentList", reflect.TypeOf((*MockPaymentService)(nil).PaymentList), ctx, req)
}

func (m *MockPaymentService) PaymentGet(ctx context.Context, req *connect.Request[v1.PaymentGetRequest]) (*connect.Response[v1.PaymentGetResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "PaymentGet", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.PaymentGetResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockPaymentServiceMockRecorder) PaymentGet(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PaymentGet", reflect.TypeOf((*MockPaymentService)(nil).PaymentGet), ctx, req)
}

func (m *MockPaymentServiceClient) PaymentGet(ctx context.Context, req *connect.Request[v1.PaymentGetRequest]) (*connect.Response[v1.PaymentGetResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "PaymentGet", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.PaymentGetResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockPaymentServiceClientMockRecorder) PaymentGet(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PaymentGet", reflect.TypeOf((*MockPaymentService)(nil).PaymentGet), ctx, req)
}

