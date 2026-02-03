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

func (m *MockConfigurationLimitService) OweDefaultLimitGet(ctx context.Context, req *connect.Request[v1.OweDefaultLimitGetRequest]) (*connect.Response[v1.OweDefaultLimitGetResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OweDefaultLimitGet", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OweDefaultLimitGetResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockConfigurationLimitServiceMockRecorder) OweDefaultLimitGet(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OweDefaultLimitGet", reflect.TypeOf((*MockConfigurationLimitService)(nil).OweDefaultLimitGet), ctx, req)
}

func (m *MockConfigurationLimitServiceClient) OweDefaultLimitGet(ctx context.Context, req *connect.Request[v1.OweDefaultLimitGetRequest]) (*connect.Response[v1.OweDefaultLimitGetResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OweDefaultLimitGet", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OweDefaultLimitGetResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockConfigurationLimitServiceClientMockRecorder) OweDefaultLimitGet(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OweDefaultLimitGet", reflect.TypeOf((*MockConfigurationLimitService)(nil).OweDefaultLimitGet), ctx, req)
}

func (m *MockConfigurationLimitService) OweDefaultLimitEdit(ctx context.Context, req *connect.Request[v1.OweDefaultLimitEditRequest]) (*connect.Response[v1.OweDefaultLimitEditResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OweDefaultLimitEdit", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OweDefaultLimitEditResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockConfigurationLimitServiceMockRecorder) OweDefaultLimitEdit(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OweDefaultLimitEdit", reflect.TypeOf((*MockConfigurationLimitService)(nil).OweDefaultLimitEdit), ctx, req)
}

func (m *MockConfigurationLimitServiceClient) OweDefaultLimitEdit(ctx context.Context, req *connect.Request[v1.OweDefaultLimitEditRequest]) (*connect.Response[v1.OweDefaultLimitEditResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OweDefaultLimitEdit", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OweDefaultLimitEditResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockConfigurationLimitServiceClientMockRecorder) OweDefaultLimitEdit(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OweDefaultLimitEdit", reflect.TypeOf((*MockConfigurationLimitService)(nil).OweDefaultLimitEdit), ctx, req)
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

func (m *MockConfigurationLimitService) OweLimitCustomDelete(ctx context.Context, req *connect.Request[v1.OweLimitCustomDeleteRequest]) (*connect.Response[v1.OweLimitCustomDeleteResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OweLimitCustomDelete", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OweLimitCustomDeleteResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockConfigurationLimitServiceMockRecorder) OweLimitCustomDelete(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OweLimitCustomDelete", reflect.TypeOf((*MockConfigurationLimitService)(nil).OweLimitCustomDelete), ctx, req)
}

func (m *MockConfigurationLimitServiceClient) OweLimitCustomDelete(ctx context.Context, req *connect.Request[v1.OweLimitCustomDeleteRequest]) (*connect.Response[v1.OweLimitCustomDeleteResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "OweLimitCustomDelete", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.OweLimitCustomDeleteResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockConfigurationLimitServiceClientMockRecorder) OweLimitCustomDelete(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OweLimitCustomDelete", reflect.TypeOf((*MockConfigurationLimitService)(nil).OweLimitCustomDelete), ctx, req)
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

func (m *MockConfigurationLimitService) SellingLimitProfitGet(ctx context.Context, req *connect.Request[v1.SellingLimitProfitGetRequest]) (*connect.Response[v1.SellingLimitProfitGetResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "SellingLimitProfitGet", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.SellingLimitProfitGetResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockConfigurationLimitServiceMockRecorder) SellingLimitProfitGet(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SellingLimitProfitGet", reflect.TypeOf((*MockConfigurationLimitService)(nil).SellingLimitProfitGet), ctx, req)
}

func (m *MockConfigurationLimitServiceClient) SellingLimitProfitGet(ctx context.Context, req *connect.Request[v1.SellingLimitProfitGetRequest]) (*connect.Response[v1.SellingLimitProfitGetResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "SellingLimitProfitGet", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.SellingLimitProfitGetResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockConfigurationLimitServiceClientMockRecorder) SellingLimitProfitGet(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SellingLimitProfitGet", reflect.TypeOf((*MockConfigurationLimitService)(nil).SellingLimitProfitGet), ctx, req)
}

func (m *MockConfigurationLimitService) SellingLimitProfitEdit(ctx context.Context, req *connect.Request[v1.SellingLimitProfitEditRequest]) (*connect.Response[v1.SellingLimitProfitEditResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "SellingLimitProfitEdit", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.SellingLimitProfitEditResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockConfigurationLimitServiceMockRecorder) SellingLimitProfitEdit(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SellingLimitProfitEdit", reflect.TypeOf((*MockConfigurationLimitService)(nil).SellingLimitProfitEdit), ctx, req)
}

func (m *MockConfigurationLimitServiceClient) SellingLimitProfitEdit(ctx context.Context, req *connect.Request[v1.SellingLimitProfitEditRequest]) (*connect.Response[v1.SellingLimitProfitEditResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "SellingLimitProfitEdit", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.SellingLimitProfitEditResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockConfigurationLimitServiceClientMockRecorder) SellingLimitProfitEdit(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SellingLimitProfitEdit", reflect.TypeOf((*MockConfigurationLimitService)(nil).SellingLimitProfitEdit), ctx, req)
}

