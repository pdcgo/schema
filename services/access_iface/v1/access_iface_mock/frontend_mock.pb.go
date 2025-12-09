package access_iface_mock

import (
    "context"
    "github.com/stretchr/testify/mock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/access_iface/v1"
)

type MockFrontendAccessService struct {
    mock.Mock
}

func (m *MockFrontendAccessService) SetupAccess(ctx context.Context, req *connect.Request[v1.SetupAccessRequest]) (*connect.Response[v1.SetupAccessResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.SetupAccessResponse]), args.Error(1)
}

func (m *MockFrontendAccessService) MenuAccess(ctx context.Context, req *connect.Request[v1.MenuAccessRequest]) (*connect.Response[v1.MenuAccessResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.MenuAccessResponse]), args.Error(1)
}

