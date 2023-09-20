package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func FileConnection() *gorm.DB {
	db, error := gorm.Open(sqlite.Open("./gorm.db"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if error != nil {
		panic("failed to connect database")
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
