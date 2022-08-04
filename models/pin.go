package models

import (
	"time"
)

type Pin struct {
	BaseModel

	UserID        int `gorm:"constraint:OnDelete:CASCADE;"`
	LocationID    int
	IsPlanToVisit bool
	Date          *time.Time
}
