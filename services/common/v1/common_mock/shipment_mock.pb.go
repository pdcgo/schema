package common_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/common/v1"
)

type MockShipmentService struct {
    ctrl     *gomock.Controller
    recorder *MockShipmentServiceMockRecorder
}

type MockShipmentServiceMockRecorder struct {
    mock *MockShipmentService
}

func NewMockShipmentService(ctrl *gomock.Controller) *MockShipmentService {
    mock := &MockShipmentService{ctrl: ctrl}
    mock.recorder = &MockShipmentServiceMockRecorder{mock}
    return mock
}

func (m *MockShipmentService) EXPECT() *MockShipmentServiceMockRecorder {
    return m.recorder
}

type MockShipmentServiceClient struct {
    ctrl     *gomock.Controller
    recorder *MockShipmentServiceClientMockRecorder
}

type MockShipmentServiceClientMockRecorder struct {
    mock *MockShipmentServiceClient
}

func NewMockShipmentServiceClient(ctrl *gomock.Controller) *MockShipmentServiceClient {
    mock := &MockShipmentServiceClient{ctrl: ctrl}
    mock.recorder = &MockShipmentServiceClientMockRecorder{mock}
    return mock
}

func (m *MockShipmentServiceClient) EXPECT() *MockShipmentServiceClientMockRecorder {
    return m.recorder
}

func (m *MockShipmentService) PublicShipmentIDs(ctx context.Context, req *connect.Request[v1.PublicShipmentIDsRequest]) (*connect.Response[v1.PublicShipmentIDsResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "PublicShipmentIDs", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.PublicShipmentIDsResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockShipmentServiceMockRecorder) PublicShipmentIDs(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublicShipmentIDs", reflect.TypeOf((*MockShipmentService)(nil).PublicShipmentIDs), ctx, req)
}

func (m *MockShipmentServiceClient) PublicShipmentIDs(ctx context.Context, req *connect.Request[v1.PublicShipmentIDsRequest]) (*connect.Response[v1.PublicShipmentIDsResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "PublicShipmentIDs", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.PublicShipmentIDsResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockShipmentServiceClientMockRecorder) PublicShipmentIDs(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublicShipmentIDs", reflect.TypeOf((*MockShipmentService)(nil).PublicShipmentIDs), ctx, req)
}

func (m *MockShipmentService) PublicShipmentList(ctx context.Context, req *connect.Request[v1.PublicShipmentListRequest]) (*connect.Response[v1.PublicShipmentListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "PublicShipmentList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.PublicShipmentListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockShipmentServiceMockRecorder) PublicShipmentList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublicShipmentList", reflect.TypeOf((*MockShipmentService)(nil).PublicShipmentList), ctx, req)
}

func (m *MockShipmentServiceClient) PublicShipmentList(ctx context.Context, req *connect.Request[v1.PublicShipmentListRequest]) (*connect.Response[v1.PublicShipmentListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "PublicShipmentList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.PublicShipmentListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockShipmentServiceClientMockRecorder) PublicShipmentList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublicShipmentList", reflect.TypeOf((*MockShipmentService)(nil).PublicShipmentList), ctx, req)
}

