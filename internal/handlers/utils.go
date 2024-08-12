package handlers

import (
	"net/http"

	"book_hotel/internal/pkg/logger"
)

func setError(w http.ResponseWriter, message string, code int) {
	http.Error(w, message, code)
	logger.LogErrorf(message)
}
