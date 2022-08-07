package models

import (
	"time"
	"travel_backend/requests"
	"travel_backend/responses"
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
	LocationID  int       `json:"location_id"`
	ImageURL    string    `json:"image_url" gorm:"unique"`
	AWSImageKey string    `json:"aws_key" gorm:"unique"`
	Place       string    `json:"place"`
	Description *string   `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

func (imageModel *ImageModel) createImageObject(data requests.ImageCreate, awsKey, imageURL string, uid int) *Image {
	return &Image{
		UserID:      uid,
		LocationID:  data.LocationID,
		ImageURL:    imageURL,
		AWSImageKey: awsKey,
		Place:       data.Place,
		Description: data.Description,
	}
}

func (imageModel *ImageModel) CreateImage(data requests.ImageCreate, awsKey, imageURL string, uid int) error {
	image := imageModel.createImageObject(data, awsKey, imageURL, uid)

	result := imageModel.Database.Create(&image)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (imageModel *ImageModel) GetImageByID(uid, id int) (Image, error) {
	var image Image

	result := imageModel.Database.Where("user_id = ?", uid).Where("id = ?", id).First(&image)
	if result.Error != nil {
		return Image{}, result.Error
	}

	return image, nil
}

func (imageModel *ImageModel) GetImagesByUserID(uid, page int) ([]responses.Image, error) {
	var images []responses.Image

	result := imageModel.Database.Scopes(utils.Paginate(page)).Preload(clause.Associations).Where("user_id = ?", uid).Find(&images)
	if result.Error != nil {
		return nil, result.Error
	}

	return images, nil
}

func (imageModel *ImageModel) GetImagesByCountry(uid, page int, country string) ([]Image, error) {
	var images []Image

	rawSQL := `Select * From Images as i
	Inner Join Locations as l on l.id=i.location_id
	Inner Join Users as u on u.id=i.user_id
	Where country = ?`

	result := imageModel.Database.Raw(rawSQL, country).Scan(&images)
	if result.Error != nil {
		return nil, result.Error
	}

	return images, nil
}

func (imageModel *ImageModel) GetImagesByLocation(uid, page, locationID int) ([]Image, error) {
	var images []Image

	result := imageModel.Database.Scopes(utils.Paginate(page)).Preload(clause.Associations).
		Where("user_id = ?", uid).Where("location_id = ?", locationID).Find(&images)
	if result.Error != nil {
		return nil, result.Error
	}

	return images, nil
}

func (imageModel *ImageModel) DeleteImageByID(uid, id int) error {
	result := imageModel.Database.Where("user_id = ?", uid).Where("id = ?", id).Delete(&Image{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
