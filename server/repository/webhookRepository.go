package repository

import (
	"github.com/mudandstars/notifier/models"

	"gorm.io/gorm"
)

type WebhookRepository struct {
	db *gorm.DB
}

func NewWebhookRepository(db *gorm.DB) *WebhookRepository {
	return &WebhookRepository{
		db: db,
	}
}

func (controller *WebhookRepository) Store(body *models.Webhook) error {
	if err := controller.db.Create(&body).Error; err != nil {
		return err
	}

	return nil
}

func (controller *WebhookRepository) All() ([]models.Webhook, error) {
	var users []models.Webhook

	result := controller.db.Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (controller *WebhookRepository) Delete(id uint) error {
	if err := controller.db.Delete(&models.Webhook{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (controller *WebhookRepository) Exists(name string) bool {
	count := int64(0)

	controller.db.Model(&models.Webhook{}).
		Where("name = ?", name).
		Count(&count)

	return count > 0
}
