package repository

import (
	"github.com/mudandstars/notifier/models"
	"gorm.io/gorm"
)

type UserConfigRepository struct {
	db *gorm.DB
}

func NewUserConfigRepository(db *gorm.DB) *UserConfigRepository {
	return &UserConfigRepository{
		db: db,
	}
}

func (controller *UserConfigRepository) Store(body *models.UserConfig) error {
	if err := controller.db.Create(&body).Error; err != nil {
		return err
	}

	return nil
}

func (controller *UserConfigRepository) Get() (models.UserConfig, error) {
	var config []models.UserConfig

	results := controller.db.Find(&config)

	if results.Error != nil {
		return models.UserConfig{}, results.Error
	}

	if len(config) > 0 {
		return config[0], nil
	}
	return models.UserConfig{}, nil
}

func (controller *UserConfigRepository) Delete() error {
	if err := controller.db.Delete(&models.UserConfig{}, "0 = 0").Error; err != nil {
		return err
	}

	return nil
}
