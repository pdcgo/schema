package accounting_iface_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/accounting_iface/v1"
)

type MockLedgerService struct {
    ctrl     *gomock.Controller
    recorder *MockLedgerServiceMockRecorder
}

type MockLedgerServiceMockRecorder struct {
    mock *MockLedgerService
}

func NewMockLedgerService(ctrl *gomock.Controller) *MockLedgerService {
    mock := &MockLedgerService{ctrl: ctrl}
    mock.recorder = &MockLedgerServiceMockRecorder{mock}
    return mock
}

func (m *MockLedgerService) EXPECT() *MockLedgerServiceMockRecorder {
    return m.recorder
}

type MockLedgerServiceClient struct {
    ctrl     *gomock.Controller
    recorder *MockLedgerServiceClientMockRecorder
}

type MockLedgerServiceClientMockRecorder struct {
    mock *MockLedgerServiceClient
}

func NewMockLedgerServiceClient(ctrl *gomock.Controller) *MockLedgerServiceClient {
    mock := &MockLedgerServiceClient{ctrl: ctrl}
    mock.recorder = &MockLedgerServiceClientMockRecorder{mock}
    return mock
}

func (m *MockLedgerServiceClient) EXPECT() *MockLedgerServiceClientMockRecorder {
    return m.recorder
}

func (m *MockLedgerService) EntryList(ctx context.Context, req *connect.Request[v1.EntryListRequest]) (*connect.Response[v1.EntryListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "EntryList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.EntryListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockLedgerServiceMockRecorder) EntryList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EntryList", reflect.TypeOf((*MockLedgerService)(nil).EntryList), ctx, req)
}

func (m *MockLedgerServiceClient) EntryList(ctx context.Context, req *connect.Request[v1.EntryListRequest]) (*connect.Response[v1.EntryListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "EntryList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.EntryListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockLedgerServiceClientMockRecorder) EntryList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EntryList", reflect.TypeOf((*MockLedgerService)(nil).EntryList), ctx, req)
}

func (m *MockLedgerService) EntryListExtra(ctx context.Context, req *connect.Request[v1.EntryListExtraRequest]) (*connect.Response[v1.EntryListExtraResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "EntryListExtra", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.EntryListExtraResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockLedgerServiceMockRecorder) EntryListExtra(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EntryListExtra", reflect.TypeOf((*MockLedgerService)(nil).EntryListExtra), ctx, req)
}

func (m *MockLedgerServiceClient) EntryListExtra(ctx context.Context, req *connect.Request[v1.EntryListExtraRequest]) (*connect.Response[v1.EntryListExtraResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "EntryListExtra", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.EntryListExtraResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockLedgerServiceClientMockRecorder) EntryListExtra(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EntryListExtra", reflect.TypeOf((*MockLedgerService)(nil).EntryListExtra), ctx, req)
}

func (m *MockLedgerService) EntryListExport(ctx context.Context, req *connect.Request[v1.EntryListExportRequest], stream *connect.ServerStream[v1.EntryListExportResponse]) error {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "EntryListExport", ctx, req, stream)
    ret0, _ := ret[0].(error)
    return ret0
}

func (mr *MockLedgerServiceMockRecorder) EntryListExport(ctx, req, stream interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EntryListExport", reflect.TypeOf((*MockLedgerService)(nil).EntryListExport), ctx, req, stream)
}

func (m *MockLedgerServiceClient) EntryListExport(ctx context.Context, req *connect.Request[v1.EntryListExportRequest]) (*connect.ServerStreamForClient[v1.EntryListExportResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "EntryListExport", ctx, req)
    ret0, _ := ret[0].(*connect.ServerStreamForClient[v1.EntryListExportResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockLedgerServiceClientMockRecorder) EntryListExport(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EntryListExport", reflect.TypeOf((*MockLedgerService)(nil).EntryListExport), ctx, req)
}

func (m *MockLedgerService) TransactionDetail(ctx context.Context, req *connect.Request[v1.TransactionDetailRequest]) (*connect.Response[v1.TransactionDetailResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TransactionDetail", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.TransactionDetailResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockLedgerServiceMockRecorder) TransactionDetail(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionDetail", reflect.TypeOf((*MockLedgerService)(nil).TransactionDetail), ctx, req)
}

func (m *MockLedgerServiceClient) TransactionDetail(ctx context.Context, req *connect.Request[v1.TransactionDetailRequest]) (*connect.Response[v1.TransactionDetailResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TransactionDetail", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.TransactionDetailResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockLedgerServiceClientMockRecorder) TransactionDetail(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionDetail", reflect.TypeOf((*MockLedgerService)(nil).TransactionDetail), ctx, req)
}

