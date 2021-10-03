package models

type Student struct {
	FullName     string `json:"full_name"`
	Email        string `json:"email"`
	Career       string `json:"career"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
