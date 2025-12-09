package access_iface_mock

import (
    "context"
    "github.com/stretchr/testify/mock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/access_iface/v1"
)

type MockHelloService struct {
    mock.Mock
}

func (m *MockHelloService) Hello(ctx context.Context, req *connect.Request[v1.HelloRequest]) (*connect.Response[v1.HelloResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.HelloResponse]), args.Error(1)
}

