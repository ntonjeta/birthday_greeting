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

func greeting(friends []Friend, date time.Time) (string, error) {

	for _, f := range friends {
		if isBirthday(f, date) {
			return fmt.Sprintf("Subject: Happy birthday!\nHappy birthday, dear %s!", f.name), nil
		}
	}
	return " ", nil // TODO erros ?
}

func isBirthday(friend Friend, date time.Time) bool {
	return friend.birthday.Day() == date.Day() &&
		friend.birthday.Month() == date.Month()
}

func main() {
	fmt.Printf("Birthday greeting! \n")
}
