package handler

import (
	"encoding/json"
	"net/http"

	"github.com/mudandstars/notifier/repository"
)

type StoreWebhookHandler struct {
	Repo repository.WebhookRepository
}

type storeWebhookRequest struct {
	Name string `json:"Name"`
}

func (handler *StoreWebhookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		Name string `json:"Name"`
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	handler.Repo.Store(repository.CreateWebhookBody{
		Name: requestBody.Name,
	})
}
