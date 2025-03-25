package controller

import (
	"net/http"
	"pjt1/app/service"
	"pjt1/pkg/api"
	"pjt1/pkg/e"
)

type UserController interface {
	SaveUserDetails(w http.ResponseWriter, r *http.Request)
	LoginUser(w http.ResponseWriter, r *http.Request)
	ExampleHandler(w http.ResponseWriter, r *http.Request)
}

type UserControllerImpl struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		userService: userService,
	}
}

func (c *UserControllerImpl) SaveUserDetails(w http.ResponseWriter, r *http.Request) {
	resp, err := c.userService.SaveUserDetails(r)
	if err != nil {
		apiErr := e.NewAPIError(err, "failed to create user")
		api.Fail(w, apiErr.StatusCode, apiErr.Code, apiErr.Message, err.Error())
		return
	}
	api.Success(w, http.StatusOK, resp)
}

func (c *UserControllerImpl) LoginUser(w http.ResponseWriter, r *http.Request) {
	resp, err := c.userService.LoginUser(r)
	if err != nil {
		apiErr := e.NewAPIError(err, "failed to login user")
		api.Fail(w, apiErr.StatusCode, apiErr.Code, apiErr.Message, err.Error())
		return
	}
	api.Success(w, http.StatusOK, resp)
}

func (c *UserControllerImpl) ExampleHandler(w http.ResponseWriter, r *http.Request) {
	err := c.userService.ExampleHandler(r)
	if err != nil {
		apiErr := e.NewAPIError(err, "failed to login user")
		api.Fail(w, apiErr.StatusCode, apiErr.Code, apiErr.Message, err.Error())
		return
	}
	api.Success(w, http.StatusOK, "resp")
}
