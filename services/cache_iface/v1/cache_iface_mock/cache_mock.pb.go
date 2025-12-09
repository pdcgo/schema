package cache_iface_mock

import (
    "context"
    "github.com/stretchr/testify/mock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/cache_iface/v1"
)

type MockCacheService struct {
    mock.Mock
}

func (m *MockCacheService) Add(ctx context.Context, req *connect.Request[v1.AddRequest]) (*connect.Response[v1.AddResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.AddResponse]), args.Error(1)
}

func (m *MockCacheService) Replace(ctx context.Context, req *connect.Request[v1.ReplaceRequest]) (*connect.Response[v1.ReplaceResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.ReplaceResponse]), args.Error(1)
}

func (m *MockCacheService) Get(ctx context.Context, req *connect.Request[v1.GetRequest]) (*connect.Response[v1.GetResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.GetResponse]), args.Error(1)
}

func (m *MockCacheService) Delete(ctx context.Context, req *connect.Request[v1.DeleteRequest]) (*connect.Response[v1.DeleteResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.DeleteResponse]), args.Error(1)
}

func (m *MockCacheService) Flush(ctx context.Context, req *connect.Request[v1.FlushRequest]) (*connect.Response[v1.FlushResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.FlushResponse]), args.Error(1)
}

