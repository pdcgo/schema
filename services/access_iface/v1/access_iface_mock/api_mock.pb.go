package access_iface_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/access_iface/v1"
)

type MockHelloService struct {
    ctrl     *gomock.Controller
    recorder *MockHelloServiceMockRecorder
}

type MockHelloServiceMockRecorder struct {
    mock *MockHelloService
}

func NewMockHelloService(ctrl *gomock.Controller) *MockHelloService {
    mock := &MockHelloService{ctrl: ctrl}
    mock.recorder = &MockHelloServiceMockRecorder{mock}
    return mock
}

func (m *MockHelloService) EXPECT() *MockHelloServiceMockRecorder {
    return m.recorder
}

type MockHelloServiceClient struct {
    ctrl     *gomock.Controller
    recorder *MockHelloServiceClientMockRecorder
}

type MockHelloServiceClientMockRecorder struct {
    mock *MockHelloServiceClient
}

func NewMockHelloServiceClient(ctrl *gomock.Controller) *MockHelloServiceClient {
    mock := &MockHelloServiceClient{ctrl: ctrl}
    mock.recorder = &MockHelloServiceClientMockRecorder{mock}
    return mock
}

func (m *MockHelloServiceClient) EXPECT() *MockHelloServiceClientMockRecorder {
    return m.recorder
}

func (m *MockHelloService) Hello(ctx context.Context, req *connect.Request[v1.HelloRequest]) (*connect.Response[v1.HelloResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "Hello", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.HelloResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockHelloServiceMockRecorder) Hello(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hello", reflect.TypeOf((*MockHelloService)(nil).Hello), ctx, req)
}

func (m *MockHelloServiceClient) Hello(ctx context.Context, req *connect.Request[v1.HelloRequest]) (*connect.Response[v1.HelloResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "Hello", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.HelloResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockHelloServiceClientMockRecorder) Hello(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hello", reflect.TypeOf((*MockHelloService)(nil).Hello), ctx, req)
}

func (m *MockHelloService) HelloClientStream(ctx context.Context, stream *connect.ClientStream[v1.HelloClientStreamRequest]) (*connect.Response[v1.HelloClientStreamResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "HelloClientStream", ctx, stream)
    ret0, _ := ret[0].(*connect.Response[v1.HelloClientStreamResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockHelloServiceMockRecorder) HelloClientStream(ctx, stream interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HelloClientStream", reflect.TypeOf((*MockHelloService)(nil).HelloClientStream), ctx, stream)
}

func (m *MockHelloServiceClient) HelloClientStream(ctx context.Context) *connect.ClientStreamForClient[v1.HelloClientStreamRequest, v1.HelloClientStreamResponse] {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "HelloClientStream", ctx)
    ret0, _ := ret[0].(*connect.ClientStreamForClient[v1.HelloClientStreamRequest, v1.HelloClientStreamResponse])
    return ret0
}

func (mr *MockHelloServiceClientMockRecorder) HelloClientStream(ctx context.Context) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HelloClientStream", reflect.TypeOf((*MockHelloService)(nil).HelloClientStream), ctx)
}

func (m *MockHelloService) HelloServerStream(ctx context.Context, req *connect.Request[v1.HelloServerStreamRequest], stream *connect.ServerStream[v1.HelloServerStreamResponse]) error {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "HelloServerStream", ctx, req, stream)
    ret0, _ := ret[0].(error)
    return ret0
}

func (mr *MockHelloServiceMockRecorder) HelloServerStream(ctx, req, stream interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HelloServerStream", reflect.TypeOf((*MockHelloService)(nil).HelloServerStream), ctx, req, stream)
}

func (m *MockHelloServiceClient) HelloServerStream(ctx context.Context, req *connect.Request[v1.HelloServerStreamRequest]) (*connect.ServerStreamForClient[v1.HelloServerStreamResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "HelloServerStream", ctx, req)
    ret0, _ := ret[0].(*connect.ServerStreamForClient[v1.HelloServerStreamResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockHelloServiceClientMockRecorder) HelloServerStream(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HelloServerStream", reflect.TypeOf((*MockHelloService)(nil).HelloServerStream), ctx, req)
}

func (m *MockHelloService) HelloBidiStream(ctx context.Context, stream *connect.BidiStream[v1.HelloBidiStreamRequest, v1.HelloBidiStreamResponse]) error {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "HelloBidiStream", ctx, stream)
    ret0, _ := ret[0].(error)
    return ret0
}

func (mr *MockHelloServiceMockRecorder) HelloBidiStream(ctx, stream interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HelloBidiStream", reflect.TypeOf((*MockHelloService)(nil).HelloBidiStream), ctx, stream)
}

func (m *MockHelloServiceClient) HelloBidiStream(ctx context.Context) *connect.BidiStreamForClient[v1.HelloBidiStreamRequest, v1.HelloBidiStreamResponse] {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "HelloBidiStream", ctx)
    ret0, _ := ret[0].(*connect.BidiStreamForClient[v1.HelloBidiStreamRequest, v1.HelloBidiStreamResponse])
    return ret0
}

func (mr *MockHelloServiceClientMockRecorder) HelloBidiStream(ctx, stream interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HelloBidiStream", reflect.TypeOf((*MockHelloService)(nil).HelloBidiStream), ctx)
}

