package mock

import (
	"github.com/ntonjeta/greeting"
	"github.com/stretchr/testify/mock"
)

type MockFriendRepository struct {
	mock.Mock
}

func (m *MockFriendRepository) Get() ([]greeting.Friend, error) {

	m.Called()
	return nil, nil
}

type MockGreetingSender struct {
	mock.Mock
}

func (m *MockGreetingSender) Send(name string) error {
	m.Called(name)
	return nil
}
