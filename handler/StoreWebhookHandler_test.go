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
	webhookRepo := repository.NewWebhookRepository(database.MemoryConnection())
	webhookRepo.DB.AutoMigrate(&models.Webhook{})

	t.Run("correctly stores the record", func(t *testing.T) {
		name := "new webhook name"
		storeWebhookHandler := &StoreWebhookHandler{
			Repo: *webhookRepo,
		}

		requestBody := repository.CreateWebhookBody{
			Name: name,
		}

		body, _ := json.Marshal(requestBody)

		req, err := http.NewRequest("POST", "/webhooks", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(storeWebhookHandler.ServeHTTP)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Expected status %v, but got %v", http.StatusOK, status)
		}

		var webhook models.Webhook
		webhookRepo.DB.Where("name = ?", name).First(&webhook)

		if webhook.Name != name {
			t.Fatal("Webhook was not created successfully")
		}
	})
}
