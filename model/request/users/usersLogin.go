package users

type UsersLogin struct {
	Username string `validate:"required,max=10"`
	Password string `validate:"required"`
}
