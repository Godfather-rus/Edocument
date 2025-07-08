package app

import (
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/Godfather-rus/Edocument/internal/edocument/handlers"
	"github.com/Godfather-rus/Edocument/internal/edocument/repository"
)

type App struct {
	dbClient   *mongo.Client
	repository *repository.Repository

	router chi.Router
	http   *http.Server

	closers  []func() error
	handlers *handlers.Handlers
	closeCh  chan os.Signal
}

func NewApp() (*App, error) {
	var err error

	app := &App{}
	// - db
	err = app.initDBConn()
	if err != nil {
		return nil, fmt.Errorf("failed to init db connection: %w", err)
	}
	// - http server
	err = app.initServer()
	if err != nil {
		return nil, fmt.Errorf("failed to init http server: %w", err)
	}
	// - graceful shutdown
	err = app.initGracefulShutdown()
	if err != nil {
		return nil, fmt.Errorf("failed to init graceful shutdown: %w", err)
	}

	return app, nil
}

func (a *App) Run() error {
	go func() {
		err := a.http.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("http server failed", err)
		}
	}()

	<-a.closeCh
	for i := len(a.closers) - 1; i >= 0; i-- {
		err := a.closers[i]()
		if err != nil {
			log.Fatal("failed to close resource", i, err)
		}
	}

	return nil
}

func (a *App) initGracefulShutdown() error {
	a.closeCh = make(chan os.Signal, 1)
	signal.Notify(a.closeCh, os.Interrupt)

	return nil
}
