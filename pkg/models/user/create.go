package user

type UserCreate struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Identifier string `json:"identifier"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
	City       string `json:"city"`
	Country    string `json:"country"`
	ZipCode    string `json:"zip_code"`
	Password   string `json:"password"`
	Role       string `json:"role"`
}
