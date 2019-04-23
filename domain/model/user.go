package model

import "time"

type User struct {
	ID         uint      `json:"id"`
	Name       string    `json:"name"`
	StudentNum string    `json:"student_number"`
	Date       time.Time `json:"date"`
	TimeCards []TimeCard `gorm:"foreignkey:UserID"`
}
