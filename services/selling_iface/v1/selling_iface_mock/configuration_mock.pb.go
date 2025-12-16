package selling_iface_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/selling_iface/v1"
)

type MockConfigurationLimitService struct {
    ctrl     *gomock.Controller
    recorder *MockConfigurationLimitServiceMockRecorder
}

type MockConfigurationLimitServiceMockRecorder struct {
    mock *MockConfigurationLimitService
}

func NewMockConfigurationLimitService(ctrl *gomock.Controller) *MockConfigurationLimitService {
    mock := &MockConfigurationLimitService{ctrl: ctrl}
    mock.recorder = &MockConfigurationLimitServiceMockRecorder{mock}
    return mock
}

func (m *MockConfigurationLimitService) EXPECT() *MockConfigurationLimitServiceMockRecorder {
    return m.recorder
}

type MockConfigurationLimitServiceClient struct {
    ctrl     *gomock.Controller
    recorder *MockConfigurationLimitServiceClientMockRecorder
}

type MockConfigurationLimitServiceClientMockRecorder struct {
    mock *MockConfigurationLimitServiceClient
}

func NewMockConfigurationLimitServiceClient(ctrl *gomock.Controller) *MockConfigurationLimitServiceClient {
    mock := &MockConfigurationLimitServiceClient{ctrl: ctrl}
    mock.recorder = &MockConfigurationLimitServiceClientMockRecorder{mock}
    return mock
}

func (m *MockConfigurationLimitServiceClient) EXPECT() *MockConfigurationLimitServiceClientMockRecorder {
    return m.recorder
}

func (m *MockConfigurationLimitService) LimitInvoice(ctx context.Context, req *connect.Request[v1.LimitInvoiceRequest]) (*connect.Response[v1.LimitInvoiceResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "LimitInvoice", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.LimitInvoiceResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockConfigurationLimitServiceMockRecorder) LimitInvoice(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LimitInvoice", reflect.TypeOf((*MockConfigurationLimitService)(nil).LimitInvoice), ctx, req)
}

func (m *MockConfigurationLimitServiceClient) LimitInvoice(ctx context.Context, req *connect.Request[v1.LimitInvoiceRequest]) (*connect.Response[v1.LimitInvoiceResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "LimitInvoice", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.LimitInvoiceResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockConfigurationLimitServiceClientMockRecorder) LimitInvoice(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LimitInvoice", reflect.TypeOf((*MockConfigurationLimitService)(nil).LimitInvoice), ctx, req)
}

func (m *MockConfigurationLimitService) OweDefaultLimit(ctx context.Context, req *connect.Request[v1.OweDefaultLimitRequest]) (*connect.Response[v1.OweDefaultLimitResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OweDefaultLimit", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OweDefaultLimitResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockConfigurationLimitServiceMockRecorder) OweDefaultLimit(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OweDefaultLimit", reflect.TypeOf((*MockConfigurationLimitService)(nil).OweDefaultLimit), ctx, req)
}

func (m *MockConfigurationLimitServiceClient) OweDefaultLimit(ctx context.Context, req *connect.Request[v1.OweDefaultLimitRequest]) (*connect.Response[v1.OweDefaultLimitResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OweDefaultLimit", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OweDefaultLimitResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockConfigurationLimitServiceClientMockRecorder) OweDefaultLimit(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OweDefaultLimit", reflect.TypeOf((*MockConfigurationLimitService)(nil).OweDefaultLimit), ctx, req)
}

func (m *MockConfigurationLimitService) OweLimitCustomCreate(ctx context.Context, req *connect.Request[v1.OweLimitCustomCreateRequest]) (*connect.Response[v1.OweLimitCustomCreateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OweLimitCustomCreate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OweLimitCustomCreateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockConfigurationLimitServiceMockRecorder) OweLimitCustomCreate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OweLimitCustomCreate", reflect.TypeOf((*MockConfigurationLimitService)(nil).OweLimitCustomCreate), ctx, req)
}

func (m *MockConfigurationLimitServiceClient) OweLimitCustomCreate(ctx context.Context, req *connect.Request[v1.OweLimitCustomCreateRequest]) (*connect.Response[v1.OweLimitCustomCreateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OweLimitCustomCreate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OweLimitCustomCreateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockConfigurationLimitServiceClientMockRecorder) OweLimitCustomCreate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OweLimitCustomCreate", reflect.TypeOf((*MockConfigurationLimitService)(nil).OweLimitCustomCreate), ctx, req)
}

func (m *MockConfigurationLimitService) OweLimitCustomList(ctx context.Context, req *connect.Request[v1.OweLimitCustomListRequest]) (*connect.Response[v1.OweLimitCustomListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OweLimitCustomList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OweLimitCustomListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockConfigurationLimitServiceMockRecorder) OweLimitCustomList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OweLimitCustomList", reflect.TypeOf((*MockConfigurationLimitService)(nil).OweLimitCustomList), ctx, req)
}

func (m *MockConfigurationLimitServiceClient) OweLimitCustomList(ctx context.Context, req *connect.Request[v1.OweLimitCustomListRequest]) (*connect.Response[v1.OweLimitCustomListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OweLimitCustomList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OweLimitCustomListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockConfigurationLimitServiceClientMockRecorder) OweLimitCustomList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OweLimitCustomList", reflect.TypeOf((*MockConfigurationLimitService)(nil).OweLimitCustomList), ctx, req)
}

func (m *MockConfigurationLimitService) OweLimitDelete(ctx context.Context, req *connect.Request[v1.OweLimitDeleteRequest]) (*connect.Response[v1.OweLimitDeleteResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OweLimitDelete", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OweLimitDeleteResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockConfigurationLimitServiceMockRecorder) OweLimitDelete(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OweLimitDelete", reflect.TypeOf((*MockConfigurationLimitService)(nil).OweLimitDelete), ctx, req)
}

func (m *MockConfigurationLimitServiceClient) OweLimitDelete(ctx context.Context, req *connect.Request[v1.OweLimitDeleteRequest]) (*connect.Response[v1.OweLimitDeleteResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OweLimitDelete", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OweLimitDeleteResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockConfigurationLimitServiceClientMockRecorder) OweLimitDelete(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OweLimitDelete", reflect.TypeOf((*MockConfigurationLimitService)(nil).OweLimitDelete), ctx, req)
}

func (m *MockConfigurationLimitService) OweLimitCustomByIDs(ctx context.Context, req *connect.Request[v1.OweLimitCustomByIDsRequest]) (*connect.Response[v1.OweLimitCustomByIDsResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OweLimitCustomByIDs", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OweLimitCustomByIDsResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockConfigurationLimitServiceMockRecorder) OweLimitCustomByIDs(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OweLimitCustomByIDs", reflect.TypeOf((*MockConfigurationLimitService)(nil).OweLimitCustomByIDs), ctx, req)
}

func (m *MockConfigurationLimitServiceClient) OweLimitCustomByIDs(ctx context.Context, req *connect.Request[v1.OweLimitCustomByIDsRequest]) (*connect.Response[v1.OweLimitCustomByIDsResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OweLimitCustomByIDs", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OweLimitCustomByIDsResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockConfigurationLimitServiceClientMockRecorder) OweLimitCustomByIDs(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OweLimitCustomByIDs", reflect.TypeOf((*MockConfigurationLimitService)(nil).OweLimitCustomByIDs), ctx, req)
}

func (m *MockConfigurationLimitService) CheckOweLimit(ctx context.Context, req *connect.Request[v1.CheckOweLimitRequest]) (*connect.Response[v1.CheckOweLimitResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "CheckOweLimit", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.CheckOweLimitResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockConfigurationLimitServiceMockRecorder) CheckOweLimit(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckOweLimit", reflect.TypeOf((*MockConfigurationLimitService)(nil).CheckOweLimit), ctx, req)
}

func (m *MockConfigurationLimitServiceClient) CheckOweLimit(ctx context.Context, req *connect.Request[v1.CheckOweLimitRequest]) (*connect.Response[v1.CheckOweLimitResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "CheckOweLimit", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.CheckOweLimitResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockConfigurationLimitServiceClientMockRecorder) CheckOweLimit(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckOweLimit", reflect.TypeOf((*MockConfigurationLimitService)(nil).CheckOweLimit), ctx, req)
}

