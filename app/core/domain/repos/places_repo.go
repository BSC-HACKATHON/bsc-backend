package repos

import (
	"github.com/lai0xn/bsc-auth/db"
	"github.com/lai0xn/bsc-auth/domain/enteties"
)

type touristicPlaceRepository struct {
	GenericRepo[enteties.TouristicPlace]
}

func NewTouristicPlaceRepository() *touristicPlaceRepository {
	return &touristicPlaceRepository{
		GenericRepo: *(new(GenericRepo[enteties.TouristicPlace])),
	}
}

func (r *touristicPlaceRepository) FindPlacesByCity(city string) ([]enteties.TouristicPlace, error) {
	database := db.GetDB()
	var places []enteties.TouristicPlace
	result := database.Where("city = ?", city).Find(&places)
	if result.Error != nil {
		return nil, result.Error
	}
	return places, nil
}

func (r *touristicPlaceRepository) FindPlacesByType(placeType string) ([]enteties.TouristicPlace, error) {
	database := db.GetDB()
	var places []enteties.TouristicPlace
	result := database.Where("type = ?", placeType).Find(&places)
	if result.Error != nil {
		return nil, result.Error
	}
	return places, nil
}
