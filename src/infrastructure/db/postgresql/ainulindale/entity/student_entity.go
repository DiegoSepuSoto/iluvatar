package entity

type StudentEntity struct {
	ID     string `gorm:"primaryKey"`
	Email  string
	Career string
}

func (StudentEntity) TableName() string {
	return "STUDENT"
}
