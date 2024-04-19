package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"

	"github.com/uptrace/bunrouter"
	"net/http"
	"rover/agent/handlers"
	"time"
)

func main() {
	logger, err := setupLogger()
	if err != nil {
		panic(err)
	}

	registry := prometheus.NewRegistry()
	setupMetrics(registry)
	setupRoutes(logger, registry, 8080)
}

func setupMetrics(registry *prometheus.Registry) {
	// Default collectors
	registry.MustRegister(collectors.NewGoCollector())
	registry.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
}

func setupLogger() (*zap.Logger, error) {
	return zap.NewProduction()
}

func setupRoutes(logger *zap.Logger, registry *prometheus.Registry, port int) {
	router := bunrouter.New()
	router.Compat().Handle("GET", "/metrics", func(writer http.ResponseWriter, request *http.Request) {
		promhttp.HandlerFor(registry, promhttp.HandlerOpts{}).ServeHTTP(writer, request)
	})

	// Management routes
	router.Handle(http.MethodGet, "/info", handlers.Info)
	srv := &http.Server{
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		Addr:         fmt.Sprintf(":%d", port),
	}
	logger.Info("Starting server")
	logger.Error("failed to start server", zap.Error(srv.ListenAndServe()))
}
