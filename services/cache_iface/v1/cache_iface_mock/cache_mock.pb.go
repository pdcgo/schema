package cache_iface_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/cache_iface/v1"
)

type MockCacheService struct {
    ctrl     *gomock.Controller
    recorder *MockCacheServiceMockRecorder
}

type MockCacheServiceMockRecorder struct {
    mock *MockCacheService
}

func NewMockCacheService(ctrl *gomock.Controller) *MockCacheService {
    mock := &MockCacheService{ctrl: ctrl}
    mock.recorder = &MockCacheServiceMockRecorder{mock}
    return mock
}

func (m *MockCacheService) EXPECT() *MockCacheServiceMockRecorder {
    return m.recorder
}

func (m *MockCacheService) Add(ctx context.Context, req *connect.Request[v1.AddRequest]) (*connect.Response[v1.AddResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "Add", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.AddResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockCacheServiceMockRecorder) Add(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockCacheService)(nil).Add), ctx, req)
}

func (m *MockCacheService) Replace(ctx context.Context, req *connect.Request[v1.ReplaceRequest]) (*connect.Response[v1.ReplaceResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "Replace", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ReplaceResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockCacheServiceMockRecorder) Replace(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Replace", reflect.TypeOf((*MockCacheService)(nil).Replace), ctx, req)
}

func (m *MockCacheService) Get(ctx context.Context, req *connect.Request[v1.GetRequest]) (*connect.Response[v1.GetResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "Get", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.GetResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockCacheServiceMockRecorder) Get(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCacheService)(nil).Get), ctx, req)
}

func (m *MockCacheService) Delete(ctx context.Context, req *connect.Request[v1.DeleteRequest]) (*connect.Response[v1.DeleteResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "Delete", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.DeleteResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockCacheServiceMockRecorder) Delete(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCacheService)(nil).Delete), ctx, req)
}

func (m *MockCacheService) Flush(ctx context.Context, req *connect.Request[v1.FlushRequest]) (*connect.Response[v1.FlushResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "Flush", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.FlushResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockCacheServiceMockRecorder) Flush(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Flush", reflect.TypeOf((*MockCacheService)(nil).Flush), ctx, req)
}

