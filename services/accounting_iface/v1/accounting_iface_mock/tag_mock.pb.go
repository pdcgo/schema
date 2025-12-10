package accounting_iface_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/accounting_iface/v1"
)

type MockTagService struct {
    ctrl     *gomock.Controller
    recorder *MockTagServiceMockRecorder
}

type MockTagServiceMockRecorder struct {
    mock *MockTagService
}

func NewMockTagService(ctrl *gomock.Controller) *MockTagService {
    mock := &MockTagService{ctrl: ctrl}
    mock.recorder = &MockTagServiceMockRecorder{mock}
    return mock
}

func (m *MockTagService) EXPECT() *MockTagServiceMockRecorder {
    return m.recorder
}

func (m *MockTagService) TagCreate(ctx context.Context, req *connect.Request[v1.TagCreateRequest]) (*connect.Response[v1.TagCreateResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TagCreate", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.TagCreateResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockTagServiceMockRecorder) TagCreate(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagCreate", reflect.TypeOf((*MockTagService)(nil).TagCreate), ctx, req)
}

func (m *MockTagService) TagList(ctx context.Context, req *connect.Request[v1.TagListRequest]) (*connect.Response[v1.TagListResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TagList", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.TagListResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockTagServiceMockRecorder) TagList(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagList", reflect.TypeOf((*MockTagService)(nil).TagList), ctx, req)
}

func (m *MockTagService) TagIDs(ctx context.Context, req *connect.Request[v1.TagIDsRequest]) (*connect.Response[v1.TagIDsResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "TagIDs", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.TagIDsResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockTagServiceMockRecorder) TagIDs(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagIDs", reflect.TypeOf((*MockTagService)(nil).TagIDs), ctx, req)
}

