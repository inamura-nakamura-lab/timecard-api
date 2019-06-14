package response

import "time"

type GetAttendance struct {
	DateFrom      time.Time    `json:"date_from"`
	DateTo        time.Time    `json:"date_to"`
	ID            uint         `json:"id"`
	Name          string       `json:"name"`
	StudentNumber string       `json:"student_number"`
	Attendances   []Attendance `json:"attendance"`
}
