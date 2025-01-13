package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var friends = []Friend{
	{name: "Mary", surname: "Ann", birthday: time.Date(1974, 11, 21, 0, 0, 0, 0, time.Local)},
	{name: "John", surname: "Doe", birthday: time.Date(1974, 10, 1, 0, 0, 0, 0, time.Local)},
}

type MockFriendRepository struct {
	mock.Mock
}

func (m *MockFriendRepository) get() ([]Friend, error) {

	m.Called()
	return friends, nil
}

type MockGreetingSender struct {
	mock.Mock
}

func (m *MockGreetingSender) send(name string) error {
	m.Called(name)
	return nil
}

func TestAcceptance_MaryBirthday(t *testing.T) {
	var date = time.Date(2024, 11, 21, 0, 0, 0, 0, time.Local)

	mockFriendRepository := new(MockFriendRepository)
	mockFriendRepository.On("get").Return(friends, nil)

	mockGreetingSender := new(MockGreetingSender)
	mockGreetingSender.On("send", "Mary").Return(nil)

	var e = greeting(mockFriendRepository, mockGreetingSender, date)

	assert.Nil(t, e)

	mockFriendRepository.AssertExpectations(t)
	mockGreetingSender.AssertExpectations(t)
	// mockGreetingSender.AssertCalled(t, "send", "mary")
}

func TestAcceptance_JohnBirthday(t *testing.T) {
	var date = time.Date(2024, 10, 1, 0, 0, 0, 0, time.Local)

	mockFriendRepository := new(MockFriendRepository)
	mockFriendRepository.On("get").Return(friends, nil)

	mockGreetingSender := new(MockGreetingSender)
	mockGreetingSender.On("send", "John").Return(nil)

	var e = greeting(mockFriendRepository, mockGreetingSender, date)

	assert.Nil(t, e)

	mockFriendRepository.AssertExpectations(t)
	mockGreetingSender.AssertExpectations(t)
}
