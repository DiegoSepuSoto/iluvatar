package requests

type StudentLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
