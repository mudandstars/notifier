package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mudandstars/notifier/database"
	"github.com/mudandstars/notifier/models"
	"github.com/mudandstars/notifier/repository"
)

func TestIndexWebhook(t *testing.T) {
	db := database.MemoryConnection()
	db.AutoMigrate(&models.Webhook{})

	webhookRepo := repository.NewWebhookRepository(db)

	t.Run("correctly retrieves all records as json", func(t *testing.T) {
		webhookHandler := NewWebhookHandler(*webhookRepo)

		webhookRepo.Store(&models.Webhook{Name: " some name"})
		webhookRepo.Store(&models.Webhook{Name: "some other name"})
		webhooks, _ := webhookRepo.All()
		expectedLength := len(webhooks)

		req, err := http.NewRequest("GET", "/webhooks", nil)

		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(webhookHandler.Index)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Expected status %v, but got %v", http.StatusOK, status)
		}

		var response indexResponse
		err = json.NewDecoder(rr.Body).Decode(&response)
		if err != nil {
			t.Errorf("Failed to decode JSON response: %v", err)
		}

		if response.Webhooks == nil {
			t.Error("Expected 'webhooks' key in response JSON, but it was missing")
		}

		if len(response.Webhooks) != expectedLength {
			t.Errorf("Expected 'webhooks' to have a length of %d, but got %d", expectedLength, len(response.Webhooks))
		}
	})
}
