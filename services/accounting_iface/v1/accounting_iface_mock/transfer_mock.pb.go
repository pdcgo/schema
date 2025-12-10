package accounting_iface_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/accounting_iface/v1"
)

type MockTransferService struct {
    ctrl     *gomock.Controller
    recorder *MockTransferServiceMockRecorder
}

type MockTransferServiceMockRecorder struct {
    mock *MockTransferService
}

func NewMockTransferService(ctrl *gomock.Controller) *MockTransferService {
    mock := &MockTransferService{ctrl: ctrl}
    mock.recorder = &MockTransferServiceMockRecorder{mock}
    return mock
}

func (m *MockTransferService) EXPECT() *MockTransferServiceMockRecorder {
    return m.recorder
}

func (m *MockTransferService) TransferTeam(ctx context.Context, req *connect.Request[v1.TransferTeamRequest]) (*connect.Response[v1.TransferTeamResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TransferTeam", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.TransferTeamResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockTransferServiceMockRecorder) TransferTeam(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferTeam", reflect.TypeOf((*MockTransferService)(nil).TransferTeam), ctx, req)
}

func (m *MockTransferService) TransferList(ctx context.Context, req *connect.Request[v1.TransferListRequest]) (*connect.Response[v1.TransferListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TransferList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.TransferListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockTransferServiceMockRecorder) TransferList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferList", reflect.TypeOf((*MockTransferService)(nil).TransferList), ctx, req)
}

