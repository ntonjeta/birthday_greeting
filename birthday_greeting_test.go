package greeting_test

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/ntonjeta/greeting"
	"github.com/ntonjeta/greeting/mock"
)

var friends = []greeting.Friend{
	{Name: "Mary", Surname: "Ann", Birthday: time.Date(1974, 11, 21, 0, 0, 0, 0, time.Local)},
	{Name: "John", Surname: "Doe", Birthday: time.Date(1974, 10, 1, 0, 0, 0, 0, time.Local)},
}

func TestAcceptance_MaryBirthday(t *testing.T) {
	var date = time.Date(2024, 11, 21, 0, 0, 0, 0, time.Local)

	mockFriendRepository := new(mock.MockFriendRepository)
	mockFriendRepository.On("Get").Return(friends, nil)

	mockGreetingSender := new(mock.MockGreetingSender)
	mockGreetingSender.On("Send", "Mary").Return(nil)

	var e = greeting.Greeting(mockFriendRepository, mockGreetingSender, date)

	assert.Nil(t, e)

	mockFriendRepository.AssertExpectations(t)
	mockGreetingSender.AssertExpectations(t)
}

func TestAcceptance_JohnBirthday(t *testing.T) {
	var date = time.Date(2024, 10, 1, 0, 0, 0, 0, time.Local)

	mockFriendRepository := new(mock.MockFriendRepository)
	mockFriendRepository.On("Get").Return(friends, nil)

	mockGreetingSender := new(mock.MockGreetingSender)
	mockGreetingSender.On("Send", "John").Return(nil)

	var e = greeting.Greeting(mockFriendRepository, mockGreetingSender, date)

	assert.Nil(t, e)

	mockFriendRepository.AssertExpectations(t)
	mockGreetingSender.AssertExpectations(t)
}

func TestAcceptance_RepositoryError(t *testing.T) {
	var date = time.Date(2024, 10, 1, 0, 0, 0, 0, time.Local)

	mockFriendRepository := new(mock.MockFriendRepository)
	mockFriendRepository.On("Get").Return([]greeting.Friend{}, errors.New("error reading friend file"))

	mockGreetingSender := new(mock.MockGreetingSender)

	var e = greeting.Greeting(mockFriendRepository, mockGreetingSender, date)

	assert.NotNil(t, e)

	mockFriendRepository.AssertExpectations(t)
}
