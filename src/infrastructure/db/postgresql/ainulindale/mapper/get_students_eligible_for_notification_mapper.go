package mapper

import "iluvatar/src/infrastructure/db/postgresql/ainulindale/entity"

func GetDevicesIDFromStudentsEntitySlice(students []entity.StudentEntity) []string {
	var studentsDeviceID []string

	for _, student := range students {
		studentsDeviceID = append(studentsDeviceID, student.DeviceID)
	}

	return studentsDeviceID
}
