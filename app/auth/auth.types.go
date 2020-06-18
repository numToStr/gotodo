package auth

// LoginDTO is the payload for login
type LoginDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"password"`
}

// LoginRes is the login response
type LoginRes struct {
	Hash string `json:"hash"`
}
