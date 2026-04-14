package selling_iface_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/selling_iface/v1"
)

type MockSellingStatService struct {
    ctrl     *gomock.Controller
    recorder *MockSellingStatServiceMockRecorder
}

type MockSellingStatServiceMockRecorder struct {
    mock *MockSellingStatService
}

func NewMockSellingStatService(ctrl *gomock.Controller) *MockSellingStatService {
    mock := &MockSellingStatService{ctrl: ctrl}
    mock.recorder = &MockSellingStatServiceMockRecorder{mock}
    return mock
}

func (m *MockSellingStatService) EXPECT() *MockSellingStatServiceMockRecorder {
    return m.recorder
}

type MockSellingStatServiceClient struct {
    ctrl     *gomock.Controller
    recorder *MockSellingStatServiceClientMockRecorder
}

type MockSellingStatServiceClientMockRecorder struct {
    mock *MockSellingStatServiceClient
}

func NewMockSellingStatServiceClient(ctrl *gomock.Controller) *MockSellingStatServiceClient {
    mock := &MockSellingStatServiceClient{ctrl: ctrl}
    mock.recorder = &MockSellingStatServiceClientMockRecorder{mock}
    return mock
}

func (m *MockSellingStatServiceClient) EXPECT() *MockSellingStatServiceClientMockRecorder {
    return m.recorder
}

func (m *MockSellingStatService) Stat(ctx context.Context, req *connect.Request[v1.StatRequest]) (*connect.Response[v1.StatResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "Stat", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.StatResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockSellingStatServiceMockRecorder) Stat(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stat", reflect.TypeOf((*MockSellingStatService)(nil).Stat), ctx, req)
}

func (m *MockSellingStatServiceClient) Stat(ctx context.Context, req *connect.Request[v1.StatRequest]) (*connect.Response[v1.StatResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "Stat", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.StatResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockSellingStatServiceClientMockRecorder) Stat(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stat", reflect.TypeOf((*MockSellingStatService)(nil).Stat), ctx, req)
}

func (m *MockSellingStatService) StatStream(ctx context.Context, req *connect.Request[v1.StatStreamRequest], stream *connect.ServerStream[v1.StatStreamResponse]) error {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "StatStream", ctx, req, stream)
    ret0, _ := ret[0].(error)
    return ret0
}

func (mr *MockSellingStatServiceMockRecorder) StatStream(ctx, req, stream interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatStream", reflect.TypeOf((*MockSellingStatService)(nil).StatStream), ctx, req, stream)
}

func (m *MockSellingStatServiceClient) StatStream(ctx context.Context, req *connect.Request[v1.StatStreamRequest]) (*connect.ServerStreamForClient[v1.StatStreamResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "StatStream", ctx, req)
    ret0, _ := ret[0].(*connect.ServerStreamForClient[v1.StatStreamResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockSellingStatServiceClientMockRecorder) StatStream(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatStream", reflect.TypeOf((*MockSellingStatService)(nil).StatStream), ctx, req)
}

