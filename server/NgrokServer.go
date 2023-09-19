package server

import (
	"github.com/mudandstars/notifier/utils"

	"context"
	"log"
	"net/http"
	"os"

	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
)

func RunNgrokServer(ctx context.Context) error {
	tunnel, err := ngrok.Listen(ctx,
		config.HTTPEndpoint(
			config.WithDomain(os.Getenv("NGROK_PUBLIC_URL")),
		),
		ngrok.WithAuthtokenFromEnv(),
	)
	if err != nil {
		return err
	}

	router := http.NewServeMux()

	router.HandleFunc("/notifier", baseHandler)
	// router.HandleFunc("/anotherendpoint", anotherEndpointHandler)

	log.Println("tunnel created:", tunnel.URL())

	return http.Serve(tunnel, router)
}

func baseHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	utils.Notify("Hallo Felix", "Get rekt");
}
