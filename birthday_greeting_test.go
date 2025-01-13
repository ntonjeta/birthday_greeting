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
	return friends, nil
}

func TestAcceptance_MaryBirthday(t *testing.T) {
	var date = time.Date(2024, 11, 21, 0, 0, 0, 0, time.Local)

	testRepository := new(MockFriendRepository)
	testRepository.On("get").Return(friends, nil)

	var mail, e = greeting(testRepository, date)

	assert.Nil(t, e)

	assert.Equal(
		t,
		mail,
		"Subject: Happy birthday!\nHappy birthday, dear Mary!",
		"mails shoulld be equals")
}

func TestAcceptance_JohnBirthday(t *testing.T) {
	var date = time.Date(2024, 10, 1, 0, 0, 0, 0, time.Local)

	testRepository := new(MockFriendRepository)
	testRepository.On("get").Return(friends, nil)

	var mail, e = greeting(testRepository, date)

	assert.Nil(t, e)

	assert.Equal(
		t,
		mail,
		"Subject: Happy birthday!\nHappy birthday, dear John!",
		"mails shoulld be equals")
}
