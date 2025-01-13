package repository

import (
	"time"

	g "github.com/ntonjeta/greeting"
)

type ConsoleFriendRepository struct {
}

func (c *ConsoleFriendRepository) Get() ([]g.Friend, error) {
	return []g.Friend{
		{Name: "Mary", Surname: "Ann", Birthday: time.Date(1974, 11, 21, 0, 0, 0, 0, time.Local)},
		{Name: "John", Surname: "Doe", Birthday: time.Date(1974, 10, 1, 0, 0, 0, 0, time.Local)},
	}, nil
}
