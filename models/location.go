package models

import "gorm.io/gorm"

type LocationModel struct {
	Database *gorm.DB
}

func NewLocationModel(database *gorm.DB) *LocationModel {
	return &LocationModel{
		Database: database,
	}
}

type Location struct {
	gorm.Model

	ID              int `gorm:"primaryKey"`
	CountryISO2Code string
	CountryISO3Code string
	Country         string `gorm:"index:country_admin_city,unique"`
	Admin           string `gorm:"index:country_admin_city,unique"`
	City            string `gorm:"index:country_admin_city,unique"`
	Latitude        float64
	Longitude       float64
}
