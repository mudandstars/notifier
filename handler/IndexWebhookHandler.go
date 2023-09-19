package handler

import (
	"log"
	"net/http"

	"github.com/mudandstars/notifier/repository"
	"github.com/mudandstars/notifier/utils"
)

type IndexWebhookHandler struct {
	Repo repository.WebhookRepository
}

func (handler *IndexWebhookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
