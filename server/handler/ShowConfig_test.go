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

func TestShowConfig(t *testing.T) {
	db := database.MemoryConnection()
	db.AutoMigrate(&models.Config{})

	configRepo := repository.NewConfigRepository(db)
		configHandler := NewConfigHandler(*configRepo)

	t.Run("correctly retrieves the record as json", func(t *testing.T) {
		storedConfig := models.Config{NgrokAuthToken: "s12n309asdf", NgrokPublicUrl: "123asf.free-app.com"}
		configRepo.Upsert(&storedConfig)
		req, err := http.NewRequest("GET", "/config", nil)

		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(configHandler.Show)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Expected status %v, but got %v", http.StatusOK, status)
		}

		var response configBody
		err = json.NewDecoder(rr.Body).Decode(&response)
		if err != nil {
			t.Errorf("Failed to decode JSON response: %v", err)
		}

		if response.NgrokAuthToken != storedConfig.NgrokAuthToken {
			t.Error("Auth token is incorrect")
		}

		if response.NgrokPublicUrl != storedConfig.NgrokPublicUrl {
			t.Error("Public Url is incorrect")
		}
	})
}
