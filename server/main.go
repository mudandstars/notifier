package main

import (
	"github.com/mudandstars/notifier/database"
	"github.com/mudandstars/notifier/local_server"
	"github.com/mudandstars/notifier/models"
	"github.com/mudandstars/notifier/ngrok_server"

	"context"
	"log"
)

func main() {
	db := database.FileConnection()
	db.AutoMigrate(&models.Webhook{})
	db.AutoMigrate(&models.Config{})

	go func() {
		ngrok_server.RunNgrokServer(db)
	}()

	log.Fatal(local_server.RunLocalServer(context.Background()))
}
