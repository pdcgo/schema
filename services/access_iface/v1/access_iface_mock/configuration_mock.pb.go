package access_iface_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/access_iface/v1"
)

type MockConfigurationService struct {
    ctrl     *gomock.Controller
    recorder *MockConfigurationServiceMockRecorder
}

type MockConfigurationServiceMockRecorder struct {
    mock *MockConfigurationService
}

func NewMockConfigurationService(ctrl *gomock.Controller) *MockConfigurationService {
    mock := &MockConfigurationService{ctrl: ctrl}
    mock.recorder = &MockConfigurationServiceMockRecorder{mock}
    return mock
}

func (m *MockConfigurationService) EXPECT() *MockConfigurationServiceMockRecorder {
    return m.recorder
}

type MockConfigurationServiceClient struct {
    ctrl     *gomock.Controller
    recorder *MockConfigurationServiceClientMockRecorder
}

type MockConfigurationServiceClientMockRecorder struct {
    mock *MockConfigurationServiceClient
}

func NewMockConfigurationServiceClient(ctrl *gomock.Controller) *MockConfigurationServiceClient {
    mock := &MockConfigurationServiceClient{ctrl: ctrl}
    mock.recorder = &MockConfigurationServiceClientMockRecorder{mock}
    return mock
}

func (m *MockConfigurationServiceClient) EXPECT() *MockConfigurationServiceClientMockRecorder {
    return m.recorder
}

func (m *MockConfigurationService) ExtensionConfiguration(ctx context.Context, req *connect.Request[v1.ExtensionConfigurationRequest]) (*connect.Response[v1.ExtensionConfigurationResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ExtensionConfiguration", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ExtensionConfigurationResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockConfigurationServiceMockRecorder) ExtensionConfiguration(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtensionConfiguration", reflect.TypeOf((*MockConfigurationService)(nil).ExtensionConfiguration), ctx, req)
}

func (m *MockConfigurationServiceClient) ExtensionConfiguration(ctx context.Context, req *connect.Request[v1.ExtensionConfigurationRequest]) (*connect.Response[v1.ExtensionConfigurationResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ExtensionConfiguration", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ExtensionConfigurationResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockConfigurationServiceClientMockRecorder) ExtensionConfiguration(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtensionConfiguration", reflect.TypeOf((*MockConfigurationService)(nil).ExtensionConfiguration), ctx, req)
}

func (m *MockConfigurationService) ExtensionConfigurationReplace(ctx context.Context, req *connect.Request[v1.ExtensionConfigurationReplaceRequest]) (*connect.Response[v1.ExtensionConfigurationReplaceResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ExtensionConfigurationReplace", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ExtensionConfigurationReplaceResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockConfigurationServiceMockRecorder) ExtensionConfigurationReplace(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtensionConfigurationReplace", reflect.TypeOf((*MockConfigurationService)(nil).ExtensionConfigurationReplace), ctx, req)
}

func (m *MockConfigurationServiceClient) ExtensionConfigurationReplace(ctx context.Context, req *connect.Request[v1.ExtensionConfigurationReplaceRequest]) (*connect.Response[v1.ExtensionConfigurationReplaceResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ExtensionConfigurationReplace", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.ExtensionConfigurationReplaceResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockConfigurationServiceClientMockRecorder) ExtensionConfigurationReplace(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtensionConfigurationReplace", reflect.TypeOf((*MockConfigurationService)(nil).ExtensionConfigurationReplace), ctx, req)
}

func (m *MockConfigurationService) AndroidCheckLatestVersion(ctx context.Context, req *connect.Request[v1.AndroidCheckLatestVersionRequest]) (*connect.Response[v1.AndroidCheckLatestVersionResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "AndroidCheckLatestVersion", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.AndroidCheckLatestVersionResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockConfigurationServiceMockRecorder) AndroidCheckLatestVersion(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AndroidCheckLatestVersion", reflect.TypeOf((*MockConfigurationService)(nil).AndroidCheckLatestVersion), ctx, req)
}

func (m *MockConfigurationServiceClient) AndroidCheckLatestVersion(ctx context.Context, req *connect.Request[v1.AndroidCheckLatestVersionRequest]) (*connect.Response[v1.AndroidCheckLatestVersionResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "AndroidCheckLatestVersion", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.AndroidCheckLatestVersionResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockConfigurationServiceClientMockRecorder) AndroidCheckLatestVersion(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AndroidCheckLatestVersion", reflect.TypeOf((*MockConfigurationService)(nil).AndroidCheckLatestVersion), ctx, req)
}

func (m *MockConfigurationService) AndroidReleases(ctx context.Context, req *connect.Request[v1.AndroidReleasesRequest]) (*connect.Response[v1.AndroidReleasesResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "AndroidReleases", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.AndroidReleasesResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockConfigurationServiceMockRecorder) AndroidReleases(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AndroidReleases", reflect.TypeOf((*MockConfigurationService)(nil).AndroidReleases), ctx, req)
}

func (m *MockConfigurationServiceClient) AndroidReleases(ctx context.Context, req *connect.Request[v1.AndroidReleasesRequest]) (*connect.Response[v1.AndroidReleasesResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "AndroidReleases", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.AndroidReleasesResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockConfigurationServiceClientMockRecorder) AndroidReleases(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AndroidReleases", reflect.TypeOf((*MockConfigurationService)(nil).AndroidReleases), ctx, req)
}

func (m *MockConfigurationService) AndroidReleaseGet(ctx context.Context, req *connect.Request[v1.AndroidReleaseGetRequest]) (*connect.Response[v1.AndroidReleaseGetResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "AndroidReleaseGet", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.AndroidReleaseGetResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockConfigurationServiceMockRecorder) AndroidReleaseGet(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AndroidReleaseGet", reflect.TypeOf((*MockConfigurationService)(nil).AndroidReleaseGet), ctx, req)
}

func (m *MockConfigurationServiceClient) AndroidReleaseGet(ctx context.Context, req *connect.Request[v1.AndroidReleaseGetRequest]) (*connect.Response[v1.AndroidReleaseGetResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "AndroidReleaseGet", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.AndroidReleaseGetResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockConfigurationServiceClientMockRecorder) AndroidReleaseGet(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AndroidReleaseGet", reflect.TypeOf((*MockConfigurationService)(nil).AndroidReleaseGet), ctx, req)
}

