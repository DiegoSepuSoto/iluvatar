package entity

type StudentEntity struct {
	ID       string `gorm:"primaryKey"`
	Email    string
	Career   string
	DeviceID string
}

func (StudentEntity) TableName() string {
	return "STUDENT"
}
