package handler

import (
	"encoding/json"
	"fmt"
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
	db.AutoMigrate(&models.Config{})

	webhookRepo := repository.NewWebhookRepository(db)
	configRepo := repository.NewConfigRepository(db)

	url := "test-url.com"
	name := "some other webhook name"
	configRepo.Delete()
	configRepo.Upsert(&models.Config{
		NgrokPublicUrl: url,
		NgrokAuthToken: "soansd092h019",
	})
	webhookRepo.Store(&models.Webhook{Name: name})

	t.Run("correctly retrieves all records as json", func(t *testing.T) {
		webhookHandler := NewWebhookHandler(*webhookRepo)

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

	t.Run("correctly sends url content", func(t *testing.T) {
		webhookHandler := NewWebhookHandler(*webhookRepo)

		req, _ := http.NewRequest("GET", "/webhooks", nil)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(webhookHandler.Index)

		handler.ServeHTTP(rr, req)

		var response indexResponse
		json.NewDecoder(rr.Body).Decode(&response)

		if response.Webhooks[0].Name != name || response.Webhooks[0].Url != fmt.Sprintf("%s/notifier?name=%s", url, name) {
			t.Error("Webhook name or Url were incorrect")
		}
	})

}
