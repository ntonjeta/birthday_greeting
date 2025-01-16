package repository

import (
	"context"
	"log"
	"path/filepath"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mysql"
)

func startMySql(ctx context.Context) (mysql.MySQLContainer, error) {
	container, err := mysql.Run(ctx,
		"mysql:8.0.36",
		// mysql.WithConfigFile(filepath.Join("testdata", "my_8.cnf")),
		mysql.WithDatabase("test"),
		mysql.WithUsername("root"),
		mysql.WithPassword("password"),
		mysql.WithScripts(filepath.Join("testdata", "schema.sql")),
	)
	if err != nil {
		log.Printf("failed to start container: %s", err)
		return mysql.MySQLContainer{}, err
	}

	return *container, nil
}

func stopMySql(mysql mysql.MySQLContainer) {
	if err := testcontainers.TerminateContainer(mysql); err != nil {
		log.Printf("failed to terminate container: %s", err)
	}
}