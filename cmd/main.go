package main

import (
	"time"
)

func main() {
	var date = time.Date(2023, 11, 21, 0, 0, 0, 0, time.Local)
	var friends ConsoleFriendRepository
	var sender ConsoleGreetingSender

	greeting(&friends, &sender, date)
}
