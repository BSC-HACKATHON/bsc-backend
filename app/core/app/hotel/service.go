package hotel

import (
	"github.com/lai0xn/bsc-auth/domain/enteties"
)

type HotelService interface {
	GetHotelsByCity(city string) ([]enteties.Hotel, error)
	GetHotelsByRating(minRating float32) ([]enteties.Hotel, error)
	UpdateHotel(id int, d enteties.Hotel) (enteties.Hotel, error)
	DeleteHotel(id int) (enteties.Hotel, error)
}

type hotelService struct {
	HotelRepo HotelRepo
}

func NewHotelService(repo HotelRepo) HotelService {
	return &hotelService{
		HotelRepo: repo,
	}
}

func (s *hotelService) GetHotelsByCity(city string) ([]enteties.Hotel, error) {
	// Use the repo method to fetch hotels by city
	return s.HotelRepo.FindHotelsByCity(city)
}

func (s *hotelService) GetHotelsByRating(minRating float32) ([]enteties.Hotel, error) {
	// Use the repo method to fetch hotels by rating
	return s.HotelRepo.FindHotelsByRating(minRating)
}

func (s *hotelService) CreateHotel(dto enteties.Hotel) (enteties.Hotel, error) {
	// Use the GenericRepo's Create method to create a hotel
	return s.HotelRepo.Create(dto)
}

func (s *hotelService) UpdateHotel(id int, dto enteties.Hotel) (enteties.Hotel, error) {
	// Use the GenericRepo's Update method to update a hotel
	return s.HotelRepo.Update(id, dto)
}

func (s *hotelService) DeleteHotel(id int) (enteties.Hotel, error) {
	// Use the GenericRepo's Delete method to delete a hotel
	return s.HotelRepo.Delete(id)
}
