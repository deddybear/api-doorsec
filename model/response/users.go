package response

import "time"

type Users struct {
	Id        int
	Username  string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
