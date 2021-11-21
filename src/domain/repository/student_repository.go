package repository

import "iluvatar/src/domain/models"

type StudentRepository interface {
	UpsertStudent(student *models.Student) (string, error)
	GetStudentsEligibleForNotification() ([]string, error)
}