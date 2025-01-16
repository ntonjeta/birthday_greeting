package repository

import (
	"database/sql"
	"log"

	"github.com/ntonjeta/greeting"
)

type MySqlFriendsRepository struct {
	db *sql.DB
}

func (repository *MySqlFriendsRepository) Get() ([]greeting.Friend, error) {
	rows, err := repository.db.Query("SELECT NAME, SURNAME, BIRTHDAY FROM FRIENDS;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var friends []greeting.Friend

	for rows.Next() {
		var friend greeting.Friend
		if err := rows.Scan(&friend.Name, &friend.Surname, &friend.Birthday); err != nil {
			return nil, err
		}
		friends = append(friends, friend)
	}
	if err := rows.Err(); err != nil {
		return friends, err
	}

	return friends, nil
}
