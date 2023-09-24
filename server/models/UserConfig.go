package models

import (
	"gorm.io/gorm"
)

type UserConfig struct {
	gorm.Model
	NgrokAuthToken string `gorm:"unique, not null"`
	NgrokPublicUrl string `gorm:"unique, not null"`
}
