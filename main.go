package main

import (
	"github.com/mudandstars/notifier/database"
	"github.com/mudandstars/notifier/models"
	"github.com/mudandstars/notifier/server"
	"github.com/mudandstars/notifier/utils"

	"context"
	"log"
)

func main() {
	utils.LoadEnvironment()

	database.FileConnection().AutoMigrate(&models.Webhook{})

	// if err := server.RunNgrokServer(context.Background()); err != nil {
	// 	log.Fatal(err)
	// }

	if err := server.RunLocalServer(context.Background()); err != nil {
		log.Fatal(err)
	}
}
