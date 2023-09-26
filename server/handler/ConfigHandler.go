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

type ConfigHandler struct {
	repo repository.ConfigRepository
}

func NewConfigHandler(repo repository.ConfigRepository) ConfigHandler {
	return ConfigHandler{
		repo: repo,
	}
}

type configBody struct {
	NgrokAuthToken string `json:"ngrokAuthToken"`
	NgrokPublicUrl string `json:"ngrokPublicUrl"`
	ID             uint   `json:"id"`
}

func (handler *ConfigHandler) Upsert(w http.ResponseWriter, r *http.Request) {
	var requestBody configBody

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if strings.Trim(requestBody.NgrokAuthToken, " ") == "" || strings.Trim(requestBody.NgrokPublicUrl, " ") == "" {
		http.Error(w, "Entries cannot be empty", http.StatusUnprocessableEntity)
		return
	}

	handler.repo.Upsert(&models.Config{
		NgrokAuthToken: strings.Trim(requestBody.NgrokAuthToken, " "),
		NgrokPublicUrl: strings.Trim(requestBody.NgrokPublicUrl, " "),
	})
}

func (handler *ConfigHandler) Show(w http.ResponseWriter, r *http.Request) {
	config, err := handler.repo.Get()

	if err != nil {
		http.Error(w, "Unkown error", http.StatusInternalServerError)
		log.Fatal(err)
	}

	utils.WriteJson(w, config)
}
