package api

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/castrintt/econ-api/internal/application/products"
	"github.com/castrintt/econ-api/internal/infrastructure/db"
	"github.com/castrintt/econ-api/internal/shared"
	"github.com/go-chi/chi/v5"
)

func (app *application) mount() http.Handler {
	router := chi.NewRouter()
	applymiddlewares(router)
	router.Get("/health", func(response http.ResponseWriter, request *http.Request) {
		shared.WriteAsJSON(response, http.StatusOK, "Healthy! 🚀")
	})

	productService := products.NewService()
	productHandler := products.NewHandler(productService)

	router.Get("/products", productHandler.GetProducts)

	return router
}

func (app *application) run(handler http.Handler) error {
	port := app.configuration.Port
	server := &http.Server{
		Addr:         port,
		Handler:      handler,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Println("Starting server on port", port)
	return server.ListenAndServe()
}

func InitializeApplication() error {
	config := db.DatabaseConfigurationFromEnv()
	if err := db.InitializeDatabase(context.Background(), config.ConnectionString); err != nil {
		return nil
	}
	app := &application{configuration: config}
	return app.run(app.mount())
}

type application struct {
	configuration db.DatabaseConfiguration
	//logger
	//db driver
}
