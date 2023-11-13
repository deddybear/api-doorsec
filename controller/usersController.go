package controller

import (
	exception "api-doorsec/exceptions"
	"api-doorsec/helper"
	"api-doorsec/model/request/users"
	"api-doorsec/model/response"
	"api-doorsec/services"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type UsersControllerInterface interface {
	SignUp(res http.ResponseWriter, req *http.Request, params httprouter.Params)
	SignIn(res http.ResponseWriter, req *http.Request, params httprouter.Params)
	FindById(res http.ResponseWriter, req *http.Request, params httprouter.Params)
	Update(res http.ResponseWriter, req *http.Request, params httprouter.Params)
	Delete(res http.ResponseWriter, req *http.Request, params httprouter.Params)
}

type UsersController struct {
	UsersService services.UsersService
}

func NewUsersController(usersServices services.UsersService) UsersControllerInterface {
	return &UsersController{
		UsersService: usersServices,
	}
}

func (controller *UsersController) SignUp(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	userCreateReq := users.UsersCreate{}

	helper.ReadFromRequestBody(req, &userCreateReq)

	userResponse := controller.UsersService.SignUp(req.Context(), userCreateReq)

	apiRes := response.Api{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(res, apiRes)
}

func (controller *UsersController) SignIn(res http.ResponseWriter, req *http.Request, params httprouter.Params) {

	userLoginReq := users.UsersLogin{}

	helper.ReadFromRequestBody(req, &userLoginReq)

	userResponse := controller.UsersService.SignIn(req.Context(), userLoginReq)

	Token, err := helper.CreateTokenJwt(userResponse.Username, userResponse.Name)

	if err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}

	apiRes := response.Login{
		Code:   200,
		Token:  Token,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(res, apiRes)
}

func (controller *UsersController) FindById(res http.ResponseWriter, req *http.Request, params httprouter.Params) {

}

func (controller *UsersController) Update(res http.ResponseWriter, req *http.Request, params httprouter.Params) {

}

func (controller *UsersController) Delete(res http.ResponseWriter, req *http.Request, params httprouter.Params) {

}
