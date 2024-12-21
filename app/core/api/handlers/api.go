package handlers

import "github.com/lai0xn/bsc-auth/app/user"

type APIHandler struct {
	userService user.UserService
}

func NewAPIHandler(usrv user.UserService) *APIHandler {
	return &APIHandler{
		userService: usrv,
	}
}
