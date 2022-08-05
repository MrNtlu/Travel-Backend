package models

import (
	"time"
)

type Pin struct {
	BaseModel

	UserID        int        `json:"user_id" gorm:"constraint:OnDelete:CASCADE;"`
	LocationID    int        `json:"location_id"`
	IsPlanToVisit bool       `json:"is_plan_to_visit"`
	Date          *time.Time `json:"date"`
}
