package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/mudandstars/notifier/models"
	"github.com/mudandstars/notifier/repository"
)

type ConfigHandler struct {
	repo repository.UserConfigRepository
}

func NewConfigHandler(repo repository.UserConfigRepository) ConfigHandler {
	return ConfigHandler{
		repo: repo,
	}
}

func (handler *ConfigHandler) Store(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		NgrokAuthToken string `json:"ngrokAuthToken"`
		NgrokPublicUrl string `json:"ngrokPublicUrl"`
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	userConfig, err := handler.repo.Get()
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if userConfig != (models.Config{}) {
		http.Error(w, "Cannot store a second user config", http.StatusNotAcceptable)
		return
	}

	if strings.Trim(requestBody.NgrokAuthToken, " ") == "" || strings.Trim(requestBody.NgrokPublicUrl, " ") == "" {
		http.Error(w, "Entries cannot be empty", http.StatusUnprocessableEntity)
		return
	}

	handler.repo.Store(&models.Config{
		NgrokAuthToken: strings.Trim(requestBody.NgrokAuthToken, " "),
		NgrokPublicUrl: strings.Trim(requestBody.NgrokPublicUrl, " "),
	})
}
