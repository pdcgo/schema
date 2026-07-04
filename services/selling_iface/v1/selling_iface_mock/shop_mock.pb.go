package selling_iface_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/selling_iface/v1"
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

type MockShopServiceClient struct {
    ctrl     *gomock.Controller
    recorder *MockShopServiceClientMockRecorder
}

type MockShopServiceClientMockRecorder struct {
    mock *MockShopServiceClient
}

func NewMockShopServiceClient(ctrl *gomock.Controller) *MockShopServiceClient {
    mock := &MockShopServiceClient{ctrl: ctrl}
    mock.recorder = &MockShopServiceClientMockRecorder{mock}
    return mock
}

func (m *MockShopServiceClient) EXPECT() *MockShopServiceClientMockRecorder {
    return m.recorder
}

func (m *MockShopService) ShopList(ctx context.Context, req *connect.Request[v1.ShopListRequest]) (*connect.Response[v1.ShopListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ShopList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ShopListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockShopServiceMockRecorder) ShopList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShopList", reflect.TypeOf((*MockShopService)(nil).ShopList), ctx, req)
}

func (m *MockShopServiceClient) ShopList(ctx context.Context, req *connect.Request[v1.ShopListRequest]) (*connect.Response[v1.ShopListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ShopList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ShopListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockShopServiceClientMockRecorder) ShopList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShopList", reflect.TypeOf((*MockShopService)(nil).ShopList), ctx, req)
}

func (m *MockShopService) ShopCreate(ctx context.Context, req *connect.Request[v1.ShopCreateRequest]) (*connect.Response[v1.ShopCreateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ShopCreate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ShopCreateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockShopServiceMockRecorder) ShopCreate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShopCreate", reflect.TypeOf((*MockShopService)(nil).ShopCreate), ctx, req)
}

func (m *MockShopServiceClient) ShopCreate(ctx context.Context, req *connect.Request[v1.ShopCreateRequest]) (*connect.Response[v1.ShopCreateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ShopCreate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ShopCreateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockShopServiceClientMockRecorder) ShopCreate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShopCreate", reflect.TypeOf((*MockShopService)(nil).ShopCreate), ctx, req)
}

func (m *MockShopService) ShopUpdate(ctx context.Context, req *connect.Request[v1.ShopUpdateRequest]) (*connect.Response[v1.ShopUpdateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ShopUpdate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ShopUpdateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockShopServiceMockRecorder) ShopUpdate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShopUpdate", reflect.TypeOf((*MockShopService)(nil).ShopUpdate), ctx, req)
}

func (m *MockShopServiceClient) ShopUpdate(ctx context.Context, req *connect.Request[v1.ShopUpdateRequest]) (*connect.Response[v1.ShopUpdateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ShopUpdate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ShopUpdateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockShopServiceClientMockRecorder) ShopUpdate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShopUpdate", reflect.TypeOf((*MockShopService)(nil).ShopUpdate), ctx, req)
}

func (m *MockShopService) ShopDelete(ctx context.Context, req *connect.Request[v1.ShopDeleteRequest]) (*connect.Response[v1.ShopDeleteResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ShopDelete", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ShopDeleteResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockShopServiceMockRecorder) ShopDelete(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShopDelete", reflect.TypeOf((*MockShopService)(nil).ShopDelete), ctx, req)
}

func (m *MockShopServiceClient) ShopDelete(ctx context.Context, req *connect.Request[v1.ShopDeleteRequest]) (*connect.Response[v1.ShopDeleteResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ShopDelete", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ShopDeleteResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockShopServiceClientMockRecorder) ShopDelete(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShopDelete", reflect.TypeOf((*MockShopService)(nil).ShopDelete), ctx, req)
}

func (m *MockShopService) ShopUserList(ctx context.Context, req *connect.Request[v1.ShopUserListRequest]) (*connect.Response[v1.ShopUserListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ShopUserList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ShopUserListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockShopServiceMockRecorder) ShopUserList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShopUserList", reflect.TypeOf((*MockShopService)(nil).ShopUserList), ctx, req)
}

func (m *MockShopServiceClient) ShopUserList(ctx context.Context, req *connect.Request[v1.ShopUserListRequest]) (*connect.Response[v1.ShopUserListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ShopUserList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ShopUserListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockShopServiceClientMockRecorder) ShopUserList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShopUserList", reflect.TypeOf((*MockShopService)(nil).ShopUserList), ctx, req)
}

func (m *MockShopService) ShopUserUpdate(ctx context.Context, req *connect.Request[v1.ShopUserUpdateRequest]) (*connect.Response[v1.ShopUserUpdateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ShopUserUpdate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ShopUserUpdateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockShopServiceMockRecorder) ShopUserUpdate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShopUserUpdate", reflect.TypeOf((*MockShopService)(nil).ShopUserUpdate), ctx, req)
}

func (m *MockShopServiceClient) ShopUserUpdate(ctx context.Context, req *connect.Request[v1.ShopUserUpdateRequest]) (*connect.Response[v1.ShopUserUpdateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ShopUserUpdate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ShopUserUpdateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockShopServiceClientMockRecorder) ShopUserUpdate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShopUserUpdate", reflect.TypeOf((*MockShopService)(nil).ShopUserUpdate), ctx, req)
}

