package user

type UserUpdate struct {
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	Address    string    `json:"address"`
	City       string    `json:"city"`
	Country    string    `json:"country"`
	ZipCode    string    `json:"zip_code"`
	Role string `json:"role"`
}
