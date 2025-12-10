package common_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/common/v1"
)

type MockShopService struct {
    ctrl     *gomock.Controller
    recorder *MockShopServiceMockRecorder
}

type MockShopServiceMockRecorder struct {
    mock *MockShopService
}

func NewMockShopService(ctrl *gomock.Controller) *MockShopService {
    mock := &MockShopService{ctrl: ctrl}
    mock.recorder = &MockShopServiceMockRecorder{mock}
    return mock
}

func (m *MockShopService) EXPECT() *MockShopServiceMockRecorder {
    return m.recorder
}

func (m *MockShopService) PublicShopIDs(ctx context.Context, req *connect.Request[v1.PublicShopIDsRequest]) (*connect.Response[v1.PublicShopIDsResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "PublicShopIDs", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.PublicShopIDsResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockShopServiceMockRecorder) PublicShopIDs(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublicShopIDs", reflect.TypeOf((*MockShopService)(nil).PublicShopIDs), ctx, req)
}

func (m *MockShopService) PublicShopList(ctx context.Context, req *connect.Request[v1.PublicShopListRequest]) (*connect.Response[v1.PublicShopListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "PublicShopList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.PublicShopListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockShopServiceMockRecorder) PublicShopList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublicShopList", reflect.TypeOf((*MockShopService)(nil).PublicShopList), ctx, req)
}

