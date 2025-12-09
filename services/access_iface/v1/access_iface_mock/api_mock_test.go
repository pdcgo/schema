package access_iface_mock_test

import (
	"errors"
	"testing"

	connect "connectrpc.com/connect"
	"github.com/pdcgo/schema/services/access_iface/v1"
	"github.com/pdcgo/schema/services/access_iface/v1/access_iface_mock"
	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	hello := new(access_iface_mock.MockHelloService)

	hello.On("Hello", &connect.Response[access_iface.HelloResponse]{}, nil).Return(nil, errors.New("asdasd"))

	_, err := hello.Hello(t.Context(), &connect.Request[access_iface.HelloRequest]{
		Msg: &access_iface.HelloRequest{},
	})

	hello.AssertExpectations(t)
	assert.Nil(t, err)
}
