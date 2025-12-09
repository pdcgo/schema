package common_mock

import (
    "context"
    "github.com/stretchr/testify/mock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/common/v1"
)

type MockTeamService struct {
    mock.Mock
}

func (m *MockTeamService) PublicTeamIDs(ctx context.Context, req *connect.Request[v1.PublicTeamIDsRequest]) (*connect.Response[v1.PublicTeamIDsResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.PublicTeamIDsResponse]), args.Error(1)
}

func (m *MockTeamService) PublicTeamList(ctx context.Context, req *connect.Request[v1.PublicTeamListRequest]) (*connect.Response[v1.PublicTeamListResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.PublicTeamListResponse]), args.Error(1)
}

