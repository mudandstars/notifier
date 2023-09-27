package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/mudandstars/notifier/models"
	"github.com/mudandstars/notifier/repository"
	"github.com/mudandstars/notifier/utils"
)

type WebhookHandler struct {
	repo repository.WebhookRepository
}

func NewWebhookHandler(repo repository.WebhookRepository) WebhookHandler {
	return WebhookHandler{
		repo: repo,
	}
}

type indexResponse struct {
	Webhooks []indexWebhook `json:"webhooks"`
}

type indexWebhook struct {
	Name string `json:"name"`
	Url  string `json:"url"`
	Id   uint   `json:"id"`
}

func (handler *WebhookHandler) Index(w http.ResponseWriter, r *http.Request) {
	allWebhooks, error := handler.repo.All()

	configRepo := repository.NewConfigRepository(handler.repo.DB)
	config, err := configRepo.Get()

	var tunnelUrl string
	if config != (models.Config{}) && err == nil {
		tunnelUrl = config.NgrokPublicUrl
	} else {
		tunnelUrl = ""
	}

	if error != nil {
		log.Fatal(error)
	}

	var webhooksBody []indexWebhook

	for _, webhook := range allWebhooks {
		webhooksBody = append(webhooksBody, indexWebhook{
			Name: webhook.Name,
			Url:  tunnelUrl + "/notifier?name=" + webhook.Name,
			Id:   webhook.ID,
		})
	}

	responseObject := indexResponse{
		Webhooks: webhooksBody,
	}

	utils.WriteJson(w, responseObject)
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

	name := strings.Trim(requestBody.Name, " ")

	if name == "" {
		http.Error(w, "Name cannot be empty", http.StatusUnprocessableEntity)
		return
	}

	if handler.repo.Exists(name) {
		http.Error(w, "Entry with this name already exists", http.StatusNotAcceptable)
		return
	}

	handler.repo.Store(&models.Webhook{
		Name: requestBody.Name,
	})
}

func (handler *WebhookHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, error := utils.Path(r.URL.Path, "webhooks")

	if error != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	handler.repo.Delete(id)
}
