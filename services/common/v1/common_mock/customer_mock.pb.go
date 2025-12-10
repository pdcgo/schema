package common_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/common/v1"
)

type MockCustomerDataService struct {
    ctrl     *gomock.Controller
    recorder *MockCustomerDataServiceMockRecorder
}

type MockCustomerDataServiceMockRecorder struct {
    mock *MockCustomerDataService
}

func NewMockCustomerDataService(ctrl *gomock.Controller) *MockCustomerDataService {
    mock := &MockCustomerDataService{ctrl: ctrl}
    mock.recorder = &MockCustomerDataServiceMockRecorder{mock}
    return mock
}

func (m *MockCustomerDataService) EXPECT() *MockCustomerDataServiceMockRecorder {
    return m.recorder
}

func (m *MockCustomerDataService) CustomerIDs(ctx context.Context, req *connect.Request[v1.CustomerIDsRequest]) (*connect.Response[v1.CustomerIDsResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "CustomerIDs", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.CustomerIDsResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockCustomerDataServiceMockRecorder) CustomerIDs(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CustomerIDs", reflect.TypeOf((*MockCustomerDataService)(nil).CustomerIDs), ctx, req)
}

