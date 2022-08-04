package requests

type Pagination struct {
	Page int `json:"page" validate:"required"`
}
