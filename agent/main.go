package main

import (
	"flag"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/uptrace/bunrouter"
	"go.uber.org/zap"
	"net/http"
	"rover/agent/handlers"
	"rover/utils"
	"time"
)

var configFileName string

func init() {
	flag.StringVar(&configFileName, "config", "config.json", "configuration file")
}

func main() {
	flag.Parse()

	logger, err := setupLogger()
	if err != nil {
		panic(err)
	}

	config, err := utils.ParseJsonFile[Config](configFileName)
	if err != nil {
		logger.Fatal("failed to read config file", zap.Error(err))
	}

	registry := prometheus.NewRegistry()
	setupMetrics(registry)

	go DialDrivers(config, logger)

	startRouter(logger, registry, config)
}

func setupMetrics(registry *prometheus.Registry) {
	// Default collectors
	registry.MustRegister(collectors.NewGoCollector())
	registry.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
}

func setupLogger() (*zap.Logger, error) {
	return zap.NewProduction()
}

func startRouter(logger *zap.Logger, registry *prometheus.Registry, config *Config) {
	router := bunrouter.New()
	router.Compat().Handle(http.MethodGet, "/metrics", func(writer http.ResponseWriter, request *http.Request) {
		promhttp.HandlerFor(registry, promhttp.HandlerOpts{}).ServeHTTP(writer, request)
	})

	// Management routes
	router.Handle(http.MethodGet, "/info", handlers.Info)
	srv := &http.Server{
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		Addr:         config.BindAddr,
	}
	logger.Info("Starting server")
	logger.Error("failed to start server", zap.Error(srv.ListenAndServe()))
}
