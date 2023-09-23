package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/mudandstars/notifier/database"
	"github.com/mudandstars/notifier/handler"
	"github.com/mudandstars/notifier/repository"
	"github.com/rs/cors"
)

func RunLocalServer(ctx context.Context) error {
	http.HandleFunc("/webhooks", projectsRouter)
	http.HandleFunc("/webhooks/", projectsRouter)

	address := fmt.Sprintf(":%s", os.Getenv("PORT"))

	c := cors.AllowAll()

    handler := c.Handler(http.DefaultServeMux)

	log.Printf("Running local server on port %s", os.Getenv("PORT"))
	return http.ListenAndServe(address, handler)
}

func projectsRouter(w http.ResponseWriter, r *http.Request) {
	webhookHandler := handler.NewWebhookHandler(
		*repository.NewWebhookRepository(database.FileConnection()),
	)

	if r.Method == http.MethodPost {
		webhookHandler.Store(w, r)
	}

	if r.Method == http.MethodGet {
		log.Print("Received GET..")
		webhookHandler.Index(w, r)
		return
	}

	if r.Method == http.MethodDelete {
		webhookHandler.Delete(w, r)
		return
	}

	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
