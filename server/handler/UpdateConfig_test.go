package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/mudandstars/notifier/database"
	"github.com/mudandstars/notifier/models"
	"github.com/mudandstars/notifier/repository"
)

func TestUpdateConfig(t *testing.T) {
	db := database.MemoryConnection()
	db.AutoMigrate(&models.Config{})

	configRepo := repository.NewConfigRepository(db)
	configHandler := NewConfigHandler(*configRepo)
	config := models.Config{
		NgrokAuthToken: "1j029jasdf",
		NgrokPublicUrl: "asodsdn.free-app.com",
	}
	configRepo.Delete()
	configRepo.Upsert(&config)

	t.Run("correctly updates the config", func(t *testing.T) {
		requestBody := configBody{
			NgrokAuthToken: "somasd123j910",
			NgrokPublicUrl: "12asdf12.free-app.com",
			ID:             config.ID,
		}

		rr := updateConfigRequest(t, configHandler, requestBody)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Expected status %v, but got %v", http.StatusOK, status)
		}

		userConfig, _ := configHandler.repo.Get()

		if userConfig.NgrokAuthToken != requestBody.NgrokAuthToken || userConfig.NgrokPublicUrl != requestBody.NgrokPublicUrl {
			t.Fatal("Webhook was not updated successfully")
		}
	})

	t.Run("values cannot be empty strings", func(t *testing.T) {
		requestBody := configBody{
			NgrokAuthToken: "",
			NgrokPublicUrl: "12asdf12.free-app.com",
			ID:             config.ID,
		}
		rr := updateConfigRequest(t, configHandler, requestBody)

		if status := rr.Code; status != http.StatusUnprocessableEntity {
			t.Errorf("Expected status %v, but got %v", http.StatusUnprocessableEntity, status)
		}

		requestBody = configBody{
			NgrokAuthToken: "",
			NgrokPublicUrl: "12asdf12.free-app.com",
			ID:             config.ID,
		}

		rr = updateConfigRequest(t, configHandler, requestBody)

		if status := rr.Code; status != http.StatusUnprocessableEntity {
			t.Errorf("Expected status %v, but got %v", http.StatusUnprocessableEntity, status)
		}
	})

	t.Run("values get trimmed", func(t *testing.T) {
		authToken := " somasd1a23j910 "
		publicUrl := "   12asdf12.free-app.com  "

		requestBody := configBody{
			NgrokAuthToken: authToken,
			NgrokPublicUrl: publicUrl,
			ID:             config.ID,
		}

		rr := updateConfigRequest(t, configHandler, requestBody)

		userConfig, _ := configHandler.repo.Get()

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Expected status %v, but got %v", http.StatusOK, status)
		}

		if userConfig.NgrokAuthToken != strings.Trim(authToken, " ") || userConfig.NgrokPublicUrl != strings.Trim(publicUrl, " ") {
			t.Fatal("Values were not trimmed correctly")
		}
	})
}

func updateConfigRequest(t *testing.T, userConfigHandler ConfigHandler, requestBody configBody) *httptest.ResponseRecorder {
	body, _ := json.Marshal(requestBody)

	req, err := http.NewRequest("PUT", "/config", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userConfigHandler.Upsert)

	handler.ServeHTTP(rr, req)

	return rr
}
