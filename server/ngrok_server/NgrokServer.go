package ngrok_server

import (
	"sync"
	"time"

	"github.com/mudandstars/notifier/models"
	"github.com/mudandstars/notifier/repository"
	"github.com/mudandstars/notifier/utils"
	"gorm.io/gorm"

	"context"
	"log"
	"net/http"

	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
)

var server *http.Server
var db *gorm.DB

func RunNgrokServer(newDB *gorm.DB) {
	db = newDB
	configRepo := repository.NewConfigRepository(db)
	config, error := configRepo.Get()

	if error == nil && config != (models.Config{}) {
		go func() {
			log.Fatal(spinUpServer(config))
		}()
	}

	for {
		newConfig, err := configRepo.Get()

		if err != nil {
			log.Printf("Error getting config: %v", err)
		}

		if config.NgrokAuthToken != newConfig.NgrokAuthToken || config.NgrokPublicUrl != newConfig.NgrokPublicUrl {
			log.Println("Config changed, shutting down the server...")
			shutdownServer()
			config = newConfig

			go func() {
				log.Fatal(spinUpServer(config))
			}()
		}

		time.Sleep(5 * time.Second)
	}
}

func shutdownServer() {
	if server != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Printf("Error shutting down the server: %v", err)
		} else {
			log.Println("Server shutdown successfully")
		}
	}
}

func spinUpServer(userConfig models.Config) error {
	serverCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tunnel, err := ngrok.Listen(serverCtx,
		config.HTTPEndpoint(
			config.WithDomain(userConfig.NgrokPublicUrl),
		),
		ngrok.WithAuthtoken(userConfig.NgrokAuthToken),
	)

	if err != nil {
		return err
	}

	router := http.NewServeMux()

	router.HandleFunc("/notifier", notificationHandler)

	log.Println("tunnel created:", tunnel.URL())

	server = &http.Server{Handler: router}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		if err := server.Serve(tunnel); err != nil {
			if err != http.ErrServerClosed {
				log.Printf("HTTP server error: %v", err)
			}

			if err := tunnel.Close(); err != nil {
				log.Printf("ngrok tunnel close error: %v", err)
			} else {
				log.Print("ngrok tunnel closed successfully")
			}
		}
	}()

	<-serverCtx.Done()
	wg.Wait()

	return nil
}

func notificationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	projectName := r.URL.Query()["name"][0]

	webhookRepo := repository.NewWebhookRepository(db)

	if webhookRepo.Exists(projectName) {
		utils.Notify("Project "+projectName, "Triggered URL")
	}
}
