package repository

import "iluvatar/src/domain/models"

type AuthRepository interface {
	Login(email, password string) (*models.Student, string, error)
}
