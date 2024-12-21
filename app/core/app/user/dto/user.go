package dto

import "time"

type UserDto struct {
	Email        string    `json:"email"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Password     string    `json:"password"`
	DateOfBirth  time.Time `json:"date_of_birth"`
	PlaceOfBirth string    `json:"place_of_birth"`
	Gender       string    `json:"gender"`
}
