package auth

type AuthLogin struct {
	Identifier string `json:"identifier" validate:"required"`
	Password string `json:"password" validate:"required"`
}