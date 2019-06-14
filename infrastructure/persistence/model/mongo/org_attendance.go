package mongo

import "time"

type OrgAttendance struct {
	DateFrom      time.Time    `bson:"date_from"`
	DateTo        time.Time    `bson:"date_to"`
	ID            uint         `bson:"id"`
	Name          string       `bson:"name"`
	StudentNumber string       `bson:"student_number"`
	Attendances   []Attendance `bson:"attendance"`
}
