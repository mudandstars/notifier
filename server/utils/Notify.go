package utils

import (
	"github.com/gen2brain/beeep"
)

func Notify (title string, description string) {
	err := beeep.Alert(title, description, "assets/warning.png")
	if err != nil {
		panic(err)
	}
}
