package controller

import "go-hospital-server/internal/core/service"

type Controller struct {
	Auth *AuthController
	User *UserController
}

func NewController(srv *service.Service) *Controller {
	return &Controller{
		Auth: NewAuthController(srv.Auth),
		User: NewUserController(srv.User),
	}
}
