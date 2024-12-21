package repos

import (
	"github.com/lai0xn/bsc-auth/db"
	"github.com/lai0xn/bsc-auth/pkg/utils"
)

type GenericRepoI[T any] interface {
	Create(dto T) (T, error)
	Get(id int) (T, error)
	Update(id int, dto T) (T, error)
	Delete(id int) (T, error)
}

type GenericRepo[T any] struct {
}

func (r *GenericRepo[T]) Create(dto T) (T, error) {
	database := db.GetDB()
	result := database.Create(&dto)
	if result.Error != nil {
		return utils.Zero[T](), result.Error
	}
	return dto, nil
}

func (r *GenericRepo[T]) Get(id int) (T, error) {
	database := db.GetDB()
	var row T
	result := database.First(&row, "id = ?", id)
	if result.Error != nil {
		return utils.Zero[T](), result.Error
	}
	return row, nil
}

func (r *GenericRepo[T]) Update(id int, dto T) (T, error) {
	database := db.GetDB()
	var row T
	result := database.First(&row, "id = ?", id).Updates(dto)
	if result.Error != nil {
		return utils.Zero[T](), result.Error
	}
	return row, nil
}

func (r *GenericRepo[T]) Delete(id int) (T, error) {
	database := db.GetDB()
	var row T
	result := database.Delete(&row, "id = ?", id)
	if result.Error != nil {
		return row, result.Error
	}
	return row, nil
}
