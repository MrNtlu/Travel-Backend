package requests

type LocationCountry struct {
	Country string `json:"country" validate:"required"`
}
