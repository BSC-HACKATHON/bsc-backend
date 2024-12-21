package place

import (
	"github.com/lai0xn/bsc-auth/domain/enteties"
)

type TouristicPlaceService interface {
	GetPlacesByCity(city string) ([]enteties.TouristicPlace, error)
	GetPlacesByType(locationType string) ([]enteties.TouristicPlace, error)
	UpdatePlace(id int, d enteties.TouristicPlace) (enteties.TouristicPlace, error)
	DeletePlace(id int) (enteties.TouristicPlace, error)
}

type touristicPlaceService struct {
	LocationRepo LocationRepo
}

func NewTouristicPlaceService(repo LocationRepo) TouristicPlaceService {
	return &touristicPlaceService{
		LocationRepo: repo,
	}
}

func (s *touristicPlaceService) GetPlacesByCity(city string) ([]enteties.TouristicPlace, error) {
	// Use the repo method to fetch places by city
	return s.LocationRepo.FindPlacesByCity(city)
}

func (s *touristicPlaceService) GetPlacesByType(locationType string) ([]enteties.TouristicPlace, error) {
	// Use the repo method to fetch places by type
	return s.LocationRepo.FindPlacesByType(locationType)
}

func (s *touristicPlaceService) CreatePlace(dto enteties.TouristicPlace) (enteties.TouristicPlace, error) {
	// Use the GenericRepo's Create method to create a touristic place
	return s.LocationRepo.Create(dto)
}

func (s *touristicPlaceService) UpdatePlace(id int, dto enteties.TouristicPlace) (enteties.TouristicPlace, error) {
	// Use the GenericRepo's Update method to update a touristic place
	return s.LocationRepo.Update(id, dto)
}

func (s *touristicPlaceService) DeletePlace(id int) (enteties.TouristicPlace, error) {
	// Use the GenericRepo's Delete method to delete a touristic place
	return s.LocationRepo.Delete(id)
}
