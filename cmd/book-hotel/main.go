package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"book_hotel/internal/core/order"
	"book_hotel/internal/handlers"
	"book_hotel/internal/pkg/logger"
	"book_hotel/internal/repository"
	"book_hotel/internal/server"
	"book_hotel/internal/storage"
)

func main() {
	logger.InitLogger()

	ctx := context.Background()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	// get from config
	addr := ":8080"

	db := storage.NewDB()

	orderRepo := repository.NewOrderRepo(db)
	hotelRepo := repository.NewRepo(db)

	orderService := order.NewService(&orderRepo, &hotelRepo)

	h := handlers.NewHandler(orderService)

	router := server.NewRouter(h)

	go func() {
		err := server.NewServer(addr, router)
		if err != nil {
			logger.LogFatal("cannot start server %s", err)
		}
	}()
	defer server.Shutdown(ctx)

	<-sig
}
