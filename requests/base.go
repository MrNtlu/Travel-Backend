package requests

type ID struct {
	ID int `json:"id" validate:"required,number"`
}

type Pagination struct {
	Page int `json:"page" validate:"required"`
}
