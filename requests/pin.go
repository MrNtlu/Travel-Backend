package requests

type PinCreate struct {
	LocationID    int   `form:"location_id" validate:"required"`
	IsPlanToVisit *bool `form:"is_plan_to_visit"`
}
