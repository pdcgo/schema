package selling_iface_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/selling_iface/v1"
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

func (m *MockRevenueService) HelloWorld(ctx context.Context, req *connect.Request[v1.HelloWorldRequest]) (*connect.Response[v1.HelloWorldResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "HelloWorld", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.HelloWorldResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockRevenueServiceMockRecorder) HelloWorld(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HelloWorld", reflect.TypeOf((*MockRevenueService)(nil).HelloWorld), ctx, req)
}

func (m *MockRevenueServiceClient) HelloWorld(ctx context.Context, req *connect.Request[v1.HelloWorldRequest]) (*connect.Response[v1.HelloWorldResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "HelloWorld", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.HelloWorldResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockRevenueServiceClientMockRecorder) HelloWorld(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HelloWorld", reflect.TypeOf((*MockRevenueService)(nil).HelloWorld), ctx, req)
}

func (m *MockRevenueService) WithdrawalCreate(ctx context.Context, req *connect.Request[v1.WithdrawalCreateRequest]) (*connect.Response[v1.WithdrawalCreateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "WithdrawalCreate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.WithdrawalCreateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockRevenueServiceMockRecorder) WithdrawalCreate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithdrawalCreate", reflect.TypeOf((*MockRevenueService)(nil).WithdrawalCreate), ctx, req)
}

func (m *MockRevenueServiceClient) WithdrawalCreate(ctx context.Context, req *connect.Request[v1.WithdrawalCreateRequest]) (*connect.Response[v1.WithdrawalCreateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "WithdrawalCreate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.WithdrawalCreateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockRevenueServiceClientMockRecorder) WithdrawalCreate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithdrawalCreate", reflect.TypeOf((*MockRevenueService)(nil).WithdrawalCreate), ctx, req)
}

func (m *MockRevenueService) WithdrawalUpdate(ctx context.Context, req *connect.Request[v1.WithdrawalUpdateRequest]) (*connect.Response[v1.WithdrawalUpdateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "WithdrawalUpdate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.WithdrawalUpdateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockRevenueServiceMockRecorder) WithdrawalUpdate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithdrawalUpdate", reflect.TypeOf((*MockRevenueService)(nil).WithdrawalUpdate), ctx, req)
}

func (m *MockRevenueServiceClient) WithdrawalUpdate(ctx context.Context, req *connect.Request[v1.WithdrawalUpdateRequest]) (*connect.Response[v1.WithdrawalUpdateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "WithdrawalUpdate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.WithdrawalUpdateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockRevenueServiceClientMockRecorder) WithdrawalUpdate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithdrawalUpdate", reflect.TypeOf((*MockRevenueService)(nil).WithdrawalUpdate), ctx, req)
}

func (m *MockRevenueService) WithdrawalDelete(ctx context.Context, req *connect.Request[v1.WithdrawalDeleteRequest]) (*connect.Response[v1.WithdrawalDeleteResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "WithdrawalDelete", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.WithdrawalDeleteResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockRevenueServiceMockRecorder) WithdrawalDelete(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithdrawalDelete", reflect.TypeOf((*MockRevenueService)(nil).WithdrawalDelete), ctx, req)
}

func (m *MockRevenueServiceClient) WithdrawalDelete(ctx context.Context, req *connect.Request[v1.WithdrawalDeleteRequest]) (*connect.Response[v1.WithdrawalDeleteResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "WithdrawalDelete", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.WithdrawalDeleteResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockRevenueServiceClientMockRecorder) WithdrawalDelete(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithdrawalDelete", reflect.TypeOf((*MockRevenueService)(nil).WithdrawalDelete), ctx, req)
}

func (m *MockRevenueService) WithdrawalList(ctx context.Context, req *connect.Request[v1.WithdrawalListRequest]) (*connect.Response[v1.WithdrawalListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "WithdrawalList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.WithdrawalListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockRevenueServiceMockRecorder) WithdrawalList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithdrawalList", reflect.TypeOf((*MockRevenueService)(nil).WithdrawalList), ctx, req)
}

func (m *MockRevenueServiceClient) WithdrawalList(ctx context.Context, req *connect.Request[v1.WithdrawalListRequest]) (*connect.Response[v1.WithdrawalListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "WithdrawalList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.WithdrawalListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockRevenueServiceClientMockRecorder) WithdrawalList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithdrawalList", reflect.TypeOf((*MockRevenueService)(nil).WithdrawalList), ctx, req)
}

func (m *MockRevenueService) WithdrawalLogList(ctx context.Context, req *connect.Request[v1.WithdrawalLogListRequest]) (*connect.Response[v1.WithdrawalLogListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "WithdrawalLogList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.WithdrawalLogListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockRevenueServiceMockRecorder) WithdrawalLogList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithdrawalLogList", reflect.TypeOf((*MockRevenueService)(nil).WithdrawalLogList), ctx, req)
}

func (m *MockRevenueServiceClient) WithdrawalLogList(ctx context.Context, req *connect.Request[v1.WithdrawalLogListRequest]) (*connect.Response[v1.WithdrawalLogListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "WithdrawalLogList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.WithdrawalLogListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockRevenueServiceClientMockRecorder) WithdrawalLogList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithdrawalLogList", reflect.TypeOf((*MockRevenueService)(nil).WithdrawalLogList), ctx, req)
}

