package response

import "time"

type Attendance struct {
	Date    time.Time `json:"date"`
	EnterAt time.Time `json:"enter_at"`
	ExitAt  time.Time `json:"exit_at"`
}
