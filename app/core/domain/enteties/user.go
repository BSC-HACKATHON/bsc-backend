package enteties

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email        string `gorm:"unique"`
	FirstName    string
	LastName     string
	Password     string
	DateOfBirth  time.Time
	PlaceOfBirth string
	IDNumber     int
	Gender       string
	IsActive     bool
	IsStaff      bool
}
