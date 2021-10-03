package entity

type StudentEntity struct {
	ID     int `gorm:"primaryKey"`
	Email  string
	Career string
}

func (StudentEntity) TableName() string {
	return "STUDENT"
}
