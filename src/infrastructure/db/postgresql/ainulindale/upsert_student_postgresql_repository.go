package ainulindale

import (
	"golang.org/x/crypto/bcrypt"
	"iluvatar/src/domain/models"
	"iluvatar/src/infrastructure/db/postgresql/ainulindale/entity"
)

const encryptionCost = 14

func (r *ainulindalePostgresqlRepository) UpsertStudent(student *models.Student) error {
	encryptedEmail, err := generateEncryptedUserName(student.Email)
	if err != nil {
		return err
	}
	db := r.dbConnection.GetConnection()

	var studentEntity = &entity.StudentEntity{
		Email:  encryptedEmail,
		Career: student.Career,
	}

	dbResponse := db.Model(studentEntity).Where("email = ?", encryptedEmail).Updates(&studentEntity)

	if dbResponse.RowsAffected == 0 {
		db.Create(studentEntity)
	}

	if dbResponse.Error != nil {
		return dbResponse.Error
	}

	return nil
}

func generateEncryptedUserName(studentEmail string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(studentEmail), encryptionCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}