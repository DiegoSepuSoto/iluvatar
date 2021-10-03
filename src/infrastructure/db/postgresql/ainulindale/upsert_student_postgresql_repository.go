package ainulindale

import (
	"crypto/md5"
	"fmt"
	"github.com/google/uuid"
	"iluvatar/src/domain/models"
	"iluvatar/src/infrastructure/db/postgresql/ainulindale/entity"
)

func (r *ainulindalePostgresqlRepository) UpsertStudent(student *models.Student) (string, error) {
	encryptedEmail, err := generateEncryptedUserName(student.Email)
	if err != nil {
		return "", err
	}
	db := r.dbConnection.GetConnection()

	var studentEntity = &entity.StudentEntity{
		ID:     uuid.New().String(),
		Email:  encryptedEmail,
		Career: student.Career,
	}

	var queryStudent entity.StudentEntity

	query := db.Where("email = ?", encryptedEmail).Limit(1).Find(&queryStudent)
	if query.Error != nil {
		return "", query.Error
	}

	if query.RowsAffected == 0 {
		dbResponse := db.Create(studentEntity)

		if dbResponse.Error != nil {
			return "", dbResponse.Error
		}

		return studentEntity.ID, nil
	}

	return queryStudent.ID, nil
}

func generateEncryptedUserName(studentEmail string) (string, error) {
	encryption := md5.Sum([]byte(studentEmail))
	return fmt.Sprintf("%x", encryption), nil
}
