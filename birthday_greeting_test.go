package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var friends = []Friend{
	{name: "Mary", surname: "Ann"},
	{name: "John", surname: "Doe"},
}

func TestAcceptance(t *testing.T) {
	var mail, e = greeting(friends)

	assert.Nil(t, e)

	assert.Equal(
		t,
		mail,
		"Subject: Happy birthday!\nHappy birthday, dear Mary!",
		"mails shoulld be equals")
}
