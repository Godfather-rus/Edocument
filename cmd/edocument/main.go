package main

import (
	"log"

	"github.com/Godfather-rus/Edocument/internal/pkg/app"
)

func main() {
	ap, err := app.NewApp()
	if err != nil {
		log.Fatal("failed to create app", "error", err)
		return
	}

	if err = ap.Run(); err != nil {
		log.Fatal("failed to run app", "error", err)
	}
}
