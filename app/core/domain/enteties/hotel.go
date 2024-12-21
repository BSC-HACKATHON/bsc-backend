package enteties

import "gorm.io/gorm"

type Hotel struct {
	gorm.Model
	Name          string  `gorm:"size:255;not null" json:"name"`
	Address       string  `gorm:"size:255;not null" json:"address"`
	City          string  `gorm:"size:100;not null" json:"city"`
	State         string  `gorm:"size:100" json:"state"`
	Country       string  `gorm:"size:100;not null" json:"country"`
	PostalCode    string  `gorm:"size:20" json:"postal_code"`
	Rating        float32 `gorm:"type:decimal(2,1);default:0.0" json:"rating"`
	NumberOfRooms int     `gorm:"not null" json:"number_of_rooms"`
	Amenities     string  `gorm:"type:text" json:"amenities"`
	ContactInfo   string  `gorm:"size:255" json:"contact_info"`
	Website       string  `gorm:"size:255" json:"website"`
}
