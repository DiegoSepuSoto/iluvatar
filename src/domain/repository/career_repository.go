package repository

type CareerRepository interface {
	GetCareer(accessToken string) (string, error)
}