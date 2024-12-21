package repos

import (
	"github.com/lai0xn/bsc-auth/db"
	"github.com/lai0xn/bsc-auth/domain/enteties"
)

type hotelRepository struct {
	GenericRepo[enteties.Hotel]
}

func NewHotelRepository() *hotelRepository {
	return &hotelRepository{
		GenericRepo: *(new(GenericRepo[enteties.Hotel])),
	}
}

func (r *hotelRepository) FindHotelsByCity(city string) ([]enteties.Hotel, error) {
	database := db.GetDB()
	var hotels []enteties.Hotel
	result := database.Where("city = ?", city).Find(&hotels)
	if result.Error != nil {
		return nil, result.Error
	}
	return hotels, nil
}

func (r *hotelRepository) FindHotelsByRating(minRating float32) ([]enteties.Hotel, error) {
	database := db.GetDB()
	var hotels []enteties.Hotel
	result := database.Where("rating >= ?", minRating).Find(&hotels)
	if result.Error != nil {
		return nil, result.Error
	}
	return hotels, nil
}
