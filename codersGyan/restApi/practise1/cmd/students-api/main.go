package main

import (
	"context"
	"crudoperation/internal/config"
	"crudoperation/internal/http/handlers/student"
	"crudoperation/internal/storage/sqlite"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	// load config
	cfg := config.MustLoad()

	// database setup

	_, err := sqlite.New(cfg)

	if err != nil {
		log.Fatal(err)
	}

	slog.Info("storage initialize", slog.String("env", cfg.Env))

	// setup router
	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New())

	// setup server

	fmt.Println("testing data")
	// fmt.Println(cfg.HTTPServer.Addr)

	server := http.Server{
		Addr: cfg.HTTPServer.Addr,
		// Addr:    "localhost:8082",
		Handler: router,
	}

	slog.Info("server started", slog.String("address", cfg.Addr))
	fmt.Printf("server started %s", cfg.HTTPServer.Addr)

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("failed to start server")
		}
	}()

	<-done

	slog.Info("Shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = server.Shutdown(ctx)
	if err != nil {
		slog.Error("Failed to shutdown server", slog.String("error,", err.Error()))
	}
	slog.Info("server shutdown successfully")

}
