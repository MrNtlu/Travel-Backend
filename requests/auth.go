package requests

type Login struct {
	EmailAddress string `json:"email_address" validate:"required,email"`
	Password     string `json:"password" validate:"required"`
}

type Register struct {
	EmailAddress string `json:"email_address" validate:"required,email"`
	Password     string `json:"password" validate:"required,min=6"`
}
