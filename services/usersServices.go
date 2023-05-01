package services

import (
	exception "api-iotdoor/exceptions"
	"api-iotdoor/helper"
	"api-iotdoor/model/request/users"
	"api-iotdoor/model/response"
	"api-iotdoor/model/structure"
	"api-iotdoor/repository"
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"time"
)

type UsersService interface {
	SignUp(ctx context.Context, req users.UsersCreate) response.Users
	SignIn(ctx context.Context, req users.UsersLogin) response.Users
	FindById(ctx context.Context, usersId int) response.Users
	Update(ctx context.Context, req users.UsersUpdate) response.Users
	Delete(ctx context.Context, usersId int)
}

type UsersImpl struct {
	UsersRepository repository.UsersRepositoryInterface
	DB              *sql.DB
	Validate        *validator.Validate
}

func NewUsersServices(usersRepository repository.UsersRepositoryInterface, DB *sql.DB, validate *validator.Validate) UsersService {
	return &UsersImpl{
		UsersRepository: usersRepository,
		DB:              DB,
		Validate:        validate,
	}
}

func (services *UsersImpl) SignUp(ctx context.Context, req users.UsersCreate) response.Users {
	err := services.Validate.Struct(req)
	helper.PanicIfError(err)

	tx, err := services.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	err = services.UsersRepository.FindByUsername(ctx, tx, req.Username)
	if err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}

	user := structure.Users{
		Username:  req.Username,
		Password:  req.Password,
		Name:      req.Name,
		CreatedAt: req.CreatedAt,
		UpdatedAt: req.UpdatedAt,
	}

	user = services.UsersRepository.SignUp(ctx, tx, user)

	return helper.ToUsersResponse(user)
}

func (services *UsersImpl) SignIn(ctx context.Context, req users.UsersLogin) response.Users {
	tx, err := services.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := services.UsersRepository.SignIn(ctx, tx, req.Username, req.Password)

	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToUsersResponse(user)
}

func (services *UsersImpl) FindById(ctx context.Context, userId int) response.Users {
	tx, err := services.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := services.UsersRepository.FindById(ctx, tx, userId)

	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToUsersResponse(user)
}

func (services *UsersImpl) Update(ctx context.Context, req users.UsersUpdate) response.Users {
	err := services.Validate.Struct(req)
	helper.PanicIfError(err)

	tx, err := services.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := services.UsersRepository.FindById(ctx, tx, req.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	user.Username = req.Username
	user.Name = req.Name
	user.Password = req.Password
	user.UpdatedAt = time.Now()

	services.UsersRepository.Update(ctx, tx, user)

	return helper.ToUsersResponse(user)
}

func (services *UsersImpl) Delete(ctx context.Context, userId int) {
	tx, err := services.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := services.UsersRepository.FindById(ctx, tx, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	services.UsersRepository.Delete(ctx, tx, category)
}
