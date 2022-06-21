package main

import (
	"authService/internal/app"
	"authService/internal/config"
)

func main() {
	cfg := config.GetConfig()
	mainApp := app.NewApp(cfg)
	mainApp.Start()
}
