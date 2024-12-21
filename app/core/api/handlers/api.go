package handlers

import (
	"github.com/lai0xn/bsc-auth/app/hotel"
	"github.com/lai0xn/bsc-auth/app/place"
	"github.com/lai0xn/bsc-auth/app/user"
)

type APIHandler struct {
	userService  user.UserService
	hotelService hotel.HotelService
	placeService place.TouristicPlaceService
}

func NewAPIHandler(usrv user.UserService, htlSrv hotel.HotelService, plcSrv place.TouristicPlaceService) *APIHandler {
	return &APIHandler{
		userService:  usrv,
		hotelService: htlSrv,
		placeService: plcSrv,
	}
}
