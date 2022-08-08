package requests

import "time"

type PinCreate struct {
	LocationID    int        `json:"location_id" validate:"required"`
	IsPlanToVisit *bool      `json:"is_plan_to_visit"`
	Date          *time.Time `json:"date" time_format:"2006-01-02"`
}
