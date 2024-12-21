package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/lai0xn/bsc-auth/api"
	"github.com/lai0xn/bsc-auth/api/handlers"
	"github.com/lai0xn/bsc-auth/app/user"
	"github.com/lai0xn/bsc-auth/domain/repos"
)

func Setup(r chi.Router) {

	usrRepo := repos.NewUserRepository()
	usrSrv := user.NewUserService(usrRepo)

	handler := handlers.NewAPIHandler(usrSrv)

	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", api.MakeHandler(handler.Login))
		r.Post("/register", api.MakeHandler(handler.Register))
	})
}
