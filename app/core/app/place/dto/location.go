package dto

type CreateTouristicPlaceDTO struct {
	Name         string  `json:"name" validate:"required"`
	Description  string  `json:"description"`
	Address      string  `json:"address" validate:"required"`
	City         string  `json:"city" validate:"required"`
	State        string  `json:"state"`
	Country      string  `json:"country" validate:"required"`
	PostalCode   string  `json:"postal_code"`
	Rating       float32 `json:"rating"`
	OpeningHours string  `json:"opening_hours"`
	EntryFee     float64 `json:"entry_fee"`
	Website      string  `json:"website"`
	ContactInfo  string  `json:"contact_info"`
}
