package users

import "time"

type UsersUpdate struct {
	Id        int    `validate:"required"`
	Username  string `validate:"required,max=10"`
	Password  string `validate:"required"`
	Name      string `validate:"required,max=50,min=1"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
