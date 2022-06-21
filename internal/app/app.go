package app

import (
	"authService/internal/config"
	"authService/internal/domain/auth"
	"authService/internal/domain/user"
	"authService/pkg/client/mongodb"
	"context"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type App struct {
	cfg        *config.Config
	router     *httprouter.Router
	httpServer *http.Server
	logger     *log.Logger
}

func NewApp(config *config.Config) App {
	router := httprouter.New()
	mongodb.InitClient(context.Background(), config.MongoDB.Host, config.MongoDB.Port, config.MongoDB.Database)

	userHandler := user.NewUserHandler()
	userHandler.Register(router)

	authHandler := auth.NewAuthHandler()
	authHandler.Register(router)

	return App{
		cfg:    config,
		router: router,
		logger: log.Default(),
	}
}

func (app *App) Start() {
	app.logger.Printf("server listening on PORT:%s", app.cfg.Listen.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", app.cfg.Listen.Port), app.router)
	if err != nil {
		panic(err)
	}
}
