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

func (m *MockWarehouseService) WarehouseDetail(ctx context.Context, req *connect.Request[v1.WarehouseDetailRequest]) (*connect.Response[v1.WarehouseDetailResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "WarehouseDetail", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.WarehouseDetailResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockWarehouseServiceMockRecorder) WarehouseDetail(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WarehouseDetail", reflect.TypeOf((*MockWarehouseService)(nil).WarehouseDetail), ctx, req)
}

func (m *MockWarehouseServiceClient) WarehouseDetail(ctx context.Context, req *connect.Request[v1.WarehouseDetailRequest]) (*connect.Response[v1.WarehouseDetailResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "WarehouseDetail", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.WarehouseDetailResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockWarehouseServiceClientMockRecorder) WarehouseDetail(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WarehouseDetail", reflect.TypeOf((*MockWarehouseService)(nil).WarehouseDetail), ctx, req)
}

func (m *MockWarehouseService) WarehouseCreate(ctx context.Context, req *connect.Request[v1.WarehouseCreateRequest], stream *connect.ServerStream[v1.WarehouseCreateResponse]) error {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "WarehouseCreate", ctx, req, stream)
    ret0, _ := ret[0].(error)
    return ret0
}

func (mr *MockWarehouseServiceMockRecorder) WarehouseCreate(ctx, req, stream interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WarehouseCreate", reflect.TypeOf((*MockWarehouseService)(nil).WarehouseCreate), ctx, req, stream)
}

func (m *MockWarehouseServiceClient) WarehouseCreate(ctx context.Context, req *connect.Request[v1.WarehouseCreateRequest]) (*connect.ServerStreamForClient[v1.WarehouseCreateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "WarehouseCreate", ctx, req)
    ret0, _ := ret[0].(*connect.ServerStreamForClient[v1.WarehouseCreateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockWarehouseServiceClientMockRecorder) WarehouseCreate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WarehouseCreate", reflect.TypeOf((*MockWarehouseService)(nil).WarehouseCreate), ctx, req)
}

func (m *MockWarehouseService) WarehouseUpdate(ctx context.Context, req *connect.Request[v1.WarehouseUpdateRequest]) (*connect.Response[v1.WarehouseUpdateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "WarehouseUpdate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.WarehouseUpdateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockWarehouseServiceMockRecorder) WarehouseUpdate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WarehouseUpdate", reflect.TypeOf((*MockWarehouseService)(nil).WarehouseUpdate), ctx, req)
}

func (m *MockWarehouseServiceClient) WarehouseUpdate(ctx context.Context, req *connect.Request[v1.WarehouseUpdateRequest]) (*connect.Response[v1.WarehouseUpdateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "WarehouseUpdate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.WarehouseUpdateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockWarehouseServiceClientMockRecorder) WarehouseUpdate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WarehouseUpdate", reflect.TypeOf((*MockWarehouseService)(nil).WarehouseUpdate), ctx, req)
}

func (m *MockWarehouseService) WarehouseDelete(ctx context.Context, req *connect.Request[v1.WarehouseDeleteRequest]) (*connect.Response[v1.WarehouseDeleteResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "WarehouseDelete", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.WarehouseDeleteResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockWarehouseServiceMockRecorder) WarehouseDelete(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WarehouseDelete", reflect.TypeOf((*MockWarehouseService)(nil).WarehouseDelete), ctx, req)
}

func (m *MockWarehouseServiceClient) WarehouseDelete(ctx context.Context, req *connect.Request[v1.WarehouseDeleteRequest]) (*connect.Response[v1.WarehouseDeleteResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "WarehouseDelete", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.WarehouseDeleteResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockWarehouseServiceClientMockRecorder) WarehouseDelete(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WarehouseDelete", reflect.TypeOf((*MockWarehouseService)(nil).WarehouseDelete), ctx, req)
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

func (m *MockWarehouseService) TransactionNoteCreate(ctx context.Context, req *connect.Request[v1.TransactionNoteCreateRequest]) (*connect.Response[v1.TransactionNoteCreateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TransactionNoteCreate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.TransactionNoteCreateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockWarehouseServiceMockRecorder) TransactionNoteCreate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionNoteCreate", reflect.TypeOf((*MockWarehouseService)(nil).TransactionNoteCreate), ctx, req)
}

func (m *MockWarehouseServiceClient) TransactionNoteCreate(ctx context.Context, req *connect.Request[v1.TransactionNoteCreateRequest]) (*connect.Response[v1.TransactionNoteCreateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TransactionNoteCreate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.TransactionNoteCreateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockWarehouseServiceClientMockRecorder) TransactionNoteCreate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionNoteCreate", reflect.TypeOf((*MockWarehouseService)(nil).TransactionNoteCreate), ctx, req)
}

func (m *MockWarehouseService) TransactionNoteList(ctx context.Context, req *connect.Request[v1.TransactionNoteListRequest]) (*connect.Response[v1.TransactionNoteListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TransactionNoteList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.TransactionNoteListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockWarehouseServiceMockRecorder) TransactionNoteList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionNoteList", reflect.TypeOf((*MockWarehouseService)(nil).TransactionNoteList), ctx, req)
}

func (m *MockWarehouseServiceClient) TransactionNoteList(ctx context.Context, req *connect.Request[v1.TransactionNoteListRequest]) (*connect.Response[v1.TransactionNoteListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TransactionNoteList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.TransactionNoteListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockWarehouseServiceClientMockRecorder) TransactionNoteList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionNoteList", reflect.TypeOf((*MockWarehouseService)(nil).TransactionNoteList), ctx, req)
}

func (m *MockWarehouseService) SellingTeamList(ctx context.Context, req *connect.Request[v1.SellingTeamListRequest]) (*connect.Response[v1.SellingTeamListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "SellingTeamList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.SellingTeamListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockWarehouseServiceMockRecorder) SellingTeamList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SellingTeamList", reflect.TypeOf((*MockWarehouseService)(nil).SellingTeamList), ctx, req)
}

func (m *MockWarehouseServiceClient) SellingTeamList(ctx context.Context, req *connect.Request[v1.SellingTeamListRequest]) (*connect.Response[v1.SellingTeamListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "SellingTeamList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.SellingTeamListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockWarehouseServiceClientMockRecorder) SellingTeamList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SellingTeamList", reflect.TypeOf((*MockWarehouseService)(nil).SellingTeamList), ctx, req)
}

func (m *MockWarehouseService) TeamRackList(ctx context.Context, req *connect.Request[v1.TeamRackListRequest]) (*connect.Response[v1.TeamRackListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TeamRackList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.TeamRackListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockWarehouseServiceMockRecorder) TeamRackList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeamRackList", reflect.TypeOf((*MockWarehouseService)(nil).TeamRackList), ctx, req)
}

func (m *MockWarehouseServiceClient) TeamRackList(ctx context.Context, req *connect.Request[v1.TeamRackListRequest]) (*connect.Response[v1.TeamRackListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TeamRackList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.TeamRackListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockWarehouseServiceClientMockRecorder) TeamRackList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeamRackList", reflect.TypeOf((*MockWarehouseService)(nil).TeamRackList), ctx, req)
}

func (m *MockWarehouseService) Stat(ctx context.Context, req *connect.Request[v1.StatRequest]) (*connect.Response[v1.StatResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "Stat", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.StatResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockWarehouseServiceMockRecorder) Stat(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stat", reflect.TypeOf((*MockWarehouseService)(nil).Stat), ctx, req)
}

func (m *MockWarehouseServiceClient) Stat(ctx context.Context, req *connect.Request[v1.StatRequest]) (*connect.Response[v1.StatResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "Stat", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.StatResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockWarehouseServiceClientMockRecorder) Stat(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stat", reflect.TypeOf((*MockWarehouseService)(nil).Stat), ctx, req)
}

