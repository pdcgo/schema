package common_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/common/v1"
)

type MockSupplierService struct {
    ctrl     *gomock.Controller
    recorder *MockSupplierServiceMockRecorder
}

type MockSupplierServiceMockRecorder struct {
    mock *MockSupplierService
}

func NewMockSupplierService(ctrl *gomock.Controller) *MockSupplierService {
    mock := &MockSupplierService{ctrl: ctrl}
    mock.recorder = &MockSupplierServiceMockRecorder{mock}
    return mock
}

func (m *MockSupplierService) EXPECT() *MockSupplierServiceMockRecorder {
    return m.recorder
}

func (m *MockSupplierService) PublicSupplierIDs(ctx context.Context, req *connect.Request[v1.PublicSupplierIDsRequest]) (*connect.Response[v1.PublicSupplierIDsResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "PublicSupplierIDs", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.PublicSupplierIDsResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockSupplierServiceMockRecorder) PublicSupplierIDs(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublicSupplierIDs", reflect.TypeOf((*MockSupplierService)(nil).PublicSupplierIDs), ctx, req)
}

