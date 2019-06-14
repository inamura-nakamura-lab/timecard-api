package mongo

import "time"

type Attendance struct {
	Date    time.Time `bson:"date"`
	EnterAt time.Time `bson:"enter_at"`
	ExitAt  time.Time `bson:"exit_at"`
}
