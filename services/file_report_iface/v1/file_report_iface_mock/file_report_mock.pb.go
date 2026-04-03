package file_report_iface_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/file_report_iface/v1"
)

type MockFileReportService struct {
    ctrl     *gomock.Controller
    recorder *MockFileReportServiceMockRecorder
}

type MockFileReportServiceMockRecorder struct {
    mock *MockFileReportService
}

func NewMockFileReportService(ctrl *gomock.Controller) *MockFileReportService {
    mock := &MockFileReportService{ctrl: ctrl}
    mock.recorder = &MockFileReportServiceMockRecorder{mock}
    return mock
}

func (m *MockFileReportService) EXPECT() *MockFileReportServiceMockRecorder {
    return m.recorder
}

type MockFileReportServiceClient struct {
    ctrl     *gomock.Controller
    recorder *MockFileReportServiceClientMockRecorder
}

type MockFileReportServiceClientMockRecorder struct {
    mock *MockFileReportServiceClient
}

func NewMockFileReportServiceClient(ctrl *gomock.Controller) *MockFileReportServiceClient {
    mock := &MockFileReportServiceClient{ctrl: ctrl}
    mock.recorder = &MockFileReportServiceClientMockRecorder{mock}
    return mock
}

func (m *MockFileReportServiceClient) EXPECT() *MockFileReportServiceClientMockRecorder {
    return m.recorder
}

func (m *MockFileReportService) AvailableReport(ctx context.Context, req *connect.Request[v1.AvailableReportRequest]) (*connect.Response[v1.AvailableReportResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "AvailableReport", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.AvailableReportResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockFileReportServiceMockRecorder) AvailableReport(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AvailableReport", reflect.TypeOf((*MockFileReportService)(nil).AvailableReport), ctx, req)
}

func (m *MockFileReportServiceClient) AvailableReport(ctx context.Context, req *connect.Request[v1.AvailableReportRequest]) (*connect.Response[v1.AvailableReportResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "AvailableReport", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.AvailableReportResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockFileReportServiceClientMockRecorder) AvailableReport(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AvailableReport", reflect.TypeOf((*MockFileReportService)(nil).AvailableReport), ctx, req)
}

func (m *MockFileReportService) CreateReport(ctx context.Context, req *connect.Request[v1.CreateReportRequest]) (*connect.Response[v1.CreateReportResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "CreateReport", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.CreateReportResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockFileReportServiceMockRecorder) CreateReport(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateReport", reflect.TypeOf((*MockFileReportService)(nil).CreateReport), ctx, req)
}

func (m *MockFileReportServiceClient) CreateReport(ctx context.Context, req *connect.Request[v1.CreateReportRequest]) (*connect.Response[v1.CreateReportResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "CreateReport", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.CreateReportResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockFileReportServiceClientMockRecorder) CreateReport(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateReport", reflect.TypeOf((*MockFileReportService)(nil).CreateReport), ctx, req)
}

