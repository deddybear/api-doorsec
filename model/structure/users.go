package structure

import "time"

type Users struct {
	Id        int
	Username  string
	Password  string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
