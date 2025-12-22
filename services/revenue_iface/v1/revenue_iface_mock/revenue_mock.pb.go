package revenue_iface_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/revenue_iface/v1"
)

type MockRevenueService struct {
    ctrl     *gomock.Controller
    recorder *MockRevenueServiceMockRecorder
}

type MockRevenueServiceMockRecorder struct {
    mock *MockRevenueService
}

func NewMockRevenueService(ctrl *gomock.Controller) *MockRevenueService {
    mock := &MockRevenueService{ctrl: ctrl}
    mock.recorder = &MockRevenueServiceMockRecorder{mock}
    return mock
}

func (m *MockRevenueService) EXPECT() *MockRevenueServiceMockRecorder {
    return m.recorder
}

type MockRevenueServiceClient struct {
    ctrl     *gomock.Controller
    recorder *MockRevenueServiceClientMockRecorder
}

type MockRevenueServiceClientMockRecorder struct {
    mock *MockRevenueServiceClient
}

func NewMockRevenueServiceClient(ctrl *gomock.Controller) *MockRevenueServiceClient {
    mock := &MockRevenueServiceClient{ctrl: ctrl}
    mock.recorder = &MockRevenueServiceClientMockRecorder{mock}
    return mock
}

func (m *MockRevenueServiceClient) EXPECT() *MockRevenueServiceClientMockRecorder {
    return m.recorder
}

func (m *MockRevenueService) OnOrder(ctx context.Context, req *connect.Request[v1.OnOrderRequest]) (*connect.Response[v1.OnOrderResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OnOrder", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OnOrderResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockRevenueServiceMockRecorder) OnOrder(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnOrder", reflect.TypeOf((*MockRevenueService)(nil).OnOrder), ctx, req)
}

func (m *MockRevenueServiceClient) OnOrder(ctx context.Context, req *connect.Request[v1.OnOrderRequest]) (*connect.Response[v1.OnOrderResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OnOrder", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OnOrderResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockRevenueServiceClientMockRecorder) OnOrder(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnOrder", reflect.TypeOf((*MockRevenueService)(nil).OnOrder), ctx, req)
}

func (m *MockRevenueService) OrderCancel(ctx context.Context, req *connect.Request[v1.OrderCancelRequest]) (*connect.Response[v1.OrderCancelResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OrderCancel", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OrderCancelResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockRevenueServiceMockRecorder) OrderCancel(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderCancel", reflect.TypeOf((*MockRevenueService)(nil).OrderCancel), ctx, req)
}

func (m *MockRevenueServiceClient) OrderCancel(ctx context.Context, req *connect.Request[v1.OrderCancelRequest]) (*connect.Response[v1.OrderCancelResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OrderCancel", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OrderCancelResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockRevenueServiceClientMockRecorder) OrderCancel(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderCancel", reflect.TypeOf((*MockRevenueService)(nil).OrderCancel), ctx, req)
}

func (m *MockRevenueService) SellingReceivableAdjustment(ctx context.Context, req *connect.Request[v1.SellingReceivableAdjustmentRequest]) (*connect.Response[v1.SellingReceivableAdjustmentResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "SellingReceivableAdjustment", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.SellingReceivableAdjustmentResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockRevenueServiceMockRecorder) SellingReceivableAdjustment(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SellingReceivableAdjustment", reflect.TypeOf((*MockRevenueService)(nil).SellingReceivableAdjustment), ctx, req)
}

func (m *MockRevenueServiceClient) SellingReceivableAdjustment(ctx context.Context, req *connect.Request[v1.SellingReceivableAdjustmentRequest]) (*connect.Response[v1.SellingReceivableAdjustmentResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "SellingReceivableAdjustment", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.SellingReceivableAdjustmentResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockRevenueServiceClientMockRecorder) SellingReceivableAdjustment(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SellingReceivableAdjustment", reflect.TypeOf((*MockRevenueService)(nil).SellingReceivableAdjustment), ctx, req)
}

func (m *MockRevenueService) OrderReturnAsync(ctx context.Context, req *connect.Request[v1.OrderReturnAsyncRequest]) (*connect.Response[v1.OrderReturnAsyncResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OrderReturnAsync", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OrderReturnAsyncResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockRevenueServiceMockRecorder) OrderReturnAsync(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderReturnAsync", reflect.TypeOf((*MockRevenueService)(nil).OrderReturnAsync), ctx, req)
}

func (m *MockRevenueServiceClient) OrderReturnAsync(ctx context.Context, req *connect.Request[v1.OrderReturnAsyncRequest]) (*connect.Response[v1.OrderReturnAsyncResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OrderReturnAsync", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OrderReturnAsyncResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockRevenueServiceClientMockRecorder) OrderReturnAsync(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderReturnAsync", reflect.TypeOf((*MockRevenueService)(nil).OrderReturnAsync), ctx, req)
}

func (m *MockRevenueService) OrderReturn(ctx context.Context, req *connect.Request[v1.OrderReturnRequest]) (*connect.Response[v1.OrderReturnResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OrderReturn", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OrderReturnResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockRevenueServiceMockRecorder) OrderReturn(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderReturn", reflect.TypeOf((*MockRevenueService)(nil).OrderReturn), ctx, req)
}

func (m *MockRevenueServiceClient) OrderReturn(ctx context.Context, req *connect.Request[v1.OrderReturnRequest]) (*connect.Response[v1.OrderReturnResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OrderReturn", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OrderReturnResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockRevenueServiceClientMockRecorder) OrderReturn(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderReturn", reflect.TypeOf((*MockRevenueService)(nil).OrderReturn), ctx, req)
}

func (m *MockRevenueService) OrderCompleted(ctx context.Context, req *connect.Request[v1.OrderCompletedRequest]) (*connect.Response[v1.OrderCompletedResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OrderCompleted", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OrderCompletedResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockRevenueServiceMockRecorder) OrderCompleted(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderCompleted", reflect.TypeOf((*MockRevenueService)(nil).OrderCompleted), ctx, req)
}

func (m *MockRevenueServiceClient) OrderCompleted(ctx context.Context, req *connect.Request[v1.OrderCompletedRequest]) (*connect.Response[v1.OrderCompletedResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OrderCompleted", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OrderCompletedResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockRevenueServiceClientMockRecorder) OrderCompleted(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderCompleted", reflect.TypeOf((*MockRevenueService)(nil).OrderCompleted), ctx, req)
}

func (m *MockRevenueService) RevenueAdjustment(ctx context.Context, req *connect.Request[v1.RevenueAdjustmentRequest]) (*connect.Response[v1.RevenueAdjustmentResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "RevenueAdjustment", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.RevenueAdjustmentResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockRevenueServiceMockRecorder) RevenueAdjustment(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevenueAdjustment", reflect.TypeOf((*MockRevenueService)(nil).RevenueAdjustment), ctx, req)
}

func (m *MockRevenueServiceClient) RevenueAdjustment(ctx context.Context, req *connect.Request[v1.RevenueAdjustmentRequest]) (*connect.Response[v1.RevenueAdjustmentResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "RevenueAdjustment", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.RevenueAdjustmentResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockRevenueServiceClientMockRecorder) RevenueAdjustment(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevenueAdjustment", reflect.TypeOf((*MockRevenueService)(nil).RevenueAdjustment), ctx, req)
}

func (m *MockRevenueService) Withdrawal(ctx context.Context, req *connect.Request[v1.WithdrawalRequest]) (*connect.Response[v1.WithdrawalResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "Withdrawal", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.WithdrawalResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockRevenueServiceMockRecorder) Withdrawal(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Withdrawal", reflect.TypeOf((*MockRevenueService)(nil).Withdrawal), ctx, req)
}

func (m *MockRevenueServiceClient) Withdrawal(ctx context.Context, req *connect.Request[v1.WithdrawalRequest]) (*connect.Response[v1.WithdrawalResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "Withdrawal", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.WithdrawalResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockRevenueServiceClientMockRecorder) Withdrawal(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Withdrawal", reflect.TypeOf((*MockRevenueService)(nil).Withdrawal), ctx, req)
}

func (m *MockRevenueService) RevenueStream(ctx context.Context, stream *connect.ClientStream[v1.RevenueStreamRequest]) (*connect.Response[v1.RevenueStreamResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "RevenueStream", ctx, stream)
    ret0, _ := ret[0].(*connect.Response[v1.RevenueStreamResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockRevenueServiceMockRecorder) RevenueStream(ctx, stream interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevenueStream", reflect.TypeOf((*MockRevenueService)(nil).RevenueStream), ctx, stream)
}

func (m *MockRevenueServiceClient) RevenueStream(ctx context.Context) *connect.ClientStreamForClient[v1.RevenueStreamRequest, v1.RevenueStreamResponse] {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "RevenueStream", ctx)
    ret0, _ := ret[0].(*connect.ClientStreamForClient[v1.RevenueStreamRequest, v1.RevenueStreamResponse])
    return ret0
}

func (mr *MockRevenueServiceClientMockRecorder) RevenueStream(ctx context.Context) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevenueStream", reflect.TypeOf((*MockRevenueService)(nil).RevenueStream), ctx)
}

func (m *MockRevenueService) RevenueOther(ctx context.Context, req *connect.Request[v1.RevenueOtherRequest]) (*connect.Response[v1.RevenueOtherResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "RevenueOther", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.RevenueOtherResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockRevenueServiceMockRecorder) RevenueOther(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevenueOther", reflect.TypeOf((*MockRevenueService)(nil).RevenueOther), ctx, req)
}

func (m *MockRevenueServiceClient) RevenueOther(ctx context.Context, req *connect.Request[v1.RevenueOtherRequest]) (*connect.Response[v1.RevenueOtherResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "RevenueOther", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.RevenueOtherResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockRevenueServiceClientMockRecorder) RevenueOther(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevenueOther", reflect.TypeOf((*MockRevenueService)(nil).RevenueOther), ctx, req)
}

func (m *MockRevenueService) SellingExpenseOther(ctx context.Context, req *connect.Request[v1.SellingExpenseOtherRequest]) (*connect.Response[v1.SellingExpenseOtherResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "SellingExpenseOther", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.SellingExpenseOtherResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockRevenueServiceMockRecorder) SellingExpenseOther(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SellingExpenseOther", reflect.TypeOf((*MockRevenueService)(nil).SellingExpenseOther), ctx, req)
}

func (m *MockRevenueServiceClient) SellingExpenseOther(ctx context.Context, req *connect.Request[v1.SellingExpenseOtherRequest]) (*connect.Response[v1.SellingExpenseOtherResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "SellingExpenseOther", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.SellingExpenseOtherResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockRevenueServiceClientMockRecorder) SellingExpenseOther(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SellingExpenseOther", reflect.TypeOf((*MockRevenueService)(nil).SellingExpenseOther), ctx, req)
}

