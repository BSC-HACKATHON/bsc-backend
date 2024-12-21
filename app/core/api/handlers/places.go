package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/lai0xn/bsc-auth/api"
	"github.com/lai0xn/bsc-auth/app/place/dto"
	"github.com/lai0xn/bsc-auth/domain/enteties"
)

func (h *APIHandler) GetPlacesByCity(w http.ResponseWriter, r *http.Request) error {
	city := r.URL.Query().Get("city")
	if city == "" {
		return api.InvalidJSON(api.Map{
			"error": "City parameter is required",
		})
	}

	places, err := h.placeService.GetPlacesByCity(city)
	if err != nil {
		return api.InvalidJSON(api.Map{
			"error": "Failed to fetch places",
		})
	}

	api.WriteJSON(w, http.StatusOK, api.Map{
		"places": places,
	})
	return nil
}

func (h *APIHandler) GetPlacesByType(w http.ResponseWriter, r *http.Request) error {
	locationType := r.URL.Query().Get("type")
	if locationType == "" {
		return api.InvalidJSON(api.Map{
			"error": "Location type parameter is required",
		})
	}

	places, err := h.placeService.GetPlacesByType(locationType)
	if err != nil {
		return api.InvalidJSON(api.Map{
			"error": "Failed to fetch places",
		})
	}

	api.WriteJSON(w, http.StatusOK, api.Map{
		"places": places,
	})
	return nil
}

func (h *APIHandler) UpdatePlace(w http.ResponseWriter, r *http.Request) error {
	var payload dto.CreateTouristicPlaceDTO
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return api.InvalidJSON(api.Map{
			"error": err.Error(),
		})
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		return api.InvalidJSON(api.Map{
			"error": "Touristic place ID is required",
		})
	}
	uid, err := strconv.Atoi(id)
	if err != nil {
		return api.InvalidJSON(api.Map{
			"error": "Hotel ID is required",
		})
	}

	entity := enteties.TouristicPlace{
		Name:         payload.Name,
		ContactInfo:  payload.ContactInfo,
		PostalCode:   payload.PostalCode,
		Address:      payload.Address,
		Website:      payload.Website,
		Country:      payload.Country,
		Rating:       payload.Rating,
		State:        payload.State,
		EntryFee:     payload.EntryFee,
		Description:  payload.Description,
		OpeningHours: payload.OpeningHours,
		City:         payload.City,
	}
	place, err := h.placeService.UpdatePlace(uid, entity)
	if err != nil {
		return api.InvalidJSON(api.Map{
			"error": err.Error(),
		})
	}

	api.WriteJSON(w, http.StatusCreated, api.Map{
		"message": "Touristic place created",
		"place":   place,
	})
	return nil
}

func (h *APIHandler) DeletePlace(w http.ResponseWriter, r *http.Request) error {
	id := r.URL.Query().Get("id")
	if id == "" {
		return api.InvalidJSON(api.Map{
			"error": "Touristic place ID is required",
		})
	}

	uid, err := strconv.Atoi(id)
	if err != nil {
		return api.InvalidJSON(api.Map{
			"error": "hotel id is required",
		})
	}
	place, err := h.placeService.DeletePlace(uid)
	if err != nil {
		return api.InvalidJSON(api.Map{
			"error": err.Error(),
		})
	}

	api.WriteJSON(w, http.StatusOK, api.Map{
		"message": "Touristic place deleted",
		"place":   place,
	})
	return nil
}
