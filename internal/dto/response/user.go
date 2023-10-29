package dto

type UserResponse struct {
	ID        string `json:"user_id"`
	UUID      string `json:"uuid"`
	Name      string `json:"full_name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
