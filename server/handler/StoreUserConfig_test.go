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

func TestStoreUserConfig(t *testing.T) {
	db := database.MemoryConnection()
	db.AutoMigrate(&models.UserConfig{})

	userConfigRepo := repository.NewUserConfigRepository(db)
	userConfigHandler := NewUserConfigHandler(*userConfigRepo)

	t.Run("correctly stores the record", func(t *testing.T) {
		userConfigRepo.Delete()
	
		requestBody := models.UserConfig{
			NgrokAuthToken: "somasd123j910",
			NgrokPublicUrl: "12asdf12.free-app.com",
		}
		rr := storeUserConfigRequest(t, userConfigHandler, requestBody)

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
		requestBody := models.UserConfig{
			NgrokAuthToken: "",
			NgrokPublicUrl: "12asdf12.free-app.com",
		}
		rr := storeUserConfigRequest(t, userConfigHandler, requestBody)

		if status := rr.Code; status != http.StatusUnprocessableEntity {
			t.Errorf("Expected status %v, but got %v", http.StatusUnprocessableEntity, status)
		}

		userConfigRepo.Delete()
		requestBody = models.UserConfig{
			NgrokAuthToken: "123pjk12o31",
			NgrokPublicUrl: "",
		}
		rr = storeUserConfigRequest(t, userConfigHandler, requestBody)

		if status := rr.Code; status != http.StatusUnprocessableEntity {
			t.Errorf("Expected status %v, but got %v", http.StatusUnprocessableEntity, status)
		}
	})

	t.Run("values get trimmed", func(t *testing.T) {
		userConfigRepo.Delete()

		authToken := " somasd1a23j910 "
		publicUrl := "   12asdf12.free-app.com  "
		requestBody := models.UserConfig{
			NgrokAuthToken: authToken,
			NgrokPublicUrl: publicUrl,
		}

		rr := storeUserConfigRequest(t, userConfigHandler, requestBody)

		userConfig, _ := userConfigHandler.repo.Get()

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Expected status %v, but got %v", http.StatusOK, status)
		}

		if userConfig.NgrokAuthToken != strings.Trim(authToken, " ") || userConfig.NgrokPublicUrl != strings.Trim(publicUrl, " ") {
			t.Fatal("Values were not trimmed properly")
		}
	})

	t.Run("cannot store a second config", func(t *testing.T) {
		userConfigRepo.Delete()
		requestBody := models.UserConfig{
			NgrokAuthToken: "asfj120-j19asdf",
			NgrokPublicUrl: "12asdf12.free-app.com",
		}
		storeUserConfigRequest(t, userConfigHandler, requestBody)
		rr := storeUserConfigRequest(t, userConfigHandler, requestBody)

		if status := rr.Code; status != http.StatusNotAcceptable {
			t.Errorf("Expected status %v, but got %v", http.StatusNotAcceptable, status)
		}

		userConfigRepo.Delete()
	})
}

func storeUserConfigRequest(t *testing.T, userConfigHandler UserConfigHandler, requestBody models.UserConfig) *httptest.ResponseRecorder {
	body, _ := json.Marshal(requestBody)

	req, err := http.NewRequest("POST", "/user-configs", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userConfigHandler.Store)

	handler.ServeHTTP(rr, req)

	return rr
}
