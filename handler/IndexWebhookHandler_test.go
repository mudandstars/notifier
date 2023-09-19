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

type indexResponse struct {
	Webhooks []string `json:"webhooks"`
}

func TestIndexWebhookHandler(t *testing.T) {
	webhookRepo := repository.NewWebhookRepository(database.MemoryConnection())
	webhookRepo.DB.AutoMigrate(&models.Webhook{})

	t.Run("correctly retrieves all records as json", func(t *testing.T) {
		indexWebhookHandler := &IndexWebhookHandler{
			Repo: *webhookRepo,
		}

		webhookRepo.Store(struct{ Name string }{Name: " some name"})
		webhookRepo.Store(struct{ Name string }{Name: "some other name"})
		expectedLength := 2

		req, err := http.NewRequest("GET", "/webhooks", nil)

		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(indexWebhookHandler.ServeHTTP)

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
