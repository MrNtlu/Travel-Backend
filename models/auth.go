package models

import (
	"fmt"
	"time"
	"travel_backend/requests"
	"travel_backend/utils"

	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

type UserModel struct {
	Database *gorm.DB
}

func NewUserModel(database *gorm.DB) *UserModel {
	return &UserModel{
		Database: database,
	}
}

type User struct {
	gorm.Model

	ID           int    `gorm:"primaryKey"`
	EmailAddress string `gorm:"unique"`
	Username     string `gorm:"unique"`
	FirstName    string
	LastName     string
	Password     string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	FCMToken     string
}

func createUserObject(data requests.Register) *User {
	return &User{
		EmailAddress: data.EmailAddress,
		Username:     data.Username,
		FirstName:    data.FirstName,
		LastName:     data.LastName,
		Password:     utils.HashPassword(data.Password),
		CreatedAt:    time.Now().UTC(),
		UpdatedAt:    time.Now().UTC(),
		FCMToken:     data.FCMToken,
	}
}

func (userModel *UserModel) CreateUser(data requests.Register) error {
	user := createUserObject(data)

	result := userModel.Database.Create(&user)
	if result.Error != nil {
		if ok := result.Error.(*pgconn.PgError); ok != nil && ok.Code == "23505" {
			return fmt.Errorf(ok.ConstraintName + " unique constraint failed!")
		}

		return result.Error
	}

	return nil
}

func (userModel *UserModel) FindUserByEmail(email string) (User, error) {
	var user User

	result := userModel.Database.Where("email_address = ?", email).First(&user)
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}
