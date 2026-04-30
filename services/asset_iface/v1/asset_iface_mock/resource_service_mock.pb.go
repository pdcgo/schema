package asset_iface_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/asset_iface/v1"
)

type MockResourceService struct {
    ctrl     *gomock.Controller
    recorder *MockResourceServiceMockRecorder
}

type MockResourceServiceMockRecorder struct {
    mock *MockResourceService
}

func NewMockResourceService(ctrl *gomock.Controller) *MockResourceService {
    mock := &MockResourceService{ctrl: ctrl}
    mock.recorder = &MockResourceServiceMockRecorder{mock}
    return mock
}

func (m *MockResourceService) EXPECT() *MockResourceServiceMockRecorder {
    return m.recorder
}

type MockResourceServiceClient struct {
    ctrl     *gomock.Controller
    recorder *MockResourceServiceClientMockRecorder
}

type MockResourceServiceClientMockRecorder struct {
    mock *MockResourceServiceClient
}

func NewMockResourceServiceClient(ctrl *gomock.Controller) *MockResourceServiceClient {
    mock := &MockResourceServiceClient{ctrl: ctrl}
    mock.recorder = &MockResourceServiceClientMockRecorder{mock}
    return mock
}

func (m *MockResourceServiceClient) EXPECT() *MockResourceServiceClientMockRecorder {
    return m.recorder
}

func (m *MockResourceService) UploadResource(ctx context.Context, req *connect.Request[v1.UploadResourceRequest]) (*connect.Response[v1.UploadResourceResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "UploadResource", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.UploadResourceResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockResourceServiceMockRecorder) UploadResource(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadResource", reflect.TypeOf((*MockResourceService)(nil).UploadResource), ctx, req)
}

func (m *MockResourceServiceClient) UploadResource(ctx context.Context, req *connect.Request[v1.UploadResourceRequest]) (*connect.Response[v1.UploadResourceResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "UploadResource", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.UploadResourceResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockResourceServiceClientMockRecorder) UploadResource(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadResource", reflect.TypeOf((*MockResourceService)(nil).UploadResource), ctx, req)
}

