package main

import (
	"github.com/mudandstars/notifier/database"
	"github.com/mudandstars/notifier/models"
	"github.com/mudandstars/notifier/utils"

	"context"
	"log"
)

func main() {
	utils.LoadEnvironment()

	database.FileConnection().AutoMigrate(&models.Webhook{})

	go func() {
		log.Fatal(RunLocalServer(context.Background()))
	}()

	log.Fatal(RunNgrokServer(context.Background()))
}
