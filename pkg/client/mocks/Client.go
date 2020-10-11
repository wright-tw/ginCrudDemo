package mocks

import (
	"github.com/stretchr/testify/mock"
)

type Client struct {
	mock.Mock
}

func (m *Client) CallAPI(method string, url string, params string, header map[string]string) (string, error) {
	args := m.Called(
		method,
		url,
		params,
		header,
	)
	return args.String(0), args.Error(1)
}
