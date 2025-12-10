package withdrawal_iface_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v2 "github.com/pdcgo/schema/services/withdrawal_iface/v2"
)

type MockWithdrawalService struct {
    ctrl     *gomock.Controller
    recorder *MockWithdrawalServiceMockRecorder
}

type MockWithdrawalServiceMockRecorder struct {
    mock *MockWithdrawalService
}

func NewMockWithdrawalService(ctrl *gomock.Controller) *MockWithdrawalService {
    mock := &MockWithdrawalService{ctrl: ctrl}
    mock.recorder = &MockWithdrawalServiceMockRecorder{mock}
    return mock
}

func (m *MockWithdrawalService) EXPECT() *MockWithdrawalServiceMockRecorder {
    return m.recorder
}

type MockWithdrawalServiceClient struct {
    ctrl     *gomock.Controller
    recorder *MockWithdrawalServiceClientMockRecorder
}

type MockWithdrawalServiceClientMockRecorder struct {
    mock *MockWithdrawalServiceClient
}

func NewMockWithdrawalServiceClient(ctrl *gomock.Controller) *MockWithdrawalServiceClient {
    mock := &MockWithdrawalServiceClient{ctrl: ctrl}
    mock.recorder = &MockWithdrawalServiceClientMockRecorder{mock}
    return mock
}

func (m *MockWithdrawalServiceClient) EXPECT() *MockWithdrawalServiceClientMockRecorder {
    return m.recorder
}

func (m *MockWithdrawalService) SubmitWithdrawal(ctx context.Context, req *connect.Request[v2.SubmitWithdrawalRequest], stream *connect.ServerStream[v2.SubmitWithdrawalResponse]) error {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "SubmitWithdrawal", ctx, req, stream)
    ret0, _ := ret[0].(error)
    return ret0
}

func (mr *MockWithdrawalServiceMockRecorder) SubmitWithdrawal(ctx, req, stream interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitWithdrawal", reflect.TypeOf((*MockWithdrawalService)(nil).SubmitWithdrawal), ctx, req, stream)
}

func (m *MockWithdrawalServiceClient) SubmitWithdrawal(ctx context.Context, req *connect.Request[v2.SubmitWithdrawalRequest]) (*connect.ServerStreamForClient[v2.SubmitWithdrawalResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "SubmitWithdrawal", ctx, req)
    ret0, _ := ret[0].(*connect.ServerStreamForClient[v2.SubmitWithdrawalResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockWithdrawalServiceClientMockRecorder) SubmitWithdrawal(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitWithdrawal", reflect.TypeOf((*MockWithdrawalService)(nil).SubmitWithdrawal), ctx, req)
}

func (m *MockWithdrawalService) SubmitWithdrawalTiktok(ctx context.Context, req *connect.Request[v2.SubmitWithdrawalTiktokRequest], stream *connect.ServerStream[v2.SubmitWithdrawalTiktokResponse]) error {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "SubmitWithdrawalTiktok", ctx, req, stream)
    ret0, _ := ret[0].(error)
    return ret0
}

func (mr *MockWithdrawalServiceMockRecorder) SubmitWithdrawalTiktok(ctx, req, stream interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitWithdrawalTiktok", reflect.TypeOf((*MockWithdrawalService)(nil).SubmitWithdrawalTiktok), ctx, req, stream)
}

func (m *MockWithdrawalServiceClient) SubmitWithdrawalTiktok(ctx context.Context, req *connect.Request[v2.SubmitWithdrawalTiktokRequest]) (*connect.ServerStreamForClient[v2.SubmitWithdrawalTiktokResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "SubmitWithdrawalTiktok", ctx, req)
    ret0, _ := ret[0].(*connect.ServerStreamForClient[v2.SubmitWithdrawalTiktokResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockWithdrawalServiceClientMockRecorder) SubmitWithdrawalTiktok(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitWithdrawalTiktok", reflect.TypeOf((*MockWithdrawalService)(nil).SubmitWithdrawalTiktok), ctx, req)
}

func (m *MockWithdrawalService) SubmitWithdrawalShopee(ctx context.Context, req *connect.Request[v2.SubmitWithdrawalShopeeRequest], stream *connect.ServerStream[v2.SubmitWithdrawalShopeeResponse]) error {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "SubmitWithdrawalShopee", ctx, req, stream)
    ret0, _ := ret[0].(error)
    return ret0
}

func (mr *MockWithdrawalServiceMockRecorder) SubmitWithdrawalShopee(ctx, req, stream interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitWithdrawalShopee", reflect.TypeOf((*MockWithdrawalService)(nil).SubmitWithdrawalShopee), ctx, req, stream)
}

func (m *MockWithdrawalServiceClient) SubmitWithdrawalShopee(ctx context.Context, req *connect.Request[v2.SubmitWithdrawalShopeeRequest]) (*connect.ServerStreamForClient[v2.SubmitWithdrawalShopeeResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "SubmitWithdrawalShopee", ctx, req)
    ret0, _ := ret[0].(*connect.ServerStreamForClient[v2.SubmitWithdrawalShopeeResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockWithdrawalServiceClientMockRecorder) SubmitWithdrawalShopee(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitWithdrawalShopee", reflect.TypeOf((*MockWithdrawalService)(nil).SubmitWithdrawalShopee), ctx, req)
}

