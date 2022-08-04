package models

import (
	"time"
	"travel_backend/requests"
	"travel_backend/utils"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	BaseModel

	UserID      int       `json:"user_id"`
	User        User      `json:"user" gorm:"constraint:OnDelete:CASCADE;foreignKey:UserID"`
	LocationID  int       `json:"location_id"`
	Location    Location  `json:"location" gorm:"constraint:OnDelete:CASCADE;"`
	ImageURL    string    `json:"image_url" gorm:"unique"`
	Place       string    `json:"place"`
	Description *string   `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

func (imageModel *ImageModel) createImageObject(data requests.ImageCreate, imageURL string, uid int) *Image {
	return &Image{
		UserID:      uid,
		LocationID:  data.LocationID,
		ImageURL:    imageURL,
		Place:       data.Place,
		Description: data.Description,
	}
}

func (imageModel *ImageModel) CreateImage(data requests.ImageCreate, uid int) error {

	//TODO Upload image and get URL
	image := imageModel.createImageObject(data, data.Image.Filename, uid)

	result := imageModel.Database.Create(&image)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (imageModel *ImageModel) GetImagesByUserID(uid, page int) ([]Image, error) {
	var images []Image

	result := imageModel.Database.Scopes(utils.Paginate(page)).Preload(clause.Associations).Where("user_id = ?", uid).Find(&images)
	if result.Error != nil {
		return nil, result.Error
	}

	return images, nil
}
