package order_iface_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/order_iface/v1"
)

type MockOrderReportService struct {
    ctrl     *gomock.Controller
    recorder *MockOrderReportServiceMockRecorder
}

type MockOrderReportServiceMockRecorder struct {
    mock *MockOrderReportService
}

func NewMockOrderReportService(ctrl *gomock.Controller) *MockOrderReportService {
    mock := &MockOrderReportService{ctrl: ctrl}
    mock.recorder = &MockOrderReportServiceMockRecorder{mock}
    return mock
}

func (m *MockOrderReportService) EXPECT() *MockOrderReportServiceMockRecorder {
    return m.recorder
}

type MockOrderReportServiceClient struct {
    ctrl     *gomock.Controller
    recorder *MockOrderReportServiceClientMockRecorder
}

type MockOrderReportServiceClientMockRecorder struct {
    mock *MockOrderReportServiceClient
}

func NewMockOrderReportServiceClient(ctrl *gomock.Controller) *MockOrderReportServiceClient {
    mock := &MockOrderReportServiceClient{ctrl: ctrl}
    mock.recorder = &MockOrderReportServiceClientMockRecorder{mock}
    return mock
}

func (m *MockOrderReportServiceClient) EXPECT() *MockOrderReportServiceClientMockRecorder {
    return m.recorder
}

func (m *MockOrderReportService) DailyFundList(ctx context.Context, req *connect.Request[v1.DailyFundListRequest]) (*connect.Response[v1.DailyFundListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "DailyFundList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.DailyFundListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockOrderReportServiceMockRecorder) DailyFundList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DailyFundList", reflect.TypeOf((*MockOrderReportService)(nil).DailyFundList), ctx, req)
}

func (m *MockOrderReportServiceClient) DailyFundList(ctx context.Context, req *connect.Request[v1.DailyFundListRequest]) (*connect.Response[v1.DailyFundListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "DailyFundList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.DailyFundListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockOrderReportServiceClientMockRecorder) DailyFundList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DailyFundList", reflect.TypeOf((*MockOrderReportService)(nil).DailyFundList), ctx, req)
}

