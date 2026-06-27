package document_iface_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/document_iface/v1"
)

type MockDocumentService struct {
    ctrl     *gomock.Controller
    recorder *MockDocumentServiceMockRecorder
}

type MockDocumentServiceMockRecorder struct {
    mock *MockDocumentService
}

func NewMockDocumentService(ctrl *gomock.Controller) *MockDocumentService {
    mock := &MockDocumentService{ctrl: ctrl}
    mock.recorder = &MockDocumentServiceMockRecorder{mock}
    return mock
}

func (m *MockDocumentService) EXPECT() *MockDocumentServiceMockRecorder {
    return m.recorder
}

type MockDocumentServiceClient struct {
    ctrl     *gomock.Controller
    recorder *MockDocumentServiceClientMockRecorder
}

type MockDocumentServiceClientMockRecorder struct {
    mock *MockDocumentServiceClient
}

func NewMockDocumentServiceClient(ctrl *gomock.Controller) *MockDocumentServiceClient {
    mock := &MockDocumentServiceClient{ctrl: ctrl}
    mock.recorder = &MockDocumentServiceClientMockRecorder{mock}
    return mock
}

func (m *MockDocumentServiceClient) EXPECT() *MockDocumentServiceClientMockRecorder {
    return m.recorder
}

func (m *MockDocumentService) RequestUpload(ctx context.Context, req *connect.Request[v1.RequestUploadRequest]) (*connect.Response[v1.RequestUploadResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "RequestUpload", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.RequestUploadResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockDocumentServiceMockRecorder) RequestUpload(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestUpload", reflect.TypeOf((*MockDocumentService)(nil).RequestUpload), ctx, req)
}

func (m *MockDocumentServiceClient) RequestUpload(ctx context.Context, req *connect.Request[v1.RequestUploadRequest]) (*connect.Response[v1.RequestUploadResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "RequestUpload", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.RequestUploadResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockDocumentServiceClientMockRecorder) RequestUpload(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestUpload", reflect.TypeOf((*MockDocumentService)(nil).RequestUpload), ctx, req)
}

func (m *MockDocumentService) ConfirmUpload(ctx context.Context, req *connect.Request[v1.ConfirmUploadRequest]) (*connect.Response[v1.ConfirmUploadResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ConfirmUpload", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ConfirmUploadResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockDocumentServiceMockRecorder) ConfirmUpload(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfirmUpload", reflect.TypeOf((*MockDocumentService)(nil).ConfirmUpload), ctx, req)
}

func (m *MockDocumentServiceClient) ConfirmUpload(ctx context.Context, req *connect.Request[v1.ConfirmUploadRequest]) (*connect.Response[v1.ConfirmUploadResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ConfirmUpload", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ConfirmUploadResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockDocumentServiceClientMockRecorder) ConfirmUpload(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfirmUpload", reflect.TypeOf((*MockDocumentService)(nil).ConfirmUpload), ctx, req)
}

func (m *MockDocumentService) GetDownloadUrl(ctx context.Context, req *connect.Request[v1.GetDownloadUrlRequest]) (*connect.Response[v1.GetDownloadUrlResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "GetDownloadUrl", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.GetDownloadUrlResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockDocumentServiceMockRecorder) GetDownloadUrl(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDownloadUrl", reflect.TypeOf((*MockDocumentService)(nil).GetDownloadUrl), ctx, req)
}

func (m *MockDocumentServiceClient) GetDownloadUrl(ctx context.Context, req *connect.Request[v1.GetDownloadUrlRequest]) (*connect.Response[v1.GetDownloadUrlResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "GetDownloadUrl", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.GetDownloadUrlResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockDocumentServiceClientMockRecorder) GetDownloadUrl(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDownloadUrl", reflect.TypeOf((*MockDocumentService)(nil).GetDownloadUrl), ctx, req)
}

