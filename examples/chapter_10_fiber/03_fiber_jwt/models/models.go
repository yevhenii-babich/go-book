package models

// LoginRequest struct is a model for login request
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResponse struct is a model for login response
type LoginResponse struct {
	Token string `json:"token"`
}

// User struct is a model for user
type User struct {
	ID       int
	Email    string
	Password string
	NickName string
}
