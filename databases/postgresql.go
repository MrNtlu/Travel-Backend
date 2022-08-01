package databases

import (
	"context"
	"database/sql"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "postgres"
	port     = 5432
	user     = "postgres"
	password = "188Burak"
	dbname   = "travel_db"
)

type PostgreSQL struct {
	Database *sql.DB
}

func SetDatabase() (*PostgreSQL, error) {
	connectionString := "postgres://" + user + ":" + password + "@localhost:" + strconv.Itoa(port) + "/" + dbname + "?sslmode=disable"
	// dockerConnectionString := "postgres://" + user + ":" + password + "@" + host + ":" + strconv.Itoa(port) + "/" + dbname + "?sslmode=disable"

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return &PostgreSQL{
		Database: db,
	}, nil
}
