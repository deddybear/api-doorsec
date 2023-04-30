package request

type UsersUpdate struct {
	Id        int    `validate:"required"`
	Username  string `validate:"required,max=8,min=1"`
	Password  string `validate:"required,uuid4"`
	Name      string `validate:"required,max=50,min=1"`
	CreatedAt string `validate:"required,timestamp"`
	UpdatedAt string `validate:"required,timestamp"`
}
