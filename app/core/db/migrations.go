package db

import "github.com/lai0xn/bsc-auth/domain/enteties"

func Migrate() {
	database.AutoMigrate(&enteties.User{})
}
