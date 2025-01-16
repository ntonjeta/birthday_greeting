package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func TestWithMySql(t *testing.T) {
	ctx := context.Background()

	container, err := startMySql(ctx)
	if err != nil {
		fmt.Printf("error create mysql container: %s", err)
	}
	defer stopMySql(container)

	cs, err := container.ConnectionString(ctx)
	if err != nil {
		log.Fatal("fail retrieve connection string")
	}

	db, err := sql.Open("mysql", cs)
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected!")
}
