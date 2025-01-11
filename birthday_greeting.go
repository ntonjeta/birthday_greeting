package main

import (
	"fmt"
	"time"
)

type Friend struct {
	name     string
	surname  string
	birthday time.Time
}

func (friend *Friend) isBirthday(date time.Time) bool {
	return friend.birthday.Day() == date.Day() &&
		friend.birthday.Month() == date.Month()
}

func greeting(friends []Friend, date time.Time) (string, error) {

	for _, f := range friends {
		if f.isBirthday(date) {
			return fmt.Sprintf("Subject: Happy birthday!\nHappy birthday, dear %s!", f.name), nil
		}
	}
	return " ", nil // TODO erros ?
}

func main() {
	fmt.Printf("Birthday greeting! \n")
}
