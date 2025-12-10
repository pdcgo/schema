package accounting_iface_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/accounting_iface/v1"
)

type MockAccountService struct {
    ctrl     *gomock.Controller
    recorder *MockAccountServiceMockRecorder
}

type MockAccountServiceMockRecorder struct {
    mock *MockAccountService
}

func NewMockAccountService(ctrl *gomock.Controller) *MockAccountService {
    mock := &MockAccountService{ctrl: ctrl}
    mock.recorder = &MockAccountServiceMockRecorder{mock}
    return mock
}

func (m *MockAccountService) EXPECT() *MockAccountServiceMockRecorder {
    return m.recorder
}

type MockAccountServiceClient struct {
    ctrl     *gomock.Controller
    recorder *MockAccountServiceClientMockRecorder
}

type MockAccountServiceClientMockRecorder struct {
    mock *MockAccountServiceClient
}

func NewMockAccountServiceClient(ctrl *gomock.Controller) *MockAccountServiceClient {
    mock := &MockAccountServiceClient{ctrl: ctrl}
    mock.recorder = &MockAccountServiceClientMockRecorder{mock}
    return mock
}

func (m *MockAccountServiceClient) EXPECT() *MockAccountServiceClientMockRecorder {
    return m.recorder
}

func (m *MockAccountService) AccountCreate(ctx context.Context, req *connect.Request[v1.AccountCreateRequest]) (*connect.Response[v1.AccountCreateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "AccountCreate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.AccountCreateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockAccountServiceMockRecorder) AccountCreate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccountCreate", reflect.TypeOf((*MockAccountService)(nil).AccountCreate), ctx, req)
}

func (m *MockAccountServiceClient) AccountCreate(ctx context.Context, req *connect.Request[v1.AccountCreateRequest]) (*connect.Response[v1.AccountCreateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "AccountCreate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.AccountCreateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockAccountServiceClientMockRecorder) AccountCreate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccountCreate", reflect.TypeOf((*MockAccountService)(nil).AccountCreate), ctx, req)
}

func (m *MockAccountService) AccountList(ctx context.Context, req *connect.Request[v1.AccountListRequest]) (*connect.Response[v1.AccountListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "AccountList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.AccountListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockAccountServiceMockRecorder) AccountList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccountList", reflect.TypeOf((*MockAccountService)(nil).AccountList), ctx, req)
}

func (m *MockAccountServiceClient) AccountList(ctx context.Context, req *connect.Request[v1.AccountListRequest]) (*connect.Response[v1.AccountListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "AccountList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.AccountListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockAccountServiceClientMockRecorder) AccountList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccountList", reflect.TypeOf((*MockAccountService)(nil).AccountList), ctx, req)
}

func (m *MockAccountService) AccountDelete(ctx context.Context, req *connect.Request[v1.AccountDeleteRequest]) (*connect.Response[v1.AccountDeleteResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "AccountDelete", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.AccountDeleteResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockAccountServiceMockRecorder) AccountDelete(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccountDelete", reflect.TypeOf((*MockAccountService)(nil).AccountDelete), ctx, req)
}

func (m *MockAccountServiceClient) AccountDelete(ctx context.Context, req *connect.Request[v1.AccountDeleteRequest]) (*connect.Response[v1.AccountDeleteResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "AccountDelete", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.AccountDeleteResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockAccountServiceClientMockRecorder) AccountDelete(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccountDelete", reflect.TypeOf((*MockAccountService)(nil).AccountDelete), ctx, req)
}

func (m *MockAccountService) AccountUpdate(ctx context.Context, req *connect.Request[v1.AccountUpdateRequest]) (*connect.Response[v1.AccountUpdateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "AccountUpdate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.AccountUpdateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockAccountServiceMockRecorder) AccountUpdate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccountUpdate", reflect.TypeOf((*MockAccountService)(nil).AccountUpdate), ctx, req)
}

func (m *MockAccountServiceClient) AccountUpdate(ctx context.Context, req *connect.Request[v1.AccountUpdateRequest]) (*connect.Response[v1.AccountUpdateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "AccountUpdate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.AccountUpdateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockAccountServiceClientMockRecorder) AccountUpdate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccountUpdate", reflect.TypeOf((*MockAccountService)(nil).AccountUpdate), ctx, req)
}

func (m *MockAccountService) LabelList(ctx context.Context, req *connect.Request[v1.LabelListRequest]) (*connect.Response[v1.LabelListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "LabelList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.LabelListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockAccountServiceMockRecorder) LabelList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LabelList", reflect.TypeOf((*MockAccountService)(nil).LabelList), ctx, req)
}

func (m *MockAccountServiceClient) LabelList(ctx context.Context, req *connect.Request[v1.LabelListRequest]) (*connect.Response[v1.LabelListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "LabelList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.LabelListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockAccountServiceClientMockRecorder) LabelList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LabelList", reflect.TypeOf((*MockAccountService)(nil).LabelList), ctx, req)
}

func (m *MockAccountService) AccountTypeList(ctx context.Context, req *connect.Request[v1.AccountTypeListRequest]) (*connect.Response[v1.AccountTypeListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "AccountTypeList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.AccountTypeListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockAccountServiceMockRecorder) AccountTypeList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccountTypeList", reflect.TypeOf((*MockAccountService)(nil).AccountTypeList), ctx, req)
}

func (m *MockAccountServiceClient) AccountTypeList(ctx context.Context, req *connect.Request[v1.AccountTypeListRequest]) (*connect.Response[v1.AccountTypeListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "AccountTypeList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.AccountTypeListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockAccountServiceClientMockRecorder) AccountTypeList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccountTypeList", reflect.TypeOf((*MockAccountService)(nil).AccountTypeList), ctx, req)
}

func (m *MockAccountService) AccountByIDs(ctx context.Context, req *connect.Request[v1.AccountByIDsRequest]) (*connect.Response[v1.AccountByIDsResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "AccountByIDs", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.AccountByIDsResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockAccountServiceMockRecorder) AccountByIDs(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccountByIDs", reflect.TypeOf((*MockAccountService)(nil).AccountByIDs), ctx, req)
}

func (m *MockAccountServiceClient) AccountByIDs(ctx context.Context, req *connect.Request[v1.AccountByIDsRequest]) (*connect.Response[v1.AccountByIDsResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "AccountByIDs", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.AccountByIDsResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockAccountServiceClientMockRecorder) AccountByIDs(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccountByIDs", reflect.TypeOf((*MockAccountService)(nil).AccountByIDs), ctx, req)
}

func (m *MockAccountService) AccountPublicSearch(ctx context.Context, req *connect.Request[v1.AccountPublicSearchRequest]) (*connect.Response[v1.AccountPublicSearchResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "AccountPublicSearch", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.AccountPublicSearchResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockAccountServiceMockRecorder) AccountPublicSearch(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccountPublicSearch", reflect.TypeOf((*MockAccountService)(nil).AccountPublicSearch), ctx, req)
}

func (m *MockAccountServiceClient) AccountPublicSearch(ctx context.Context, req *connect.Request[v1.AccountPublicSearchRequest]) (*connect.Response[v1.AccountPublicSearchResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "AccountPublicSearch", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.AccountPublicSearchResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockAccountServiceClientMockRecorder) AccountPublicSearch(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccountPublicSearch", reflect.TypeOf((*MockAccountService)(nil).AccountPublicSearch), ctx, req)
}

func (m *MockAccountService) AccountBalanceInit(ctx context.Context, req *connect.Request[v1.AccountBalanceInitRequest]) (*connect.Response[v1.AccountBalanceInitResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "AccountBalanceInit", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.AccountBalanceInitResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockAccountServiceMockRecorder) AccountBalanceInit(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccountBalanceInit", reflect.TypeOf((*MockAccountService)(nil).AccountBalanceInit), ctx, req)
}

func (m *MockAccountServiceClient) AccountBalanceInit(ctx context.Context, req *connect.Request[v1.AccountBalanceInitRequest]) (*connect.Response[v1.AccountBalanceInitResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "AccountBalanceInit", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.AccountBalanceInitResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockAccountServiceClientMockRecorder) AccountBalanceInit(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccountBalanceInit", reflect.TypeOf((*MockAccountService)(nil).AccountBalanceInit), ctx, req)
}

func (m *MockAccountService) TransferCreate(ctx context.Context, req *connect.Request[v1.TransferCreateRequest]) (*connect.Response[v1.TransferCreateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TransferCreate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.TransferCreateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockAccountServiceMockRecorder) TransferCreate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferCreate", reflect.TypeOf((*MockAccountService)(nil).TransferCreate), ctx, req)
}

func (m *MockAccountServiceClient) TransferCreate(ctx context.Context, req *connect.Request[v1.TransferCreateRequest]) (*connect.Response[v1.TransferCreateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TransferCreate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.TransferCreateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockAccountServiceClientMockRecorder) TransferCreate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferCreate", reflect.TypeOf((*MockAccountService)(nil).TransferCreate), ctx, req)
}

func (m *MockAccountService) TransferCancel(ctx context.Context, req *connect.Request[v1.TransferCancelRequest]) (*connect.Response[v1.TransferCancelResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TransferCancel", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.TransferCancelResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockAccountServiceMockRecorder) TransferCancel(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferCancel", reflect.TypeOf((*MockAccountService)(nil).TransferCancel), ctx, req)
}

func (m *MockAccountServiceClient) TransferCancel(ctx context.Context, req *connect.Request[v1.TransferCancelRequest]) (*connect.Response[v1.TransferCancelResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TransferCancel", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.TransferCancelResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockAccountServiceClientMockRecorder) TransferCancel(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferCancel", reflect.TypeOf((*MockAccountService)(nil).TransferCancel), ctx, req)
}

func (m *MockAccountService) AccountMutationList(ctx context.Context, req *connect.Request[v1.AccountMutationListRequest]) (*connect.Response[v1.AccountMutationListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "AccountMutationList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.AccountMutationListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockAccountServiceMockRecorder) AccountMutationList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccountMutationList", reflect.TypeOf((*MockAccountService)(nil).AccountMutationList), ctx, req)
}

func (m *MockAccountServiceClient) AccountMutationList(ctx context.Context, req *connect.Request[v1.AccountMutationListRequest]) (*connect.Response[v1.AccountMutationListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "AccountMutationList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.AccountMutationListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockAccountServiceClientMockRecorder) AccountMutationList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccountMutationList", reflect.TypeOf((*MockAccountService)(nil).AccountMutationList), ctx, req)
}

