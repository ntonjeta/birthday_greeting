package main

import (
	"fmt"
	"time"
)

type FriendRepository interface {
	get() ([]Friend, error)
}

func greeting(
	friendRepo FriendRepository,
	date time.Time,
) (string, error) {

	// TODO test error repository
	var friends, _ = friendRepo.get()
	for _, friend := range friends {
		if friend.isBirthday(date) {
			return fmt.Sprintf("Subject: Happy birthday!\nHappy birthday, dear %s!", friend.name), nil
		}
	}
	// TODO no birthday test
	return "", nil
}

func main() {
	fmt.Printf("Birthday greeting! \n")
}
