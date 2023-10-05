package db

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	DataOfBirth string `json:"data_of_birth" binding:"required"`
	Email       string `json:"email" binding:"required"`
}

type EditUserRequest struct {
	Username    string `json:"username"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	DataOfBirth string `json:"data_of_birth"`
	Email       string `json:"email"`
}
