package main

import (
	"context"
	"net/http"

	"connectrpc.com/connect"
	"github.com/pdcgo/schema/services/accounting_iface/v1"
	"github.com/pdcgo/schema/services/accounting_iface/v1/accounting_ifaceconnect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type transferService struct{}

// TransferCreate implements accounting_ifaceconnect.TransferServiceHandler.
func (t *transferService) TransferCreate(context.Context, *connect.Request[accounting_iface.TransferCreateRequest]) (*connect.Response[accounting_iface.TransferCreateResponse], error) {
	panic("unimplemented")
}

// TransferList implements accounting_ifaceconnect.TransferServiceHandler.
func (t *transferService) TransferList(context.Context, *connect.Request[accounting_iface.TransferListRequest]) (*connect.Response[accounting_iface.TransferListResponse], error) {
	panic("unimplemented")
}

// TransferTypeList implements accounting_ifaceconnect.TransferServiceHandler.
func (t *transferService) TransferTypeList(context.Context, *connect.Request[accounting_iface.TransferTypeListRequest]) (*connect.Response[accounting_iface.TransferTypeListResponse], error) {
	panic("unimplemented")
}

// TransferUpdate implements accounting_ifaceconnect.TransferServiceHandler.
func (t *transferService) TransferUpdate(context.Context, *connect.Request[accounting_iface.TransferUpdateRequest]) (*connect.Response[accounting_iface.TransferUpdateResponse], error) {
	panic("unimplemented")
}

func main() {
	mux := http.NewServeMux()
	path, handler := accounting_ifaceconnect.NewTransferServiceHandler(&transferService{})
	mux.Handle(path, handler)
	http.ListenAndServe(
		"localhost:8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)

}
