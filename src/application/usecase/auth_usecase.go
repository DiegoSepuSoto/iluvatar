package usecase

import "iluvatar/src/domain/models"

type AuthUseCase interface {
	Login(email, password string) (*models.Student, error)
}
