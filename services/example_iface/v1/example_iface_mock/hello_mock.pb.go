package example_iface_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/example_iface/v1"
)

type MockExampleHelloService struct {
    ctrl     *gomock.Controller
    recorder *MockExampleHelloServiceMockRecorder
}

type MockExampleHelloServiceMockRecorder struct {
    mock *MockExampleHelloService
}

func NewMockExampleHelloService(ctrl *gomock.Controller) *MockExampleHelloService {
    mock := &MockExampleHelloService{ctrl: ctrl}
    mock.recorder = &MockExampleHelloServiceMockRecorder{mock}
    return mock
}

func (m *MockExampleHelloService) EXPECT() *MockExampleHelloServiceMockRecorder {
    return m.recorder
}

type MockExampleHelloServiceClient struct {
    ctrl     *gomock.Controller
    recorder *MockExampleHelloServiceClientMockRecorder
}

type MockExampleHelloServiceClientMockRecorder struct {
    mock *MockExampleHelloServiceClient
}

func NewMockExampleHelloServiceClient(ctrl *gomock.Controller) *MockExampleHelloServiceClient {
    mock := &MockExampleHelloServiceClient{ctrl: ctrl}
    mock.recorder = &MockExampleHelloServiceClientMockRecorder{mock}
    return mock
}

func (m *MockExampleHelloServiceClient) EXPECT() *MockExampleHelloServiceClientMockRecorder {
    return m.recorder
}

func (m *MockExampleHelloService) CreateToken(ctx context.Context, req *connect.Request[v1.CreateTokenRequest]) (*connect.Response[v1.CreateTokenResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "CreateToken", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.CreateTokenResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockExampleHelloServiceMockRecorder) CreateToken(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateToken", reflect.TypeOf((*MockExampleHelloService)(nil).CreateToken), ctx, req)
}

func (m *MockExampleHelloServiceClient) CreateToken(ctx context.Context, req *connect.Request[v1.CreateTokenRequest]) (*connect.Response[v1.CreateTokenResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "CreateToken", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.CreateTokenResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockExampleHelloServiceClientMockRecorder) CreateToken(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateToken", reflect.TypeOf((*MockExampleHelloService)(nil).CreateToken), ctx, req)
}

func (m *MockExampleHelloService) HelloWorld(ctx context.Context, req *connect.Request[v1.HelloWorldRequest]) (*connect.Response[v1.HelloWorldResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "HelloWorld", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.HelloWorldResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockExampleHelloServiceMockRecorder) HelloWorld(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HelloWorld", reflect.TypeOf((*MockExampleHelloService)(nil).HelloWorld), ctx, req)
}

func (m *MockExampleHelloServiceClient) HelloWorld(ctx context.Context, req *connect.Request[v1.HelloWorldRequest]) (*connect.Response[v1.HelloWorldResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "HelloWorld", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.HelloWorldResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockExampleHelloServiceClientMockRecorder) HelloWorld(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HelloWorld", reflect.TypeOf((*MockExampleHelloService)(nil).HelloWorld), ctx, req)
}

func (m *MockExampleHelloService) HelloAdmin(ctx context.Context, req *connect.Request[v1.HelloAdminRequest]) (*connect.Response[v1.HelloAdminResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "HelloAdmin", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.HelloAdminResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockExampleHelloServiceMockRecorder) HelloAdmin(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HelloAdmin", reflect.TypeOf((*MockExampleHelloService)(nil).HelloAdmin), ctx, req)
}

func (m *MockExampleHelloServiceClient) HelloAdmin(ctx context.Context, req *connect.Request[v1.HelloAdminRequest]) (*connect.Response[v1.HelloAdminResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "HelloAdmin", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.HelloAdminResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockExampleHelloServiceClientMockRecorder) HelloAdmin(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HelloAdmin", reflect.TypeOf((*MockExampleHelloService)(nil).HelloAdmin), ctx, req)
}

