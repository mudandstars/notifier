package repository

import (
	"testing"

	"github.com/mudandstars/notifier/database"
	"github.com/mudandstars/notifier/models"
)

func TestUserConfigRepository(t *testing.T) {
	db := database.MemoryConnection()
	db.AutoMigrate(&models.Config{})

	userConfigRepo := NewConfigRepository(db)

	t.Run("correctly stores the userConfig", func(t *testing.T) {
		userConfig := models.Config{NgrokAuthToken: "asojf012n12", NgrokPublicUrl: "asdf.asdf.free-app.com"}

		error := userConfigRepo.Store(&userConfig)
		if error != nil {
			t.Fatal(error)
		}
	})

	t.Run("correctly fetches the user config", func(t *testing.T) {
		userConfig := models.Config{NgrokAuthToken: "asojf012asdfn12", NgrokPublicUrl: "asdf.as123df.free-app.com"}
		userConfigRepo.Store(&userConfig)

		if _, error := userConfigRepo.Get(); error != nil {
			t.Fatal(error)
		}
	})

	t.Run("correctly deletes the userConfig", func(t *testing.T) {
		userConfig := models.Config{NgrokAuthToken: "asojf012asdfn12", NgrokPublicUrl: "asdf.as123df.free-app.com"}
		userConfigRepo.Store(&userConfig)

		error := userConfigRepo.Delete()
		if error != nil {
			t.Fatal(error)
		}
	})
}
