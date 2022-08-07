package models

import (
	"travel_backend/responses"

	"gorm.io/gorm"
)

type LocationModel struct {
	Database *gorm.DB
}

func NewLocationModel(database *gorm.DB) *LocationModel {
	return &LocationModel{
		Database: database,
	}
}

type Location struct {
	BaseModel

	CountryISO2 string  `json:"country_iso2"`
	CountryISO3 string  `json:"country_iso3"`
	Country     string  `json:"country" gorm:"index:country_admin_city,unique"`
	Admin       string  `json:"admin" gorm:"index:country_admin_city,unique"`
	City        string  `json:"city" gorm:"index:country_admin_city,unique"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
}

func (locationModel *LocationModel) GetCountryList() ([]responses.LocationCountry, error) {
	var locationCountryList []responses.LocationCountry

	rawSQL := `SELECT DISTINCT country FROM locations ORDER BY country ASC`

	result := locationModel.Database.Raw(rawSQL).Scan(&locationCountryList)
	if result.Error != nil {
		return nil, result.Error
	}

	return locationCountryList, nil
}

func (locationModel *LocationModel) GetAreaCityList(country string) ([]responses.LocationAreaCity, error) {
	var locationAreaCityList []responses.LocationAreaCity

	rawSQL := `SELECT id, admin as area, city FROM locations WHERE country = ?`

	result := locationModel.Database.Raw(rawSQL, country).Scan(&locationAreaCityList)
	if result.Error != nil {
		return nil, result.Error
	}

	return locationAreaCityList, nil
}

func (locationModel *LocationModel) GetCityList(country string) ([]responses.LocationCity, error) {
	var locationCityList []responses.LocationCity

	rawSQL := `SELECT DISTINCT ON (admin, country) id, admin as area, country
	FROM locations
	WHERE country = ? AND city=admin
	ORDER BY country ASC, admin ASC`

	result := locationModel.Database.Raw(rawSQL, country).Scan(&locationCityList)
	if result.Error != nil {
		return nil, result.Error
	}

	return locationCityList, nil
}
