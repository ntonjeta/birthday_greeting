package main

import (
	"time"
)

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
		}
	}
	// TODO no birthday test
	return nil
}

func main() {
	var date = time.Date(2024, 11, 21, 0, 0, 0, 0, time.Local)
	var friends ConsoleFriendRepository
	var sender ConsoleGreetingSender

	greeting(&friends, &sender, date)
}
