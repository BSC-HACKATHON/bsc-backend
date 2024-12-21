package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/lai0xn/bsc-auth/api"
	"github.com/lai0xn/bsc-auth/app/hotel/dto"
	"github.com/lai0xn/bsc-auth/domain/enteties"
)

func (h *APIHandler) GetHotelsByCity(w http.ResponseWriter, r *http.Request) error {
	city := r.URL.Query().Get("city")
	if city == "" {
		return api.InvalidJSON(api.Map{
			"error": "City parameter is required",
		})
	}

	hotels, err := h.hotelService.GetHotelsByCity(city)
	if err != nil {
		return api.InvalidJSON(api.Map{
			"error": "Failed to fetch hotels",
		})
	}

	api.WriteJSON(w, http.StatusOK, api.Map{
		"hotels": hotels,
	})
	return nil
}

func (h *APIHandler) UpdateHotel(w http.ResponseWriter, r *http.Request) error {
	var payload dto.CreateHotelDTO
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return api.InvalidJSON(api.Map{
			"error": err.Error(),
		})
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		return api.InvalidJSON(api.Map{
			"error": "Hotel ID is required",
		})
	}
	uid, err := strconv.Atoi(id)
	if err != nil {
		return api.InvalidJSON(api.Map{
			"error": "Hotel ID is required",
		})
	}

	entity := enteties.Hotel{
		Name:          payload.Name,
		Address:       payload.Address,
		City:          payload.City,
		State:         payload.State,
		Rating:        payload.Rating,
		Country:       payload.Country,
		Website:       payload.Website,
		Amenities:     payload.Amenities,
		PostalCode:    payload.PostalCode,
		ContactInfo:   payload.ContactInfo,
		NumberOfRooms: payload.NumberOfRooms,
	}

	hotel, err := h.hotelService.UpdateHotel(uid, entity)
	if err != nil {
		return api.InvalidJSON(api.Map{
			"error": err.Error(),
		})
	}

	api.WriteJSON(w, http.StatusOK, api.Map{
		"message": "Hotel updated",
		"hotel":   hotel,
	})
	return nil
}

func (h *APIHandler) DeleteHotel(w http.ResponseWriter, r *http.Request) error {
	id := r.URL.Query().Get("id")
	if id == "" {
		return api.InvalidJSON(api.Map{
			"error": "Hotel ID is required",
		})
	}
	uid, err := strconv.Atoi(id)
	if err != nil {
		return api.InvalidJSON(api.Map{
			"error": "Hotel ID is required",
		})
	}
	hotel, err := h.hotelService.DeleteHotel(uid)
	if err != nil {
		return api.InvalidJSON(api.Map{
			"error": err.Error(),
		})
	}

	api.WriteJSON(w, http.StatusOK, api.Map{
		"message": "Hotel deleted",
		"hotel":   hotel,
	})
	return nil
}
