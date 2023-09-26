package handler

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/mudandstars/notifier/database"
	"github.com/mudandstars/notifier/models"
	"github.com/mudandstars/notifier/repository"
)

func TestDeleteWebhook(t *testing.T) {
	db := database.MemoryConnection()
	db.AutoMigrate(&models.Webhook{})

	webhookRepo := repository.NewWebhookRepository(db)

	webhook := models.Webhook{Name: "test"}
	webhookRepo.Store(&webhook)

	webhookHandler := NewWebhookHandler(*webhookRepo)

	t.Run("correctly deletes the record", func(t *testing.T) {
		req, err := http.NewRequest("DELETE", "/webhooks/"+strconv.FormatUint(uint64(webhook.ID), 10), nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(webhookHandler.Delete)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Expected status %v, but got %v", http.StatusOK, status)
		}
	})
}
