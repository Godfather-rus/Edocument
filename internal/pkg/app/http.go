package app

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Godfather-rus/Edocument/internal/edocument/handlers"
	"github.com/Godfather-rus/Edocument/internal/edocument/repository"
)

func (a *App) initServer() error {

	a.repository = repository.NewRepository(a.dbClient)

	a.handlers = handlers.NewHandlers(a.repository)

	log.Println("Server running")

	r := chi.NewRouter()

	r.Post("/api/docs", a.handlers.CreateEdoc)

	r.Get("/api/docs", a.handlers.GetEdocsList)
	r.Get("/api/docs/{id}", a.handlers.GetEdoc)

	a.router = r

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("HTTP_PORT")),
		Handler: r,
	}

	a.http = server
	a.closers = append(a.closers, func() error {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		return a.http.Shutdown(ctx)
	})

	return nil
}
