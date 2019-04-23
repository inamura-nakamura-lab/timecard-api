package model

type User struct {
	BaseModel
	Name       string `gorm:"not null"`
	StudentNum string `gorm:"not null"`
}
