package database

import (
	"testing"
)

func TestConnection(t *testing.T) {
	t.Run("sqlite file connection", func(t *testing.T) {
		db, error := FileConnection().DB()

		if error != nil {
			t.Fatal(error)
		}

		if error = db.Ping(); error != nil {
			t.Fatal(error)
		}
	})

	t.Run("sqlite memory connection", func(t *testing.T) {
		db, error := MemoryConnection().DB()

		if error != nil {
			t.Fatal(error)
		}

		if error = db.Ping(); error != nil {
			t.Fatal(error)
		}
	})
}
