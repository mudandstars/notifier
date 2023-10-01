package database

import (
	"log"
	"os"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func FileConnection() *gorm.DB {
	dbPath, error := databasePath()

	if error != nil {
		log.Fatalf("failed to get database path: %v", error)
	}

	db, error := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if error != nil {
		log.Fatalf("failed to connect database: %v", error)
	}

	return db
}

func MemoryConnection() *gorm.DB {
	db, error := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if error != nil {
		panic("failed to connect database")
	}

	return db
}

func databasePath() (string, error) {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		return "", err
	}

	appSupportDir := filepath.Join(homeDir, "Library", "Application Support", "Notifier")
	err = os.MkdirAll(appSupportDir, 0755)

	if err != nil {
		return "", err
	}

	return filepath.Join(appSupportDir, "gorm.db"), nil
}
