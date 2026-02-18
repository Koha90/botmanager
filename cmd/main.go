package main

import (
	"log"

	"github.com/joho/godotenv"

	"botmanager/internal/app"
	"botmanager/internal/config"
)

func main() {
	_ = godotenv.Load()

	cfg := config.MustLoad()

	app := app.NewApp(cfg)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
