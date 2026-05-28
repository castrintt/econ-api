package api

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func applymiddlewares(router *chi.Mux) {
	router.Use(middleware.RequestID)                 // important for rate limiting
	router.Use(middleware.RealIP)                    // important for rate limiting
	router.Use(middleware.Logger)                    // important for logging
	router.Use(middleware.Recoverer)                 // important for panic recovery
	router.Use(middleware.Timeout(60 * time.Second)) // important for timeout
}
