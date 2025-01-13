package main

import "time"

type FriendRepository interface {
	get() ([]Friend, error)
}

type ConsoleFriendRepository struct {
}

func (c *ConsoleFriendRepository) get() ([]Friend, error) {
	return []Friend{
		{name: "Mary", surname: "Ann", birthday: time.Date(1974, 11, 21, 0, 0, 0, 0, time.Local)},
		{name: "John", surname: "Doe", birthday: time.Date(1974, 10, 1, 0, 0, 0, 0, time.Local)},
	}, nil
}
