package main

import (
	"log"
	"net/http"
	"time"

	repo "github.com/Abhishekkapoor98/Ecom_GO_API/internal/adapter/postgresql/sqlc"
	"github.com/Abhishekkapoor98/Ecom_GO_API/internal/products"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

func (app *application) mount() http.Handler {

	// Create a new chi router
	r := chi.NewRouter()

	// Add some basic middleware
	r.Use(middleware.RealIP)    //Important for rate limiting, analitics and tracing
	r.Use(middleware.RequestID) // Important for rate limiting
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer) //Recover from crashes

	//Set timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	// Define your routes here
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("All good."))
	})

	productServce := products.NewService(repo.New(app.db))
	productHandler := products.NewHandler(productServce) // Pass a real service implementation here
	r.Get("/products", productHandler.ListProducts)

	return r
}

func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Starting server on %v", app.config.addr)

	return srv.ListenAndServe()
}

type application struct {
	config config

	//logger
	db *pgx.Conn
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dsn string
}
