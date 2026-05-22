package api

import (
	"log"
	"net/http"
	"time"

	"github.com/castrintt/econ-api/cmd/env"
	"github.com/castrintt/econ-api/internal/application/products"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) mount() http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)                 // important for rate limiting
	router.Use(middleware.RealIP)                    // important for rate limiting
	router.Use(middleware.Logger)                    // important for logging
	router.Use(middleware.Recoverer)                 // important for panic recovery
	router.Use(middleware.Timeout(60 * time.Second)) // important for timeout

	router.Get("/health", func(response http.ResponseWriter, request *http.Request) {
		response.WriteHeader(http.StatusOK)
		response.Write([]byte("Healthy! 🚀"))
	})

	productService := products.NewService()
	productHandler := products.NewHandler(productService)

	router.Get("/products", productHandler.GetProducts)

	return router
}

func (app *application) run(handler http.Handler) error {
	server := &http.Server{
		Addr:         app.configuration.port,
		Handler:      handler,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Println("Starting server on port", app.configuration.port)

	return server.ListenAndServe()
}

func InitializeApplication() error {
	config := configuration{
		port: env.APPLICATION_PORT,
		database: databaseConfiguration{
			host:         env.DB_HOST,
			port:         env.DB_PORT,
			user:         env.DB_USER,
			password:     env.DB_PASSWORD,
			databaseName: env.DB_NAME,
		},
	}
	api := &application{
		configuration: config,
	}
	handler := api.mount()
	return api.run(handler)
}

type application struct {
	configuration configuration
	//logger
	//db driver
}

type configuration struct {
	port     string
	database databaseConfiguration
}

type databaseConfiguration struct {
	host         string
	port         string
	user         string
	password     string
	databaseName string
}
