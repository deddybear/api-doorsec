package routes

import (
	"api-iotdoor/controller"
	"api-iotdoor/repository"
	"api-iotdoor/services"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func UsersRoutes(router *httprouter.Router, db *sql.DB, validate *validator.Validate) {

	usersRepository := repository.NewUsersRepository()
	usersService := services.NewUsersServices(usersRepository, db, validate)
	usersController := controller.NewUsersController(usersService)

	router.POST("/api/users/signin", usersController.SignIn)
	router.POST("/api/users/signup", usersController.SignUp)
	//router.GET("/api/users/:usersId", usersController.FindById)
	router.PATCH("/api/users/update/:usersId", usersController.Update)
	router.DELETE("/api/users/delete/:usersId", usersController.Delete)
}
