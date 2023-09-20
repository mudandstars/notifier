package main

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
	webhookHandler := (&handler.WebhookHandler{
		Repo: *webhookRepository,
	})

	if r.Method == http.MethodPost {
		webhookHandler.Store(w, r)
	}

	if r.Method == http.MethodGet {
		webhookHandler.Index(w, r)
	}

	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}