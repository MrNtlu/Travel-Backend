package responses

type User struct {
	ID           uint   `json:"id" gorm:"primarykey"`
	EmailAddress string `json:"email_address"`
	Username     string `json:"username"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
}
