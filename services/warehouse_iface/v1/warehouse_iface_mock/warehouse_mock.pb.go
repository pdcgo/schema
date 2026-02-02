package warehouse_iface_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/warehouse_iface/v1"
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

func (m *MockWarehouseService) WarehouseIDs(ctx context.Context, req *connect.Request[v1.WarehouseIDsRequest]) (*connect.Response[v1.WarehouseIDsResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "WarehouseIDs", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.WarehouseIDsResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockWarehouseServiceMockRecorder) WarehouseIDs(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WarehouseIDs", reflect.TypeOf((*MockWarehouseService)(nil).WarehouseIDs), ctx, req)
}

func (m *MockWarehouseServiceClient) WarehouseIDs(ctx context.Context, req *connect.Request[v1.WarehouseIDsRequest]) (*connect.Response[v1.WarehouseIDsResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "WarehouseIDs", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.WarehouseIDsResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockWarehouseServiceClientMockRecorder) WarehouseIDs(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WarehouseIDs", reflect.TypeOf((*MockWarehouseService)(nil).WarehouseIDs), ctx, req)
}

func (m *MockWarehouseService) WarehouseList(ctx context.Context, req *connect.Request[v1.WarehouseListRequest]) (*connect.Response[v1.WarehouseListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "WarehouseList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.WarehouseListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockWarehouseServiceMockRecorder) WarehouseList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WarehouseList", reflect.TypeOf((*MockWarehouseService)(nil).WarehouseList), ctx, req)
}

func (m *MockWarehouseServiceClient) WarehouseList(ctx context.Context, req *connect.Request[v1.WarehouseListRequest]) (*connect.Response[v1.WarehouseListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "WarehouseList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.WarehouseListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockWarehouseServiceClientMockRecorder) WarehouseList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WarehouseList", reflect.TypeOf((*MockWarehouseService)(nil).WarehouseList), ctx, req)
}

func (m *MockWarehouseService) TeamWarehouseReturnInfo(ctx context.Context, req *connect.Request[v1.TeamWarehouseReturnInfoRequest]) (*connect.Response[v1.TeamWarehouseReturnInfoResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TeamWarehouseReturnInfo", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.TeamWarehouseReturnInfoResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockWarehouseServiceMockRecorder) TeamWarehouseReturnInfo(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeamWarehouseReturnInfo", reflect.TypeOf((*MockWarehouseService)(nil).TeamWarehouseReturnInfo), ctx, req)
}

func (m *MockWarehouseServiceClient) TeamWarehouseReturnInfo(ctx context.Context, req *connect.Request[v1.TeamWarehouseReturnInfoRequest]) (*connect.Response[v1.TeamWarehouseReturnInfoResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TeamWarehouseReturnInfo", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.TeamWarehouseReturnInfoResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockWarehouseServiceClientMockRecorder) TeamWarehouseReturnInfo(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeamWarehouseReturnInfo", reflect.TypeOf((*MockWarehouseService)(nil).TeamWarehouseReturnInfo), ctx, req)
}

