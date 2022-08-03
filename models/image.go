package models

import (
	"time"

	"gorm.io/gorm"
)

type ImageModel struct {
	Database *gorm.DB
}

func NewImageModel(database *gorm.DB) *ImageModel {
	return &ImageModel{
		Database: database,
	}
}

type Image struct {
	gorm.Model

	ID         int `gorm:"primaryKey"`
	UserID     int `gorm:"constraint:OnDelete:CASCADE;"`
	LocationID int
	ImageURL   string `gorm:"unique"`
	CreatedAt  time.Time
}
