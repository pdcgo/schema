package payment_iface_mock

import (
    "context"
    "github.com/stretchr/testify/mock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/payment_iface/v1"
)

type MockPaymentService struct {
    mock.Mock
}

func (m *MockPaymentService) PaymentCreate(ctx context.Context, req *connect.Request[v1.PaymentCreateRequest]) (*connect.Response[v1.PaymentCreateResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.PaymentCreateResponse]), args.Error(1)
}

func (m *MockPaymentService) PaymentCancel(ctx context.Context, req *connect.Request[v1.PaymentCancelRequest]) (*connect.Response[v1.PaymentCancelResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.PaymentCancelResponse]), args.Error(1)
}

func (m *MockPaymentService) PaymentAccept(ctx context.Context, req *connect.Request[v1.PaymentAcceptRequest]) (*connect.Response[v1.PaymentAcceptResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.PaymentAcceptResponse]), args.Error(1)
}

func (m *MockPaymentService) PaymentReject(ctx context.Context, req *connect.Request[v1.PaymentRejectRequest]) (*connect.Response[v1.PaymentRejectResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.PaymentRejectResponse]), args.Error(1)
}

func (m *MockPaymentService) PaymentList(ctx context.Context, req *connect.Request[v1.PaymentListRequest]) (*connect.Response[v1.PaymentListResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.PaymentListResponse]), args.Error(1)
}

func (m *MockPaymentService) PaymentGet(ctx context.Context, req *connect.Request[v1.PaymentGetRequest]) (*connect.Response[v1.PaymentGetResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.PaymentGetResponse]), args.Error(1)
}

