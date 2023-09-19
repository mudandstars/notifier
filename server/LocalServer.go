package server

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/mudandstars/notifier/database"
	"github.com/mudandstars/notifier/handler"
	"github.com/mudandstars/notifier/repository"
)

func RunLocalServer(ctx context.Context) error {
	http.HandleFunc("/webhooks", projectsRouter)

	address := fmt.Sprintf(":%s", os.Getenv("PORT"))

	return http.ListenAndServe(address, nil)
}

func projectsRouter(w http.ResponseWriter, r *http.Request) {
	webhookRepository := repository.NewWebhookRepository(database.FileConnection())

	if r.Method == http.MethodPost {
		(&handler.StoreWebhookHandler{
			Repo: *webhookRepository,
		}).ServeHTTP(w, r)
	}

	if r.Method == http.MethodGet {
		(&handler.IndexWebhookHandler{
			Repo: *webhookRepository,
		}).ServeHTTP(w, r)
	}

	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
