package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"

	"github.com/fedya-eremin/lo-trials/api"
	"github.com/fedya-eremin/lo-trials/logger"
	task_repo "github.com/fedya-eremin/lo-trials/repo/task"
	task_svc "github.com/fedya-eremin/lo-trials/service/task"
)

func main() {
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)

	l := logger.NewAsyncLogger(os.Stdout, 1000)
	defer l.Handler().(*logger.AsyncHandler).Close()
	slog.SetDefault(l)

	repo := task_repo.NewTaskRepo()
	service := task_svc.NewTaskService(repo)
	server := api.NewServer(service)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /tasks", server.GetTasksByStatus)
	mux.HandleFunc("GET /tasks/{id}", server.GetTaskById)
	mux.HandleFunc("POST /tasks", server.AddTask)
	handler := logger.AsyncLoggingMiddleware(mux)

	srv := &http.Server{
		Addr:    ":8001",
		Handler: handler,
	}

	serverErr := make(chan error, 1)
	go func() {
		slog.Info("Starting server", "addr", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			serverErr <- err
		}
	}()

	select {
	case err := <-serverErr:
		slog.Error("Server error", "error", err)
	case <-shutdown:
		slog.Info("Shutdown signal received")
	}

	slog.Info("Shutting down server...")
	if err := srv.Shutdown(context.Background()); err != nil {
		slog.Error("Server shutdown error", "error", err)
	}
	slog.Info("Server stopped")
}
