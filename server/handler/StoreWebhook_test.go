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

func TestStoreWebhook(t *testing.T) {
	db := database.MemoryConnection()
	db.AutoMigrate(&models.Webhook{})

	webhookRepo := repository.NewWebhookRepository(db)
	webhookHandler := NewWebhookHandler(*webhookRepo)

	t.Run("correctly stores the record", func(t *testing.T) {
		name := "new webhook name"
		rr := storeWebhookRequest(t, webhookHandler, name)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Expected status %v, but got %v", http.StatusOK, status)
		}

		var webhook models.Webhook
		db.Where("name = ?", name).First(&webhook)

		if webhook.Name != name {
			t.Fatal("Webhook was not created successfully")
		}
	})

	t.Run("name cannot be empty string", func(t *testing.T) {
		rr := storeWebhookRequest(t, webhookHandler, " ")

		if status := rr.Code; status != http.StatusUnprocessableEntity {
			t.Errorf("Expected status %v, but got %v", http.StatusUnprocessableEntity, status)
		}
	})

	t.Run("name gets trimmed", func(t *testing.T) {
		name := "  test "
		storeWebhookRequest(t, webhookHandler, name)

		var webhook models.Webhook
		db.Where("name = ?", name).First(&webhook)

		if webhook.Name != name {
			t.Fatal("Webhook was not created successfully")
		}
	})

	t.Run("cannot store duplicate names", func(t *testing.T) {
		name := "test name"
		storeWebhookRequest(t, webhookHandler, name)
		rr := storeWebhookRequest(t, webhookHandler, name)

		if status := rr.Code; status != http.StatusNotAcceptable {
			t.Errorf("Expected status %v, but got %v", http.StatusNotAcceptable, status)
		}
	})
}

func storeWebhookRequest(t *testing.T, webhookHandler WebhookHandler, name string) *httptest.ResponseRecorder {
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

	return rr
}
