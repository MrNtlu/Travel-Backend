package responses

type Location struct {
	ID          uint    `json:"id" gorm:"primarykey"`
	CountryISO2 string  `json:"country_iso2"`
	CountryISO3 string  `json:"country_iso3"`
	Country     string  `json:"country"`
	Admin       string  `json:"admin"`
	City        string  `json:"city"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
}

type LocationAreaCity struct {
	ID   uint   `json:"id" gorm:"primarykey"`
	Area string `json:"area"`
	City string `json:"city"`
}

type LocationCity struct {
	ID   uint   `json:"id" gorm:"primarykey"`
	Area string `json:"area"`
}

type LocationCountry struct {
	Country string `json:"country"`
}
