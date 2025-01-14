package greeting

import (
	"time"
)

func Greeting(
	friendRepository FriendRepository,
	greetingSender GreetingSender,
	date time.Time,
) error {

	// TODO test error repository
	var friends, _ = friendRepository.Get()

	for _, friend := range friends {
		if friend.IsBirthday(date) {
			greetingSender.Send(friend.Name)
		}
	}
	// TODO no birthday test
	return nil
}
