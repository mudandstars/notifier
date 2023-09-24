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
		webhook := models.Webhook{Name: "Test Name 1"}

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
		webhook := models.Webhook{Name: "Test Name 2"}

		webhookRepo.Store(&webhook)

		error := webhookRepo.Delete(webhook.ID)
		if error != nil {
			t.Fatal(error)
		}
	})

	t.Run("Exists() returns true if the entry exists", func(t *testing.T) {
		webhook := models.Webhook{Name: "Test Name 3"}

		error := webhookRepo.Store(&webhook)
		if error != nil {
			t.Fatal(error)
		}

		exists := webhookRepo.Exists(webhook.Name)
		if exists != true {
			t.Fatal("Expected .Exists() to return true")
		}
	})

	t.Run("Exists() returns false if the entry does not exist", func(t *testing.T) {
		exists := webhookRepo.Exists("some name that does not exist")
		if exists != false {
			t.Fatal("Expected .Exists() to return false")
		}
	})
}
