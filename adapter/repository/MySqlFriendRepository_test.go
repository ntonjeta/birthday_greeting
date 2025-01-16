package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	mySqlTestContainer "github.com/testcontainers/testcontainers-go/modules/mysql"
)

func arrange() (*sql.DB, mySqlTestContainer.MySQLContainer) {
	ctx := context.Background()

	container, err := startMySql(ctx)
	if err != nil {
		fmt.Printf("error create mysql container: %s", err)
	}

	cs, err := container.ConnectionString(ctx)
	if err != nil {
		log.Fatal("fail retrieve connection string")
	}

	db, err := sql.Open("mysql", cs)
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

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected!")
}
