package repository

import (
	"testing"

	"github.com/mudandstars/notifier/database"
	"github.com/mudandstars/notifier/models"
)

func TestRepository(t *testing.T) {
	db := database.MemoryConnection()
	db.AutoMigrate(&models.Webhook{})

	webhookRepo := NewWebhookRepository(db)

	t.Run("correctly stores the webhook", func(t *testing.T) {
		webhook := models.Webhook{Name: "Test Name"}

		error := webhookRepo.Store(&webhook)
		if error != nil {
			t.Fatal(error)
		}
	})

	t.Run("correctly fetches all entries", func(t *testing.T) {
		if _, error := webhookRepo.All(); error != nil {
			t.Fatal(error)
		}
	})

	t.Run("correctly deletes the webhook", func(t *testing.T) {
		webhook := models.Webhook{Name: "Test Name"}

		webhookRepo.Store(&webhook)

		error := webhookRepo.Delete(webhook.ID)
		if error != nil {
			t.Fatal(error)
		}
	})
}
