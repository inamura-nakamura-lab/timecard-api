package model

type TimeCard struct {
	BaseModel
	Attendances []Attendance `gorm:"foreignkey:TimeCardID"`
	UserID      uint
}
