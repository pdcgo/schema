package order_iface_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/order_iface/v1"
)

type MockOrderService struct {
    ctrl     *gomock.Controller
    recorder *MockOrderServiceMockRecorder
}

type MockOrderServiceMockRecorder struct {
    mock *MockOrderService
}

func NewMockOrderService(ctrl *gomock.Controller) *MockOrderService {
    mock := &MockOrderService{ctrl: ctrl}
    mock.recorder = &MockOrderServiceMockRecorder{mock}
    return mock
}

func (m *MockOrderService) EXPECT() *MockOrderServiceMockRecorder {
    return m.recorder
}

type MockOrderServiceClient struct {
    ctrl     *gomock.Controller
    recorder *MockOrderServiceClientMockRecorder
}

type MockOrderServiceClientMockRecorder struct {
    mock *MockOrderServiceClient
}

func NewMockOrderServiceClient(ctrl *gomock.Controller) *MockOrderServiceClient {
    mock := &MockOrderServiceClient{ctrl: ctrl}
    mock.recorder = &MockOrderServiceClientMockRecorder{mock}
    return mock
}

func (m *MockOrderServiceClient) EXPECT() *MockOrderServiceClientMockRecorder {
    return m.recorder
}

func (m *MockOrderService) OrderDraftCreate(ctx context.Context, req *connect.Request[v1.OrderDraftCreateRequest]) (*connect.Response[v1.OrderDraftCreateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OrderDraftCreate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OrderDraftCreateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockOrderServiceMockRecorder) OrderDraftCreate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderDraftCreate", reflect.TypeOf((*MockOrderService)(nil).OrderDraftCreate), ctx, req)
}

func (m *MockOrderServiceClient) OrderDraftCreate(ctx context.Context, req *connect.Request[v1.OrderDraftCreateRequest]) (*connect.Response[v1.OrderDraftCreateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OrderDraftCreate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OrderDraftCreateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockOrderServiceClientMockRecorder) OrderDraftCreate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderDraftCreate", reflect.TypeOf((*MockOrderService)(nil).OrderDraftCreate), ctx, req)
}

func (m *MockOrderService) OrderDraftDelete(ctx context.Context, req *connect.Request[v1.OrderDraftDeleteRequest]) (*connect.Response[v1.OrderDraftDeleteResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OrderDraftDelete", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OrderDraftDeleteResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockOrderServiceMockRecorder) OrderDraftDelete(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderDraftDelete", reflect.TypeOf((*MockOrderService)(nil).OrderDraftDelete), ctx, req)
}

func (m *MockOrderServiceClient) OrderDraftDelete(ctx context.Context, req *connect.Request[v1.OrderDraftDeleteRequest]) (*connect.Response[v1.OrderDraftDeleteResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OrderDraftDelete", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OrderDraftDeleteResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockOrderServiceClientMockRecorder) OrderDraftDelete(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderDraftDelete", reflect.TypeOf((*MockOrderService)(nil).OrderDraftDelete), ctx, req)
}

func (m *MockOrderService) OrderDraftList(ctx context.Context, req *connect.Request[v1.OrderDraftListRequest]) (*connect.Response[v1.OrderDraftListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OrderDraftList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OrderDraftListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockOrderServiceMockRecorder) OrderDraftList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderDraftList", reflect.TypeOf((*MockOrderService)(nil).OrderDraftList), ctx, req)
}

func (m *MockOrderServiceClient) OrderDraftList(ctx context.Context, req *connect.Request[v1.OrderDraftListRequest]) (*connect.Response[v1.OrderDraftListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OrderDraftList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OrderDraftListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockOrderServiceClientMockRecorder) OrderDraftList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderDraftList", reflect.TypeOf((*MockOrderService)(nil).OrderDraftList), ctx, req)
}

func (m *MockOrderService) OrderDraftGet(ctx context.Context, req *connect.Request[v1.OrderDraftGetRequest]) (*connect.Response[v1.OrderDraftGetResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OrderDraftGet", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OrderDraftGetResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockOrderServiceMockRecorder) OrderDraftGet(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderDraftGet", reflect.TypeOf((*MockOrderService)(nil).OrderDraftGet), ctx, req)
}

func (m *MockOrderServiceClient) OrderDraftGet(ctx context.Context, req *connect.Request[v1.OrderDraftGetRequest]) (*connect.Response[v1.OrderDraftGetResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OrderDraftGet", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OrderDraftGetResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockOrderServiceClientMockRecorder) OrderDraftGet(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderDraftGet", reflect.TypeOf((*MockOrderService)(nil).OrderDraftGet), ctx, req)
}

func (m *MockOrderService) OrderCreate(ctx context.Context, req *connect.Request[v1.OrderCreateRequest]) (*connect.Response[v1.OrderCreateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OrderCreate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OrderCreateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockOrderServiceMockRecorder) OrderCreate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderCreate", reflect.TypeOf((*MockOrderService)(nil).OrderCreate), ctx, req)
}

func (m *MockOrderServiceClient) OrderCreate(ctx context.Context, req *connect.Request[v1.OrderCreateRequest]) (*connect.Response[v1.OrderCreateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OrderCreate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OrderCreateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockOrderServiceClientMockRecorder) OrderCreate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderCreate", reflect.TypeOf((*MockOrderService)(nil).OrderCreate), ctx, req)
}

func (m *MockOrderService) OrderFundSet(ctx context.Context, stream *connect.ClientStream[v1.OrderFundSetRequest]) (*connect.Response[v1.OrderFundSetResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OrderFundSet", ctx, stream)
    ret0, _ := ret[0].(*connect.Response[v1.OrderFundSetResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockOrderServiceMockRecorder) OrderFundSet(ctx, stream interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderFundSet", reflect.TypeOf((*MockOrderService)(nil).OrderFundSet), ctx, stream)
}

func (m *MockOrderServiceClient) OrderFundSet(ctx context.Context) *connect.ClientStreamForClient[v1.OrderFundSetRequest, v1.OrderFundSetResponse] {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OrderFundSet", ctx)
    ret0, _ := ret[0].(*connect.ClientStreamForClient[v1.OrderFundSetRequest, v1.OrderFundSetResponse])
    return ret0
}

func (mr *MockOrderServiceClientMockRecorder) OrderFundSet(ctx context.Context) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderFundSet", reflect.TypeOf((*MockOrderService)(nil).OrderFundSet), ctx)
}

func (m *MockOrderService) OrderTagRemove(ctx context.Context, req *connect.Request[v1.OrderTagRemoveRequest]) (*connect.Response[v1.OrderTagRemoveResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OrderTagRemove", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OrderTagRemoveResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockOrderServiceMockRecorder) OrderTagRemove(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderTagRemove", reflect.TypeOf((*MockOrderService)(nil).OrderTagRemove), ctx, req)
}

func (m *MockOrderServiceClient) OrderTagRemove(ctx context.Context, req *connect.Request[v1.OrderTagRemoveRequest]) (*connect.Response[v1.OrderTagRemoveResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OrderTagRemove", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OrderTagRemoveResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockOrderServiceClientMockRecorder) OrderTagRemove(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderTagRemove", reflect.TypeOf((*MockOrderService)(nil).OrderTagRemove), ctx, req)
}

func (m *MockOrderService) OrderTagAdd(ctx context.Context, req *connect.Request[v1.OrderTagAddRequest]) (*connect.Response[v1.OrderTagAddResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OrderTagAdd", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OrderTagAddResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockOrderServiceMockRecorder) OrderTagAdd(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderTagAdd", reflect.TypeOf((*MockOrderService)(nil).OrderTagAdd), ctx, req)
}

func (m *MockOrderServiceClient) OrderTagAdd(ctx context.Context, req *connect.Request[v1.OrderTagAddRequest]) (*connect.Response[v1.OrderTagAddResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OrderTagAdd", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OrderTagAddResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockOrderServiceClientMockRecorder) OrderTagAdd(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderTagAdd", reflect.TypeOf((*MockOrderService)(nil).OrderTagAdd), ctx, req)
}

func (m *MockOrderService) ChangeOrderRefID(ctx context.Context, req *connect.Request[v1.ChangeOrderRefIDRequest]) (*connect.Response[v1.ChangeOrderRefIDResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ChangeOrderRefID", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ChangeOrderRefIDResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockOrderServiceMockRecorder) ChangeOrderRefID(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeOrderRefID", reflect.TypeOf((*MockOrderService)(nil).ChangeOrderRefID), ctx, req)
}

func (m *MockOrderServiceClient) ChangeOrderRefID(ctx context.Context, req *connect.Request[v1.ChangeOrderRefIDRequest]) (*connect.Response[v1.ChangeOrderRefIDResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ChangeOrderRefID", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ChangeOrderRefIDResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockOrderServiceClientMockRecorder) ChangeOrderRefID(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeOrderRefID", reflect.TypeOf((*MockOrderService)(nil).ChangeOrderRefID), ctx, req)
}

func (m *MockOrderService) ChangeEstRevenue(ctx context.Context, req *connect.Request[v1.ChangeEstRevenueRequest]) (*connect.Response[v1.ChangeEstRevenueResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ChangeEstRevenue", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ChangeEstRevenueResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockOrderServiceMockRecorder) ChangeEstRevenue(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeEstRevenue", reflect.TypeOf((*MockOrderService)(nil).ChangeEstRevenue), ctx, req)
}

func (m *MockOrderServiceClient) ChangeEstRevenue(ctx context.Context, req *connect.Request[v1.ChangeEstRevenueRequest]) (*connect.Response[v1.ChangeEstRevenueResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ChangeEstRevenue", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ChangeEstRevenueResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockOrderServiceClientMockRecorder) ChangeEstRevenue(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeEstRevenue", reflect.TypeOf((*MockOrderService)(nil).ChangeEstRevenue), ctx, req)
}

func (m *MockOrderService) OrderCompleted(ctx context.Context, req *connect.Request[v1.OrderCompletedRequest]) (*connect.Response[v1.OrderCompletedResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OrderCompleted", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OrderCompletedResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockOrderServiceMockRecorder) OrderCompleted(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderCompleted", reflect.TypeOf((*MockOrderService)(nil).OrderCompleted), ctx, req)
}

func (m *MockOrderServiceClient) OrderCompleted(ctx context.Context, req *connect.Request[v1.OrderCompletedRequest]) (*connect.Response[v1.OrderCompletedResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OrderCompleted", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OrderCompletedResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockOrderServiceClientMockRecorder) OrderCompleted(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderCompleted", reflect.TypeOf((*MockOrderService)(nil).OrderCompleted), ctx, req)
}

func (m *MockOrderService) OrderReturnArrived(ctx context.Context, req *connect.Request[v1.OrderReturnArrivedRequest]) (*connect.Response[v1.OrderReturnArrivedResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OrderReturnArrived", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OrderReturnArrivedResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockOrderServiceMockRecorder) OrderReturnArrived(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderReturnArrived", reflect.TypeOf((*MockOrderService)(nil).OrderReturnArrived), ctx, req)
}

func (m *MockOrderServiceClient) OrderReturnArrived(ctx context.Context, req *connect.Request[v1.OrderReturnArrivedRequest]) (*connect.Response[v1.OrderReturnArrivedResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OrderReturnArrived", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OrderReturnArrivedResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockOrderServiceClientMockRecorder) OrderReturnArrived(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderReturnArrived", reflect.TypeOf((*MockOrderService)(nil).OrderReturnArrived), ctx, req)
}

func (m *MockOrderService) OrderList(ctx context.Context, req *connect.Request[v1.OrderListRequest], stream *connect.ServerStream[v1.OrderListResponse]) error {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OrderList", ctx, req, stream)
    ret0, _ := ret[0].(error)
    return ret0
}

func (mr *MockOrderServiceMockRecorder) OrderList(ctx, req, stream interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderList", reflect.TypeOf((*MockOrderService)(nil).OrderList), ctx, req, stream)
}

func (m *MockOrderServiceClient) OrderList(ctx context.Context, req *connect.Request[v1.OrderListRequest]) (*connect.ServerStreamForClient[v1.OrderListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OrderList", ctx, req)
    ret0, _ := ret[0].(*connect.ServerStreamForClient[v1.OrderListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockOrderServiceClientMockRecorder) OrderList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderList", reflect.TypeOf((*MockOrderService)(nil).OrderList), ctx, req)
}

func (m *MockOrderService) OrderOverview(ctx context.Context, req *connect.Request[v1.OrderOverviewRequest]) (*connect.Response[v1.OrderOverviewResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OrderOverview", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OrderOverviewResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockOrderServiceMockRecorder) OrderOverview(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderOverview", reflect.TypeOf((*MockOrderService)(nil).OrderOverview), ctx, req)
}

func (m *MockOrderServiceClient) OrderOverview(ctx context.Context, req *connect.Request[v1.OrderOverviewRequest]) (*connect.Response[v1.OrderOverviewResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OrderOverview", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OrderOverviewResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockOrderServiceClientMockRecorder) OrderOverview(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderOverview", reflect.TypeOf((*MockOrderService)(nil).OrderOverview), ctx, req)
}

func (m *MockOrderService) MpPaymentCreate(ctx context.Context, req *connect.Request[v1.MpPaymentCreateRequest]) (*connect.Response[v1.MpPaymentCreateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "MpPaymentCreate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.MpPaymentCreateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockOrderServiceMockRecorder) MpPaymentCreate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MpPaymentCreate", reflect.TypeOf((*MockOrderService)(nil).MpPaymentCreate), ctx, req)
}

func (m *MockOrderServiceClient) MpPaymentCreate(ctx context.Context, req *connect.Request[v1.MpPaymentCreateRequest]) (*connect.Response[v1.MpPaymentCreateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "MpPaymentCreate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.MpPaymentCreateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockOrderServiceClientMockRecorder) MpPaymentCreate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MpPaymentCreate", reflect.TypeOf((*MockOrderService)(nil).MpPaymentCreate), ctx, req)
}

func (m *MockOrderService) MpPaymentOrderList(ctx context.Context, req *connect.Request[v1.MpPaymentOrderListRequest]) (*connect.Response[v1.MpPaymentOrderListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "MpPaymentOrderList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.MpPaymentOrderListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockOrderServiceMockRecorder) MpPaymentOrderList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MpPaymentOrderList", reflect.TypeOf((*MockOrderService)(nil).MpPaymentOrderList), ctx, req)
}

func (m *MockOrderServiceClient) MpPaymentOrderList(ctx context.Context, req *connect.Request[v1.MpPaymentOrderListRequest]) (*connect.Response[v1.MpPaymentOrderListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "MpPaymentOrderList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.MpPaymentOrderListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockOrderServiceClientMockRecorder) MpPaymentOrderList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MpPaymentOrderList", reflect.TypeOf((*MockOrderService)(nil).MpPaymentOrderList), ctx, req)
}

func (m *MockOrderService) MpPaymentDelete(ctx context.Context, req *connect.Request[v1.MpPaymentDeleteRequest]) (*connect.Response[v1.MpPaymentDeleteResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "MpPaymentDelete", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.MpPaymentDeleteResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockOrderServiceMockRecorder) MpPaymentDelete(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MpPaymentDelete", reflect.TypeOf((*MockOrderService)(nil).MpPaymentDelete), ctx, req)
}

func (m *MockOrderServiceClient) MpPaymentDelete(ctx context.Context, req *connect.Request[v1.MpPaymentDeleteRequest]) (*connect.Response[v1.MpPaymentDeleteResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "MpPaymentDelete", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.MpPaymentDeleteResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockOrderServiceClientMockRecorder) MpPaymentDelete(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MpPaymentDelete", reflect.TypeOf((*MockOrderService)(nil).MpPaymentDelete), ctx, req)
}

