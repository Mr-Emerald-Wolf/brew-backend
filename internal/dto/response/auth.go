package dto

type AuthResponse struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}

type RefreshResponse struct {
	AccessToken string `json:"access_token"`
}
