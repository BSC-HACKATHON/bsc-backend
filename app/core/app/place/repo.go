package place

import (
	"github.com/lai0xn/bsc-auth/domain/enteties"
	"github.com/lai0xn/bsc-auth/domain/repos"
)

type LocationRepo interface {
	FindPlacesByCity(city string) ([]enteties.TouristicPlace, error)
	FindPlacesByType(locationType string) ([]enteties.TouristicPlace, error)
	repos.GenericRepoI[enteties.TouristicPlace]
}
