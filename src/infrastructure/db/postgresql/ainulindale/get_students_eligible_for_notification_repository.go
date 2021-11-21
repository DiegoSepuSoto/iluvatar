package ainulindale

import (
	"iluvatar/src/infrastructure/db/postgresql/ainulindale/entity"
	"iluvatar/src/infrastructure/db/postgresql/ainulindale/mapper"
)

func (r *ainulindalePostgresqlRepository) GetStudentsEligibleForNotification() ([]string, error) {
	db := r.dbConnection.GetConnection()

	var students []entity.StudentEntity

	query := db.Where("device_id is not null").Where("device_id <> ?", "").Find(&students)
	if query.Error != nil {
		return nil, query.Error
	}

	return mapper.GetDevicesIDFromStudentsEntitySlice(students), nil
}
