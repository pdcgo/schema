package access_iface_mock_test

import (
	"testing"

	connect "connectrpc.com/connect"
	gomock "github.com/golang/mock/gomock"
	"github.com/pdcgo/schema/services/access_iface/v1"
	"github.com/pdcgo/schema/services/access_iface/v1/access_iface_mock"
	"github.com/pdcgo/schema/services/access_iface/v1/access_ifaceconnect"
	"github.com/zeebo/assert"
)

func TestMock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockHelloService := access_iface_mock.NewMockHelloService(ctrl)
	mockHelloService.
		EXPECT().
		Hello(gomock.Any(), gomock.Any()).
		Return(nil, nil).
		AnyTimes()

	_, err := mockHelloService.Hello(t.Context(), &connect.Request[access_iface.HelloRequest]{
		Msg: &access_iface.HelloRequest{
			Name: "asdasd",
		},
	})

	assert.Nil(t, err)

	_, err = mockHelloService.Hello(t.Context(), &connect.Request[access_iface.HelloRequest]{
		Msg: &access_iface.HelloRequest{
			Name: "asdasd",
		},
	})

	assert.Nil(t, err)

	access_ifaceconnect.NewHelloServiceHandler(mockHelloService)
}
