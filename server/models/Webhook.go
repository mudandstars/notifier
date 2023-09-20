package models

import (
	"gorm.io/gorm"
)

type Webhook struct {
	gorm.Model
	Name string `gorm:"not null"`
}
