package legacy_iface_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/legacy_iface/v1"
)

type MockLegacyInvoiceService struct {
    ctrl     *gomock.Controller
    recorder *MockLegacyInvoiceServiceMockRecorder
}

type MockLegacyInvoiceServiceMockRecorder struct {
    mock *MockLegacyInvoiceService
}

func NewMockLegacyInvoiceService(ctrl *gomock.Controller) *MockLegacyInvoiceService {
    mock := &MockLegacyInvoiceService{ctrl: ctrl}
    mock.recorder = &MockLegacyInvoiceServiceMockRecorder{mock}
    return mock
}

func (m *MockLegacyInvoiceService) EXPECT() *MockLegacyInvoiceServiceMockRecorder {
    return m.recorder
}

type MockLegacyInvoiceServiceClient struct {
    ctrl     *gomock.Controller
    recorder *MockLegacyInvoiceServiceClientMockRecorder
}

type MockLegacyInvoiceServiceClientMockRecorder struct {
    mock *MockLegacyInvoiceServiceClient
}

func NewMockLegacyInvoiceServiceClient(ctrl *gomock.Controller) *MockLegacyInvoiceServiceClient {
    mock := &MockLegacyInvoiceServiceClient{ctrl: ctrl}
    mock.recorder = &MockLegacyInvoiceServiceClientMockRecorder{mock}
    return mock
}

func (m *MockLegacyInvoiceServiceClient) EXPECT() *MockLegacyInvoiceServiceClientMockRecorder {
    return m.recorder
}

func (m *MockLegacyInvoiceService) InvoiceOverview(ctx context.Context, req *connect.Request[v1.InvoiceOverviewRequest]) (*connect.Response[v1.InvoiceOverviewResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "InvoiceOverview", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.InvoiceOverviewResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockLegacyInvoiceServiceMockRecorder) InvoiceOverview(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvoiceOverview", reflect.TypeOf((*MockLegacyInvoiceService)(nil).InvoiceOverview), ctx, req)
}

func (m *MockLegacyInvoiceServiceClient) InvoiceOverview(ctx context.Context, req *connect.Request[v1.InvoiceOverviewRequest]) (*connect.Response[v1.InvoiceOverviewResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "InvoiceOverview", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.InvoiceOverviewResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockLegacyInvoiceServiceClientMockRecorder) InvoiceOverview(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvoiceOverview", reflect.TypeOf((*MockLegacyInvoiceService)(nil).InvoiceOverview), ctx, req)
}

