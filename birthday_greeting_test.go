package greeting

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var friends = []Friend{
	{Name: "Mary", Surname: "Ann", Birthday: time.Date(1974, 11, 21, 0, 0, 0, 0, time.Local)},
	{Name: "John", Surname: "Doe", Birthday: time.Date(1974, 10, 1, 0, 0, 0, 0, time.Local)},
}

type MockFriendRepository struct {
	mock.Mock
}

func (m *MockFriendRepository) Get() ([]Friend, error) {

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

func TestAcceptance_MaryBirthday(t *testing.T) {
	var date = time.Date(2024, 11, 21, 0, 0, 0, 0, time.Local)

	mockFriendRepository := new(MockFriendRepository)
	mockFriendRepository.On("Get").Return(friends, nil)

	mockGreetingSender := new(MockGreetingSender)
	mockGreetingSender.On("Send", "Mary").Return(nil)

	var e = greeting(mockFriendRepository, mockGreetingSender, date)

	assert.Nil(t, e)

	mockFriendRepository.AssertExpectations(t)
	mockGreetingSender.AssertExpectations(t)
}

func TestAcceptance_JohnBirthday(t *testing.T) {
	var date = time.Date(2024, 10, 1, 0, 0, 0, 0, time.Local)

	mockFriendRepository := new(MockFriendRepository)
	mockFriendRepository.On("Get").Return(friends, nil)

	mockGreetingSender := new(MockGreetingSender)
	mockGreetingSender.On("Send", "John").Return(nil)

	var e = greeting(mockFriendRepository, mockGreetingSender, date)

	assert.Nil(t, e)

	mockFriendRepository.AssertExpectations(t)
	mockGreetingSender.AssertExpectations(t)
}
