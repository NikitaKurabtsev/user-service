package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/NikitaKurabtsev/user-service/internal/adapters/in/handlers"
	"github.com/NikitaKurabtsev/user-service/internal/adapters/out/repository"
	"github.com/NikitaKurabtsev/user-service/internal/core/services"
)

func main() {
	repo := repository.NewInMemoryRepository()
	service := services.NewUserService(repo)
	handler := handlers.NewUserHandler(service)

	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	server := http.Server{
		Addr:         ":8000",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	go func() {
		log.Println("starting server on :8000")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server failed %v\n", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown :%v", err)
	}

	log.Println("Server exited gracefully")
}
