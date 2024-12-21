package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lai0xn/bsc-auth/api"
	"github.com/lai0xn/bsc-auth/app/user/dto"
	"github.com/lai0xn/bsc-auth/pkg/utils"
)

func (h *APIHandler) Login(w http.ResponseWriter, r *http.Request) error {
	var payload dto.LoginDto
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return api.InvalidJSON(api.Map{
			"error": err.Error(),
		})
	}
	fmt.Println(payload)
	user, err := h.userService.Authenticate(payload.Email, payload.Password)
	if err != nil {
		return api.InvalidJSON(api.Map{
			"error": "wrong credentials",
		})
	}
	token, err := utils.GenerateJWT(user.ID, user.Email, fmt.Sprintf("%s %s", user.FirstName, user.LastName))
	if err != nil {
		return api.InvalidJSON(api.Map{
			"error": "wrong credentials",
		})
	}

	api.WriteJSON(w, http.StatusCreated, api.Map{
		"token": token,
	})
	return nil

}

func (h *APIHandler) Register(w http.ResponseWriter, r *http.Request) error {
	var payload dto.UserDto
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return api.InvalidJSON(api.Map{
			"error": err.Error(),
		})
	}
	fmt.Println(payload)
	if err := h.userService.Register(payload); err != nil {
		return api.InvalidJSON(api.Map{
			"error": err.Error(),
		})
	}
	api.WriteJSON(w, http.StatusCreated, api.Map{
		"message": "user created",
	})
	return nil

}
