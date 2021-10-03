package mapper

import (
	"iluvatar/src/domain/models"
	"iluvatar/src/infrastructure/http/repository/miutem/auth/entity"
)

func MapLoginEntityToStudentModel(loginEntity *entity.LoginEntity) *models.Student {
	return &models.Student{
		FullName: loginEntity.NombreCompleto,
		Email:    loginEntity.Correo,
	}
}
