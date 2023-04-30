package request

import "time"

type UsersCreate struct {
	Username  string `validate:"required,max=10,min=4"`
	Password  string `validate:"required"`
	Name      string `validate:"required,max=50,min=1"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
