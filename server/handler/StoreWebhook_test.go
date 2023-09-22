package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mudandstars/notifier/database"
	"github.com/mudandstars/notifier/models"
	"github.com/mudandstars/notifier/repository"
)

func TestStoreWebhookHandler(t *testing.T) {
	db := database.MemoryConnection()
	db.AutoMigrate(&models.Webhook{})

	webhookRepo := repository.NewWebhookRepository(db)
	webhookHandler := NewWebhookHandler(*webhookRepo)

	t.Run("correctly stores the record", func(t *testing.T) {
		name := "new webhook name"

		requestBody := models.Webhook{
			Name: name,
		}

		body, _ := json.Marshal(requestBody)

		req, err := http.NewRequest("POST", "/webhooks", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(webhookHandler.Store)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Expected status %v, but got %v", http.StatusOK, status)
		}

		var webhook models.Webhook
		db.Where("name = ?", name).First(&webhook)

		if webhook.Name != name {
			t.Fatal("Webhook was not created successfully")
		}
	})
}
