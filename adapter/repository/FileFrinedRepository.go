package repository

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ntonjeta/greeting"
)

type FileFriendRepository struct {
}

func (f *FileFriendRepository) Get() ([]greeting.Friend, error) {
	var friends []greeting.Friend

	content, err := os.ReadFile("friends.txt")
	if err != nil {
		fmt.Printf("error %s", err.Error())
		return nil, err
	}

	for _, l := range strings.Split(strings.TrimSpace(string(content)), "\n")[1:] {
		var attr = strings.Split(l, ",")

		friends = append(friends, greeting.Friend{
			Name:     string(attr[0]),
			Surname:  string(attr[1]),
			Birthday: birthday(attr[2]),
		})
	}
	return friends, nil
}

func birthday(birthday string) time.Time {
	attr := strings.Split(birthday, "-")

	var year, _ = strconv.Atoi(attr[0])
	var month, _ = strconv.Atoi(attr[1])
	var day, _ = strconv.Atoi(attr[2])

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
}
