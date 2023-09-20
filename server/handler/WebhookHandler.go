package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mudandstars/notifier/repository"
	"github.com/mudandstars/notifier/utils"
)

type WebhookHandler struct {
	Repo repository.WebhookRepository
}

func (handler *WebhookHandler) Index(w http.ResponseWriter, r *http.Request) {
	allWebhooks, error := handler.Repo.All()

	if error != nil {
		log.Fatal(error)
	}

	var names []string
	for _, webhook := range allWebhooks {
		names = append(names, webhook.Name)
	}

	responseObject := map[string][]string{
		"webhooks": names,
	}

	utils.WriteJson(w, responseObject)

	log.Println("something came to get data")
}

func (handler *WebhookHandler) Store(w http.ResponseWriter, r *http.Request) {
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
