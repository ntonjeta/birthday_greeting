package main

import "time"

type Friend struct {
	name     string
	surname  string
	birthday time.Time
}

func (friend *Friend) isBirthday(date time.Time) bool {
	return friend.birthday.Day() == date.Day() &&
		friend.birthday.Month() == date.Month()
}
