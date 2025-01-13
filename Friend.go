package greeting

import "time"

type Friend struct {
	Name     string
	Surname  string
	Birthday time.Time
}

func (friend *Friend) IsBirthday(date time.Time) bool {
	return friend.Birthday.Day() == date.Day() &&
		friend.Birthday.Month() == date.Month()
}
