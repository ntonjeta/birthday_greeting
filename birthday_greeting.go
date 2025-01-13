package main

import (
	"fmt"
	"time"
)

func greeting(friends []Friend, date time.Time) (string, error) {

	for _, friend := range friends {
		if friend.isBirthday(date) {
			return fmt.Sprintf("Subject: Happy birthday!\nHappy birthday, dear %s!", friend.name), nil
		}
	}
	// TODO no birthday test
	return "", nil
}

func main() {
	fmt.Printf("Birthday greeting! \n")
}
