package requests

import "mime/multipart"

type ImageCreate struct {
	LocationID  int                  `form:"location_id" validate:"required"`
	Image       multipart.FileHeader `form:"image" validate:"required" swaggerignore:"true"`
	Place       string               `form:"place" validate:"required"`
	Description *string              `form:"description"`
}
