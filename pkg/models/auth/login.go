package auth

type AuthLogin struct {
	Identifier string `json:"identifier"`
	Password string `json:"password"`
}