package greeting

import (
	"log"
	"time"
)

func Greeting(
	friendRepository FriendRepository,
	greetingSender GreetingSender,
	date time.Time,
) error {

	var friends, err = friendRepository.Get()
	if err != nil {
		log.Println(err)
		return err
	}

	for _, friend := range friends {
		if friend.IsBirthday(date) {
			greetingSender.Send(friend.Name)
		}
	}
	// TODO no birthday test
	return nil
}
