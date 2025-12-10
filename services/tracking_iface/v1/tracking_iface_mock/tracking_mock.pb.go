package tracking_iface_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/tracking_iface/v1"
)

type MockTrackingService struct {
    ctrl     *gomock.Controller
    recorder *MockTrackingServiceMockRecorder
}

type MockTrackingServiceMockRecorder struct {
    mock *MockTrackingService
}

func NewMockTrackingService(ctrl *gomock.Controller) *MockTrackingService {
    mock := &MockTrackingService{ctrl: ctrl}
    mock.recorder = &MockTrackingServiceMockRecorder{mock}
    return mock
}

func (m *MockTrackingService) EXPECT() *MockTrackingServiceMockRecorder {
    return m.recorder
}

func (m *MockTrackingService) TrackingGet(ctx context.Context, req *connect.Request[v1.TrackingGetRequest]) (*connect.Response[v1.TrackingGetResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TrackingGet", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.TrackingGetResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockTrackingServiceMockRecorder) TrackingGet(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TrackingGet", reflect.TypeOf((*MockTrackingService)(nil).TrackingGet), ctx, req)
}

