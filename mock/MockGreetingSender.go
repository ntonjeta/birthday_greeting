package mock

import (
	"github.com/stretchr/testify/mock"
)

type MockGreetingSender struct {
	mock.Mock
}

func (m *MockGreetingSender) Send(name string) error {
	args := m.Called(name)
	return args.Error(0)
}
