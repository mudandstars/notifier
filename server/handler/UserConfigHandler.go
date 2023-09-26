package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/mudandstars/notifier/models"
	"github.com/mudandstars/notifier/repository"
)

type UserConfigHandler struct {
	repo repository.UserConfigRepository
}

func NewUserConfigHandler(repo repository.UserConfigRepository) UserConfigHandler {
	return UserConfigHandler{
		repo: repo,
	}
}

func (handler *UserConfigHandler) Store(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		NgrokAuthToken string `json:"ngrokAuthToken"`
		NgrokPublicUrl string `json:"ngrokPublicUrl"`
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	userConfig, _ := handler.repo.Get()

	if userConfig != (models.UserConfig{}) {
		http.Error(w, "Cannot store a second user config", http.StatusNotAcceptable)
		return
	}

	if strings.Trim(requestBody.NgrokAuthToken, " ") == "" || strings.Trim(requestBody.NgrokPublicUrl, " ") == "" {
		http.Error(w, "Entries cannot be empty", http.StatusUnprocessableEntity)
		return
	}

	handler.repo.Store(&models.UserConfig{
		NgrokAuthToken: strings.Trim(requestBody.NgrokAuthToken, " "),
		NgrokPublicUrl: strings.Trim(requestBody.NgrokPublicUrl, " "),
	})
}
