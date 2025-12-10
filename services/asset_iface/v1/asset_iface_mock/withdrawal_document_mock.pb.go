package asset_iface_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/asset_iface/v1"
)

type MockWithdrawalDocumentService struct {
    ctrl     *gomock.Controller
    recorder *MockWithdrawalDocumentServiceMockRecorder
}

type MockWithdrawalDocumentServiceMockRecorder struct {
    mock *MockWithdrawalDocumentService
}

func NewMockWithdrawalDocumentService(ctrl *gomock.Controller) *MockWithdrawalDocumentService {
    mock := &MockWithdrawalDocumentService{ctrl: ctrl}
    mock.recorder = &MockWithdrawalDocumentServiceMockRecorder{mock}
    return mock
}

func (m *MockWithdrawalDocumentService) EXPECT() *MockWithdrawalDocumentServiceMockRecorder {
    return m.recorder
}

func (m *MockWithdrawalDocumentService) Upload(ctx context.Context, req *connect.Request[v1.UploadRequest]) (*connect.Response[v1.UploadResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "Upload", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.UploadResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockWithdrawalDocumentServiceMockRecorder) Upload(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upload", reflect.TypeOf((*MockWithdrawalDocumentService)(nil).Upload), ctx, req)
}

