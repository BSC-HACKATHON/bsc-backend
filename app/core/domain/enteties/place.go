package enteties

import "gorm.io/gorm"

type TouristicPlace struct {
	gorm.Model
	Name         string  `gorm:"size:255;not null" json:"name"`
	Description  string  `gorm:"type:text" json:"description"`
	Address      string  `gorm:"size:255" json:"address"`
	City         string  `gorm:"size:100;not null" json:"city"`
	State        string  `gorm:"size:100" json:"state"`
	Country      string  `gorm:"size:100;not null" json:"country"`
	PostalCode   string  `gorm:"size:20" json:"postal_code"`
	Rating       float32 `gorm:"type:decimal(2,1);default:0.0" json:"rating"`
	OpeningHours string  `gorm:"size:255" json:"opening_hours"`
	EntryFee     float64 `gorm:"type:decimal(10,2);default:0.0" json:"entry_fee"`
	Website      string  `gorm:"size:255" json:"website"`
	ContactInfo  string  `gorm:"size:255" json:"contact_info"`
	NearbyHotels []Hotel `gorm:"many2many:hotel_touristic_places;" json:"nearby_hotels"`
}
