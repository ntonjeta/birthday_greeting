package main

import (
	"time"

	"github.com/ntonjeta/greeting"
	"github.com/ntonjeta/greeting/adapter/repository"
	"github.com/ntonjeta/greeting/adapter/sender"
)

func main() {
	var date = time.Date(2023, 11, 21, 0, 0, 0, 0, time.Local)
	var friends repository.FileFriendRepository
	var sender sender.ConsoleGreetingSender

	greeting.Greeting(&friends, &sender, date)
}
