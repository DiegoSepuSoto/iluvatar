package usecase

import (
	"iluvatar/src/domain/models"
	"iluvatar/src/domain/models/requests"
)

type AuthUseCase interface {
	Login(loginRequest *requests.StudentLoginRequest) (*models.Student, error)
	RefreshToken(studentID string) (string, error)
}
