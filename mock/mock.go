package mock

import (
	"time"

	"github.com/ntonjeta/greeting"

	"github.com/stretchr/testify/mock"
)

var friends = []greeting.Friend{
	{Name: "Mary", Surname: "Ann", Birthday: time.Date(1974, 11, 21, 0, 0, 0, 0, time.Local)},
	{Name: "John", Surname: "Doe", Birthday: time.Date(1974, 10, 1, 0, 0, 0, 0, time.Local)},
}

type MockFriendRepository struct {
	mock.Mock
}

func (m *MockFriendRepository) Get() ([]greeting.Friend, error) {

	m.Called()
	return friends, nil
}

type MockGreetingSender struct {
	mock.Mock
}

func (m *MockGreetingSender) Send(name string) error {
	m.Called(name)
	return nil
}
