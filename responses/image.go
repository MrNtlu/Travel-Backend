package responses

import "time"

type Image struct {
	UserID      int       `json:"user_id"`
	User        User      `json:"user" gorm:"constraint:OnDelete:CASCADE;foreignKey:UserID"`
	LocationID  int       `json:"location_id"`
	Location    Location  `json:"location" gorm:"constraint:OnDelete:CASCADE;"`
	ImageURL    string    `json:"image_url"`
	Place       string    `json:"place"`
	Description *string   `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
