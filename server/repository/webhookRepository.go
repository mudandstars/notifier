package repository

import (
	"github.com/mudandstars/notifier/models"

	"gorm.io/gorm"
)

type WebhookRepository struct {
	DB *gorm.DB
}

func NewWebhookRepository(db *gorm.DB) *WebhookRepository {
	return &WebhookRepository{
		DB: db,
	}
}

func (controller *WebhookRepository) Store(body *models.Webhook) error {
	if err := controller.DB.Create(&body).Error; err != nil {
		return err
	}

	return nil
}

func (controller *WebhookRepository) All() ([]models.Webhook, error) {
	var users []models.Webhook

	result := controller.DB.Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (controller *WebhookRepository) Delete(id uint) error {
	if err := controller.DB.Delete(&models.Webhook{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (controller *WebhookRepository) Exists(name string) bool {
	count := int64(0)

	controller.DB.Model(&models.Webhook{}).
		Where("name = ?", name).
		Count(&count)

	return count > 0
}
