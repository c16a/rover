package main

import (
	"context"
	"fmt"
	"os/exec"
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
	version, err := getJreVersion(handler.config.JreHome)
	if err != nil {
		return nil, err
	}
	return &schemas.InfoResponse{
		Name:    "jre",
		Version: version,
	}, nil
}

func getJreVersion(jrePath string) (string, error) {
	cmd := exec.Command(fmt.Sprintf("%s/bin/java", jrePath), "-version")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
