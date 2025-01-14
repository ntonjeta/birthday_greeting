package mock

import (
	"github.com/ntonjeta/greeting"

	"github.com/stretchr/testify/mock"
)

type MockFriendRepository struct {
	mock.Mock
}

func (m *MockFriendRepository) Get() ([]greeting.Friend, error) {

	args := m.Called()

	return args.Get(0).([]greeting.Friend), args.Error(1)
}

type MockGreetingSender struct {
	mock.Mock
}

func (m *MockGreetingSender) Send(name string) error {
	args := m.Called(name)
	return args.Error(0)
}
