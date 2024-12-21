package main

import (
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/lai0xn/bsc-auth/api/routes"
	"github.com/lai0xn/bsc-auth/db"
)

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	routes.Setup(r)

	log.Info("Trying db connection")
	db.MustConnect()
	log.Info("Database Connected")
	log.Info("Running db migrations")
	db.Migrate()
	log.Info("Server started on port 8080")

	http.ListenAndServe(":8080", r)

}
