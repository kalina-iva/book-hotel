package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"book_hotel/internal/handlers"
)

func NewRouter(h handlers.Handler) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/orders", h.CreateOrderHandler)

	return r
}
