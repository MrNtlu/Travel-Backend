package responses

import "time"

type Pin struct {
	UserID        int        `json:"user_id"`
	User          User       `json:"user" gorm:"constraint:OnDelete:CASCADE;foreignKey:UserID"`
	LocationID    int        `json:"location_id"`
	Location      Location   `json:"location" gorm:"constraint:OnDelete:CASCADE;"`
	IsPlanToVisit bool       `json:"is_plan_to_visit"`
	Date          *time.Time `json:"date"`
}
