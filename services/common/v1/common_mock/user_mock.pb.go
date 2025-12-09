package common_mock

import (
    "context"
    "github.com/stretchr/testify/mock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/common/v1"
)

type MockLoginService struct {
    mock.Mock
}

func (m *MockLoginService) Login(ctx context.Context, req *connect.Request[v1.LoginRequest]) (*connect.Response[v1.LoginResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.LoginResponse]), args.Error(1)
}

type MockUserService struct {
    mock.Mock
}

func (m *MockUserService) PublicUserIDs(ctx context.Context, req *connect.Request[v1.PublicUserIDsRequest]) (*connect.Response[v1.PublicUserIDsResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.PublicUserIDsResponse]), args.Error(1)
}

