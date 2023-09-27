package repository

import (
	"github.com/mudandstars/notifier/models"
	"gorm.io/gorm"
)

type ConfigRepository struct {
	DB *gorm.DB
}

func NewConfigRepository(db *gorm.DB) *ConfigRepository {
	return &ConfigRepository{
		DB: db,
	}
}

func (repo *ConfigRepository) Upsert(body *models.Config) error {
	var existingConfigs []models.Config
	if err := repo.DB.Find(&existingConfigs).Error; err != nil {
		return err
	}

	if len(existingConfigs) > 0 {
		// If records exist, update the first one
		if err := repo.DB.Model(&existingConfigs[0]).Updates(body).Error; err != nil {
			return err
		}
	} else {
		// If no records exist, create a new one
		if err := repo.DB.Create(body).Error; err != nil {
			return err
		}
	}

	return nil
}

func (repo *ConfigRepository) Get() (models.Config, error) {
	var config []models.Config

	results := repo.DB.Find(&config)

	if results.Error != nil {
		return models.Config{}, results.Error
	}

	if len(config) > 0 {
		return config[0], nil
	}
	return models.Config{}, nil
}

func (repo *ConfigRepository) Delete() error {
	if err := repo.DB.Delete(&models.Config{}, "0 = 0").Error; err != nil {
		return err
	}

	return nil
}
