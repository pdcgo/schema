package tracking_iface_mock

import (
    "context"
    "github.com/stretchr/testify/mock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/tracking_iface/v1"
)

type MockTrackingService struct {
    mock.Mock
}

func (m *MockTrackingService) TrackingGet(ctx context.Context, req *connect.Request[v1.TrackingGetRequest]) (*connect.Response[v1.TrackingGetResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.TrackingGetResponse]), args.Error(1)
}

