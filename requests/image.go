package requests

import "mime/multipart"

type ImageCreate struct {
	LocationID  int                  `form:"location_id" validate:"required"`
	Image       multipart.FileHeader `form:"image" validate:"required" swaggerignore:"true"`
	Place       string               `form:"place" validate:"required"`
	Description *string              `form:"description"`
}

type ImageByCountry struct {
	Page    int    `json:"page" validate:"required"`
	Country string `json:"country" validate:"required"`
}

type ImageByLocation struct {
	Page     int `json:"page" validate:"required"`
	Location int `json:"location" validate:"required,number,gte=0"`
}
