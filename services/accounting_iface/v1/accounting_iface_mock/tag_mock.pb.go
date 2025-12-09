package accounting_iface_mock

import (
    "context"
    "github.com/stretchr/testify/mock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/accounting_iface/v1"
)

type MockTagService struct {
    mock.Mock
}

func (m *MockTagService) TagCreate(ctx context.Context, req *connect.Request[v1.TagCreateRequest]) (*connect.Response[v1.TagCreateResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.TagCreateResponse]), args.Error(1)
}

func (m *MockTagService) TagList(ctx context.Context, req *connect.Request[v1.TagListRequest]) (*connect.Response[v1.TagListResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.TagListResponse]), args.Error(1)
}

func (m *MockTagService) TagIDs(ctx context.Context, req *connect.Request[v1.TagIDsRequest]) (*connect.Response[v1.TagIDsResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.TagIDsResponse]), args.Error(1)
}

