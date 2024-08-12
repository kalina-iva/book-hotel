package server

import (
	"context"
	"errors"
	"net/http"

	"book_hotel/internal/pkg/logger"
)

var server http.Server

func NewServer(addr string, r http.Handler) error {
	logger.LogInfo("Server listening on localhost:8080")
	server = http.Server{Addr: addr, Handler: r}
	err := server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		logger.LogInfo("Server closed")
		return nil
	}
	return err
}

func Shutdown(ctx context.Context) {
	logger.LogInfo("Server is destroying...")
	if err := server.Shutdown(ctx); err != nil {
		logger.LogErrorf("cannot shutdown server: %s", err)
	}
	logger.LogInfo("Server was destroyed successfully")
}
