package mocks

import (
	"github.com/stretchr/testify/mock"
)

type MockClient struct {
	mock.Mock
}

func (m *MockClient) CallAPI(method string, url string, params string, header map[string]string) (string, error) {
	args := m.Called(
		method,
		url,
		params,
		header,
	)
	return args.String(0), args.Error(1)
}
