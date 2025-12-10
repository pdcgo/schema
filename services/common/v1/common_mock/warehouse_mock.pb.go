package common_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/common/v1"
)

type MockWarehouseService struct {
    ctrl     *gomock.Controller
    recorder *MockWarehouseServiceMockRecorder
}

type MockWarehouseServiceMockRecorder struct {
    mock *MockWarehouseService
}

func NewMockWarehouseService(ctrl *gomock.Controller) *MockWarehouseService {
    mock := &MockWarehouseService{ctrl: ctrl}
    mock.recorder = &MockWarehouseServiceMockRecorder{mock}
    return mock
}

func (m *MockWarehouseService) EXPECT() *MockWarehouseServiceMockRecorder {
    return m.recorder
}

type MockWarehouseServiceClient struct {
    ctrl     *gomock.Controller
    recorder *MockWarehouseServiceClientMockRecorder
}

type MockWarehouseServiceClientMockRecorder struct {
    mock *MockWarehouseServiceClient
}

func NewMockWarehouseServiceClient(ctrl *gomock.Controller) *MockWarehouseServiceClient {
    mock := &MockWarehouseServiceClient{ctrl: ctrl}
    mock.recorder = &MockWarehouseServiceClientMockRecorder{mock}
    return mock
}

func (m *MockWarehouseServiceClient) EXPECT() *MockWarehouseServiceClientMockRecorder {
    return m.recorder
}

func (m *MockWarehouseService) PublicWarehouseIDs(ctx context.Context, req *connect.Request[v1.PublicWarehouseIDsRequest]) (*connect.Response[v1.PublicWarehouseIDsResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "PublicWarehouseIDs", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.PublicWarehouseIDsResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockWarehouseServiceMockRecorder) PublicWarehouseIDs(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublicWarehouseIDs", reflect.TypeOf((*MockWarehouseService)(nil).PublicWarehouseIDs), ctx, req)
}

func (m *MockWarehouseServiceClient) PublicWarehouseIDs(ctx context.Context, req *connect.Request[v1.PublicWarehouseIDsRequest]) (*connect.Response[v1.PublicWarehouseIDsResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "PublicWarehouseIDs", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.PublicWarehouseIDsResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockWarehouseServiceClientMockRecorder) PublicWarehouseIDs(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublicWarehouseIDs", reflect.TypeOf((*MockWarehouseService)(nil).PublicWarehouseIDs), ctx, req)
}

