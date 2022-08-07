package models

import (
	"time"
	"travel_backend/requests"

	"gorm.io/gorm"
)

type PinModel struct {
	Database *gorm.DB
}

func NewPinModel(database *gorm.DB) *PinModel {
	return &PinModel{
		Database: database,
	}
}

type Pin struct {
	BaseModel

	UserID        int        `json:"user_id" gorm:"constraint:OnDelete:CASCADE;"`
	LocationID    int        `json:"location_id"`
	IsPlanToVisit bool       `json:"is_plan_to_visit"`
	Date          *time.Time `json:"date"`
}

func (pinModel *PinModel) createPinObject(data requests.PinCreate, uid int) *Pin {
	return &Pin{
		UserID:        uid,
		LocationID:    data.LocationID,
		IsPlanToVisit: *data.IsPlanToVisit,
	}
}

func (pinModel *PinModel) CreatePin(data requests.PinCreate, uid int) error {
	pin := pinModel.createPinObject(data, uid)

	result := pinModel.Database.Create(&pin)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
