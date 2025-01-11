package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var friends = []Friend{
	{name: "Mary", surname: "Ann"},
	{name: "John", surname: "Doe"},
}

func TestAcceptance(t *testing.T) {
	date := time.Now()

	var mail, e = greeting(friends, date)

	assert.NotNil(t, e)
	assert.Empty(t, mail)
}
