package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/mudandstars/notifier/repository"
	"github.com/mudandstars/notifier/utils"
)

type WebhookHandler struct {
	Repo repository.WebhookRepository
}

type indexResponse struct {
	Webhooks []indexWebhook `json:"webhooks"`
}

type indexWebhook struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func (handler *WebhookHandler) Index(w http.ResponseWriter, r *http.Request) {
	allWebhooks, error := handler.Repo.All()

	if error != nil {
		log.Fatal(error)
	}

	var webhooksBody []indexWebhook

	for _, webhook := range allWebhooks {
		webhooksBody = append(webhooksBody, indexWebhook{
			Name: webhook.Name,
			Url:  os.Getenv("NGROK_PUBLIC_URL") + "?name=" + webhook.Name,
		})
	}

	responseObject := indexResponse{
		Webhooks: webhooksBody,
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
