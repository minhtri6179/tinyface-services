package db

type UserResponse struct {
	Username    string `json:"username"`
	FullName    string `json:"full_name"`
	Email       string `json:"email"`
	DataOfBirth string `json:"data_of_birth"`
}
