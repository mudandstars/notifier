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

func TestStoreConfig(t *testing.T) {
	db := database.MemoryConnection()
	db.AutoMigrate(&models.Config{})

	userConfigRepo := repository.NewConfigRepository(db)
	userConfigHandler := NewConfigHandler(*userConfigRepo)

	t.Run("correctly stores or updates the record", func(t *testing.T) {
		userConfigRepo.Delete()

		requestBody := models.Config{
			NgrokAuthToken: "somasd123j910",
			NgrokPublicUrl: "12asdf12.free-app.com",
		}
		rr := upsertConfigRequest(t, userConfigHandler, requestBody)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Expected status %v, but got %v", http.StatusOK, status)
		}

		userConfig, _ := userConfigHandler.repo.Get()

		if userConfig.NgrokAuthToken != requestBody.NgrokAuthToken || userConfig.NgrokPublicUrl != requestBody.NgrokPublicUrl {
			t.Fatal("Webhook was not created successfully")
		}
	})

	t.Run("values cannot be empty strings", func(t *testing.T) {
		userConfigRepo.Delete()
		requestBody := models.Config{
			NgrokAuthToken: "",
			NgrokPublicUrl: "12asdf12.free-app.com",
		}
		rr := upsertConfigRequest(t, userConfigHandler, requestBody)

		if status := rr.Code; status != http.StatusUnprocessableEntity {
			t.Errorf("Expected status %v, but got %v", http.StatusUnprocessableEntity, status)
		}

		userConfigRepo.Delete()
		requestBody = models.Config{
			NgrokAuthToken: "123pjk12o31",
			NgrokPublicUrl: "",
		}
		rr = upsertConfigRequest(t, userConfigHandler, requestBody)

		if status := rr.Code; status != http.StatusUnprocessableEntity {
			t.Errorf("Expected status %v, but got %v", http.StatusUnprocessableEntity, status)
		}
	})

	t.Run("values get trimmed", func(t *testing.T) {
		userConfigRepo.Delete()

		authToken := " somasd1a23j910 "
		publicUrl := "   12asdf12.free-app.com  "
		requestBody := models.Config{
			NgrokAuthToken: authToken,
			NgrokPublicUrl: publicUrl,
		}

		rr := upsertConfigRequest(t, userConfigHandler, requestBody)

		userConfig, _ := userConfigHandler.repo.Get()

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Expected status %v, but got %v", http.StatusOK, status)
		}

		if userConfig.NgrokAuthToken != strings.Trim(authToken, " ") || userConfig.NgrokPublicUrl != strings.Trim(publicUrl, " ") {
			t.Fatal("Values were not trimmed properly")
		}
	})
}

func upsertConfigRequest(t *testing.T, userConfigHandler ConfigHandler, requestBody models.Config) *httptest.ResponseRecorder {
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
