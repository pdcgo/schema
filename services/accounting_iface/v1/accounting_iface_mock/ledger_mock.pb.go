package accounting_iface_mock

import (
    "context"
    "github.com/stretchr/testify/mock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/accounting_iface/v1"
)

type MockLedgerService struct {
    mock.Mock
}

func (m *MockLedgerService) EntryList(ctx context.Context, req *connect.Request[v1.EntryListRequest]) (*connect.Response[v1.EntryListResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.EntryListResponse]), args.Error(1)
}

func (m *MockLedgerService) EntryListExtra(ctx context.Context, req *connect.Request[v1.EntryListExtraRequest]) (*connect.Response[v1.EntryListExtraResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.EntryListExtraResponse]), args.Error(1)
}

func (m *MockLedgerService) EntryListExport(ctx context.Context, req *connect.Request[v1.EntryListExportRequest]) (*connect.Response[v1.EntryListExportResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.EntryListExportResponse]), args.Error(1)
}

func (m *MockLedgerService) TransactionDetail(ctx context.Context, req *connect.Request[v1.TransactionDetailRequest]) (*connect.Response[v1.TransactionDetailResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.TransactionDetailResponse]), args.Error(1)
}

