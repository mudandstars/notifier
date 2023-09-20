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

type CreateWebhookBody struct {
	Name string
}

func (controller *WebhookRepository) Store(body CreateWebhookBody) error {
	if err := controller.DB.Create(&models.Webhook{
		Name: body.Name,
	}).Error; err != nil {
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
