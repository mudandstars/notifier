package repository

import "gorm.io/gorm"

type Repository interface {
	Store(body struct{}) error
	All() ([]gorm.Model, error)
}
