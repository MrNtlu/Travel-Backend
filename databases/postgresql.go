package databases

import (
	"fmt"
	"strconv"
	"travel_backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "188Burak"
	dbname   = "travel_db"
)

type PostgreSQL struct {
	Database *gorm.DB
}

func SetDatabase() (*PostgreSQL, error) {
	dbConnection := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, strconv.Itoa(port),
	)

	db, err := gorm.Open(postgres.Open(dbConnection), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Make migrations if necessary
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Location{})
	db.AutoMigrate(&models.Image{})

	return &PostgreSQL{
		Database: db,
	}, nil
}
