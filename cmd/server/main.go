package main

import (
	"log/slog"
	"net/http"
	"os"
)

type config struct {
	addr      string
	staticDir string
}

func createLogger() *slog.Logger {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}))

	return logger

}

func main() {

	logger := createLogger()

	logger.Info("Starting server")

	port := os.Getenv("PORT")
	if port == "" {
		// TODO: Validate that port starts with :
		port = ":8080"
	}

	mux := routes()

	err := http.ListenAndServe(port, mux)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	logger.Info("Stopping server...")
}
