package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/lai0xn/bsc-auth/api"
	"github.com/lai0xn/bsc-auth/api/handlers"
	"github.com/lai0xn/bsc-auth/app/hotel"
	"github.com/lai0xn/bsc-auth/app/place"
	"github.com/lai0xn/bsc-auth/app/user"
	"github.com/lai0xn/bsc-auth/domain/repos"
)

func Setup(r chi.Router) {

	// User repository and service
	usrRepo := repos.NewUserRepository()
	usrSrv := user.NewUserService(usrRepo)

	// Hotel repository and service
	hotelRepo := repos.NewHotelRepository() // Assuming you have this repository
	hotelSrv := hotel.NewHotelService(hotelRepo)

	// Touristic Place repository and service
	placeRepo := repos.NewTouristicPlaceRepository() // Assuming you have this repository
	placeSrv := place.NewTouristicPlaceService(placeRepo)

	// Create a new APIHandler with all services
	handler := handlers.NewAPIHandler(usrSrv, hotelSrv, placeSrv)

	// Authentication Routes
	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", api.MakeHandler(handler.Login))
		r.Post("/register", api.MakeHandler(handler.Register))
	})

	// Hotel Routes
	r.Route("/hotels", func(r chi.Router) {
		r.Get("/", api.MakeHandler(handler.GetHotelsByCity)) // Get hotels by city
		r.Put("/", api.MakeHandler(handler.UpdateHotel))     // Update hotel
		r.Delete("/", api.MakeHandler(handler.DeleteHotel))  // Delete hotel
	})

	// Touristic Place Routes
	r.Route("/places", func(r chi.Router) {
		r.Get("/", api.MakeHandler(handler.GetPlacesByCity))     // Get places by city
		r.Get("/type", api.MakeHandler(handler.GetPlacesByType)) // Get places by type
		r.Put("/", api.MakeHandler(handler.UpdatePlace))         // Update place
		r.Delete("/", api.MakeHandler(handler.DeletePlace))      // Delete place
	})
}
