package dto

type CreateHotelDTO struct {
	Name          string  `json:"name" validate:"required"`
	Address       string  `json:"address" validate:"required"`
	City          string  `json:"city" validate:"required"`
	State         string  `json:"state"`
	Country       string  `json:"country" validate:"required"`
	PostalCode    string  `json:"postal_code"`
	Rating        float32 `json:"rating"`
	NumberOfRooms int     `json:"number_of_rooms" validate:"required"`
	Amenities     string  `json:"amenities"`
	ContactInfo   string  `json:"contact_info"`
	Website       string  `json:"website"`
}
