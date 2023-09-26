package models

import (
	"gorm.io/gorm"
)

type Config struct {
	gorm.Model
	NgrokAuthToken string `gorm:"unique, not null"`
	NgrokPublicUrl string `gorm:"unique, not null"`
}
