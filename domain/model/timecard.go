package model

type TimeCard struct {
	ID          uint         `json:"id"`
	Attendances []Attendance `json:"attendances"`
}
