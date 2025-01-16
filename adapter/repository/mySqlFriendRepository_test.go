package repository

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ntonjeta/greeting"
	"github.com/stretchr/testify/assert"
	mySqlTestContainer "github.com/testcontainers/testcontainers-go/modules/mysql"
)

func arrange() (*sql.DB, mySqlTestContainer.MySQLContainer) {
	ctx := context.Background()

	container, err := startMySql(ctx)
	if err != nil {
		log.Printf("error create mysql container: %s", err)
		os.Exit(1)
	}

	cs, err := container.ConnectionString(ctx)
	if err != nil {
		log.Fatal("fail retrieve connection string")
	}

	db, err := sql.Open("mysql", cs+"?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	return db, container
}

func dispose(container mySqlTestContainer.MySQLContainer) {
	stopMySql(container)
}

func TestWithMySql(t *testing.T) {
	var db, container = arrange()
	defer dispose(container)

	var repository = MySqlFriendsRepository{db}

	var expectedFriends = []greeting.Friend{
		{Name: "Mary", Surname: "Ann", Birthday: time.Date(1974, 11, 11, 0, 0, 0, 0, time.UTC)},
		{Name: "John", Surname: "Doe", Birthday: time.Date(1974, 10, 1, 0, 0, 0, 0, time.UTC)},
	}
	var friends, _ = repository.Get()

	assert.Equal(t, expectedFriends, friends)
}
