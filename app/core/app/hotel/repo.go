package hotel

import (
	"github.com/lai0xn/bsc-auth/domain/enteties"
	"github.com/lai0xn/bsc-auth/domain/repos"
)

type HotelRepo interface {
	FindHotelsByCity(city string) ([]enteties.Hotel, error)
	FindHotelsByRating(minRating float32) ([]enteties.Hotel, error)
	repos.GenericRepoI[enteties.Hotel]
}
