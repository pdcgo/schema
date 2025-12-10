package stock_iface_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/stock_iface/v1"
)

type MockStockReportService struct {
    ctrl     *gomock.Controller
    recorder *MockStockReportServiceMockRecorder
}

type MockStockReportServiceMockRecorder struct {
    mock *MockStockReportService
}

func NewMockStockReportService(ctrl *gomock.Controller) *MockStockReportService {
    mock := &MockStockReportService{ctrl: ctrl}
    mock.recorder = &MockStockReportServiceMockRecorder{mock}
    return mock
}

func (m *MockStockReportService) EXPECT() *MockStockReportServiceMockRecorder {
    return m.recorder
}

func (m *MockStockReportService) StockTeam(ctx context.Context, req *connect.Request[v1.StockTeamRequest]) (*connect.Response[v1.StockTeamResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "StockTeam", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.StockTeamResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockStockReportServiceMockRecorder) StockTeam(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StockTeam", reflect.TypeOf((*MockStockReportService)(nil).StockTeam), ctx, req)
}

