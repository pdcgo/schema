package selling_iface_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/selling_iface/v1"
)

type MockSupplierService struct {
    ctrl     *gomock.Controller
    recorder *MockSupplierServiceMockRecorder
}

type MockSupplierServiceMockRecorder struct {
    mock *MockSupplierService
}

func NewMockSupplierService(ctrl *gomock.Controller) *MockSupplierService {
    mock := &MockSupplierService{ctrl: ctrl}
    mock.recorder = &MockSupplierServiceMockRecorder{mock}
    return mock
}

func (m *MockSupplierService) EXPECT() *MockSupplierServiceMockRecorder {
    return m.recorder
}

type MockSupplierServiceClient struct {
    ctrl     *gomock.Controller
    recorder *MockSupplierServiceClientMockRecorder
}

type MockSupplierServiceClientMockRecorder struct {
    mock *MockSupplierServiceClient
}

func NewMockSupplierServiceClient(ctrl *gomock.Controller) *MockSupplierServiceClient {
    mock := &MockSupplierServiceClient{ctrl: ctrl}
    mock.recorder = &MockSupplierServiceClientMockRecorder{mock}
    return mock
}

func (m *MockSupplierServiceClient) EXPECT() *MockSupplierServiceClientMockRecorder {
    return m.recorder
}

func (m *MockSupplierService) SupplierList(ctx context.Context, req *connect.Request[v1.SupplierListRequest]) (*connect.Response[v1.SupplierListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "SupplierList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.SupplierListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockSupplierServiceMockRecorder) SupplierList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SupplierList", reflect.TypeOf((*MockSupplierService)(nil).SupplierList), ctx, req)
}

func (m *MockSupplierServiceClient) SupplierList(ctx context.Context, req *connect.Request[v1.SupplierListRequest]) (*connect.Response[v1.SupplierListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "SupplierList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.SupplierListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockSupplierServiceClientMockRecorder) SupplierList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SupplierList", reflect.TypeOf((*MockSupplierService)(nil).SupplierList), ctx, req)
}

func (m *MockSupplierService) SupplierCreate(ctx context.Context, req *connect.Request[v1.SupplierCreateRequest]) (*connect.Response[v1.SupplierCreateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "SupplierCreate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.SupplierCreateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockSupplierServiceMockRecorder) SupplierCreate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SupplierCreate", reflect.TypeOf((*MockSupplierService)(nil).SupplierCreate), ctx, req)
}

func (m *MockSupplierServiceClient) SupplierCreate(ctx context.Context, req *connect.Request[v1.SupplierCreateRequest]) (*connect.Response[v1.SupplierCreateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "SupplierCreate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.SupplierCreateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockSupplierServiceClientMockRecorder) SupplierCreate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SupplierCreate", reflect.TypeOf((*MockSupplierService)(nil).SupplierCreate), ctx, req)
}

func (m *MockSupplierService) SupplierDelete(ctx context.Context, req *connect.Request[v1.SupplierDeleteRequest]) (*connect.Response[v1.SupplierDeleteResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "SupplierDelete", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.SupplierDeleteResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockSupplierServiceMockRecorder) SupplierDelete(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SupplierDelete", reflect.TypeOf((*MockSupplierService)(nil).SupplierDelete), ctx, req)
}

func (m *MockSupplierServiceClient) SupplierDelete(ctx context.Context, req *connect.Request[v1.SupplierDeleteRequest]) (*connect.Response[v1.SupplierDeleteResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "SupplierDelete", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.SupplierDeleteResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockSupplierServiceClientMockRecorder) SupplierDelete(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SupplierDelete", reflect.TypeOf((*MockSupplierService)(nil).SupplierDelete), ctx, req)
}

func (m *MockSupplierService) SupplierUpdate(ctx context.Context, req *connect.Request[v1.SupplierUpdateRequest]) (*connect.Response[v1.SupplierUpdateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "SupplierUpdate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.SupplierUpdateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockSupplierServiceMockRecorder) SupplierUpdate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SupplierUpdate", reflect.TypeOf((*MockSupplierService)(nil).SupplierUpdate), ctx, req)
}

func (m *MockSupplierServiceClient) SupplierUpdate(ctx context.Context, req *connect.Request[v1.SupplierUpdateRequest]) (*connect.Response[v1.SupplierUpdateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "SupplierUpdate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.SupplierUpdateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockSupplierServiceClientMockRecorder) SupplierUpdate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SupplierUpdate", reflect.TypeOf((*MockSupplierService)(nil).SupplierUpdate), ctx, req)
}

func (m *MockSupplierService) SupplierGet(ctx context.Context, req *connect.Request[v1.SupplierGetRequest]) (*connect.Response[v1.SupplierGetResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "SupplierGet", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.SupplierGetResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockSupplierServiceMockRecorder) SupplierGet(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SupplierGet", reflect.TypeOf((*MockSupplierService)(nil).SupplierGet), ctx, req)
}

func (m *MockSupplierServiceClient) SupplierGet(ctx context.Context, req *connect.Request[v1.SupplierGetRequest]) (*connect.Response[v1.SupplierGetResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "SupplierGet", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.SupplierGetResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockSupplierServiceClientMockRecorder) SupplierGet(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SupplierGet", reflect.TypeOf((*MockSupplierService)(nil).SupplierGet), ctx, req)
}

