package repository

import (
	"github.com/mudandstars/notifier/models"
	"gorm.io/gorm"
)

type UserConfigRepository struct {
	db *gorm.DB
}

func NewConfigRepository(db *gorm.DB) *UserConfigRepository {
	return &UserConfigRepository{
		db: db,
	}
}

func (controller *UserConfigRepository) Store(body *models.Config) error {
	if err := controller.db.Create(&body).Error; err != nil {
		return err
	}

	return nil
}

func (controller *UserConfigRepository) Get() (models.Config, error) {
	var config []models.Config

	results := controller.db.Find(&config)

	if results.Error != nil {
		return models.Config{}, results.Error
	}

	if len(config) > 0 {
		return config[0], nil
	}
	return models.Config{}, nil
}

func (controller *UserConfigRepository) Delete() error {
	if err := controller.db.Delete(&models.Config{}, "0 = 0").Error; err != nil {
		return err
	}

	return nil
}
