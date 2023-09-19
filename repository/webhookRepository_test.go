package repository

import (
	"testing"

	"github.com/mudandstars/notifier/database"
	"github.com/mudandstars/notifier/models"
)

func TestRepository(t *testing.T) {
	webhookRepo := NewWebhookRepository(database.MemoryConnection())
	webhookRepo.DB.AutoMigrate(&models.Webhook{})

	t.Run("correctly stores webhook", func(t *testing.T) {
		webhookName := "Test Name"
		webhookRepo.Store(CreateWebhookBody{Name: webhookName})

		if error := webhookRepo.Store(CreateWebhookBody{Name: webhookName}); error != nil {
			t.Fatal(error)
		}

		if webhooks, _ := webhookRepo.All(); webhooks[0].Name != webhookName {
			t.Fatal("Names don't match.")
		}
	})

	t.Run("correctly fetches all entries", func(t *testing.T) {
		if _, error := webhookRepo.All(); error != nil {
			t.Fatal(error)
		}
	})
}
