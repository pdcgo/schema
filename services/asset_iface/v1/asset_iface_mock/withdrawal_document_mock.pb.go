package asset_iface_mock

import (
    "context"
    "github.com/stretchr/testify/mock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/asset_iface/v1"
)

type MockWithdrawalDocumentService struct {
    mock.Mock
}

func (m *MockWithdrawalDocumentService) Upload(ctx context.Context, req *connect.Request[v1.UploadRequest]) (*connect.Response[v1.UploadResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.UploadResponse]), args.Error(1)
}

