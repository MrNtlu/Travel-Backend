package requests

type Login struct {
	EmailAddress string `json:"email_address" validate:"required,email"`
	Password     string `json:"password" validate:"required"`
}

type Register struct {
	EmailAddress string `json:"email_address" validate:"required,email"`
	Password     string `json:"password" validate:"required,min=6"`
	Username     string `json:"username" validate:"required,min=3"`
	FirstName    string `json:"first_name" validate:"required"`
	LastName     string `json:"last_name" validate:"required"`
	FCMToken     string `json:"fcm_token" validate:"required"`
}
