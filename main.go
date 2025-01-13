package greeting

import (
	"time"

	r "github.com/ntonjeta/greeting/repository"
	s "github.com/ntonjeta/greeting/sender"
)

func main() {
	var date = time.Date(2024, 11, 21, 0, 0, 0, 0, time.Local)
	var friends r.ConsoleFriendRepository
	var sender s.ConsoleGreetingSender

	greeting(&friends, &sender, date)
}
