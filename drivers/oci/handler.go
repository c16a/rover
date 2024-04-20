package main

import (
	"context"
	"rover/drivers/schemas"
)

type Handler struct {
	config *Config
}

func NewHandler(config *Config) *Handler {
	return &Handler{
		config: config,
	}
}

func (handler *Handler) GetHealth(ctx context.Context, req *schemas.HealthRequest) (*schemas.HealthResponse, error) {
	return &schemas.HealthResponse{Status: true}, nil
}

func (handler *Handler) GetInfo(ctx context.Context, req *schemas.InfoRequest) (*schemas.InfoResponse, error) {
	return &schemas.InfoResponse{
		Name:    "oci",
		Version: "latest",
	}, nil
}
