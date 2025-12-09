package accounting_iface_mock

import (
    "context"
    "github.com/stretchr/testify/mock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/accounting_iface/v1"
)

type MockAccountService struct {
    mock.Mock
}

func (m *MockAccountService) AccountCreate(ctx context.Context, req *connect.Request[v1.AccountCreateRequest]) (*connect.Response[v1.AccountCreateResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.AccountCreateResponse]), args.Error(1)
}

func (m *MockAccountService) AccountList(ctx context.Context, req *connect.Request[v1.AccountListRequest]) (*connect.Response[v1.AccountListResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.AccountListResponse]), args.Error(1)
}

func (m *MockAccountService) AccountDelete(ctx context.Context, req *connect.Request[v1.AccountDeleteRequest]) (*connect.Response[v1.AccountDeleteResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.AccountDeleteResponse]), args.Error(1)
}

func (m *MockAccountService) AccountUpdate(ctx context.Context, req *connect.Request[v1.AccountUpdateRequest]) (*connect.Response[v1.AccountUpdateResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.AccountUpdateResponse]), args.Error(1)
}

func (m *MockAccountService) LabelList(ctx context.Context, req *connect.Request[v1.LabelListRequest]) (*connect.Response[v1.LabelListResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.LabelListResponse]), args.Error(1)
}

func (m *MockAccountService) AccountTypeList(ctx context.Context, req *connect.Request[v1.AccountTypeListRequest]) (*connect.Response[v1.AccountTypeListResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.AccountTypeListResponse]), args.Error(1)
}

func (m *MockAccountService) AccountByIDs(ctx context.Context, req *connect.Request[v1.AccountByIDsRequest]) (*connect.Response[v1.AccountByIDsResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.AccountByIDsResponse]), args.Error(1)
}

func (m *MockAccountService) AccountPublicSearch(ctx context.Context, req *connect.Request[v1.AccountPublicSearchRequest]) (*connect.Response[v1.AccountPublicSearchResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.AccountPublicSearchResponse]), args.Error(1)
}

func (m *MockAccountService) AccountBalanceInit(ctx context.Context, req *connect.Request[v1.AccountBalanceInitRequest]) (*connect.Response[v1.AccountBalanceInitResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.AccountBalanceInitResponse]), args.Error(1)
}

func (m *MockAccountService) TransferCreate(ctx context.Context, req *connect.Request[v1.TransferCreateRequest]) (*connect.Response[v1.TransferCreateResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.TransferCreateResponse]), args.Error(1)
}

func (m *MockAccountService) TransferCancel(ctx context.Context, req *connect.Request[v1.TransferCancelRequest]) (*connect.Response[v1.TransferCancelResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.TransferCancelResponse]), args.Error(1)
}

func (m *MockAccountService) AccountMutationList(ctx context.Context, req *connect.Request[v1.AccountMutationListRequest]) (*connect.Response[v1.AccountMutationListResponse], error) {
    args := m.Called(ctx, req)
    return args.Get(0).(*connect.Response[v1.AccountMutationListResponse]), args.Error(1)
}

