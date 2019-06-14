package model

import "time"

type User struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	StudentNumber string    `json:"student_number"`
	Date          time.Time `json:"date"`
}
