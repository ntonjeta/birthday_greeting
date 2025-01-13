package main

import (
	"fmt"
	"time"
)

type GreetingSender interface {
	send(name string) error
}

func greeting(
	friendRepository FriendRepository,
	greetingSender GreetingSender,
	date time.Time,
) error {

	// TODO test error repository
	var friends, _ = friendRepository.get()
	for _, friend := range friends {
		if friend.isBirthday(date) {
			greetingSender.send(friend.name)
			// return fmt.Sprintf("Subject: Happy birthday!\nHappy birthday, dear %s!", friend.name), nil
		}
	}
	// TODO no birthday test
	return nil
}

func main() {
	fmt.Printf("Birthday greeting! \n")
}
