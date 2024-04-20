package main

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"rover/drivers/schemas"
	"time"
)

func DialDrivers(config *Config, logger *zap.Logger) {
	for _, driver := range config.Drivers {
		go func(driver *DriverConfig) {
			err := dialDriver(driver, logger)
			if err != nil {
				logger.Error("failed to dial driver", zap.Error(err), zap.String("name", driver.Name))
			}
		}(driver)
	}
}

func dialDriver(driverConfig *DriverConfig, logger *zap.Logger) error {
	options := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.NewClient(driverConfig.SocketPath, options...)
	if err != nil {
		return err
	}

	client := schemas.NewDriverClient(conn)

	ctx, cancelFn := context.WithTimeout(context.Background(), time.Duration(driverConfig.DialTimeoutSeconds)*time.Second)
	defer cancelFn()
	infoResponse, err := client.GetInfo(ctx, &schemas.InfoRequest{})
	if err != nil {
		return err
	}

	logger.Info("registered driver", zap.String("name", infoResponse.Name), zap.String("version", infoResponse.Version))

	errorCh := make(chan error)
	defer close(errorCh)

	go func() {
		for {
			err := <-errorCh
			if err != nil {
				logger.Error("failed to dial driver", zap.Error(err), zap.String("name", driverConfig.Name))
			}
		}
	}()

	ticker := time.NewTicker(time.Duration(driverConfig.DialIntervalSeconds) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			go func() {
				response, err := dialDriverLoop(client)
				if err != nil {
					errorCh <- err
				} else {
					logger.Info("received driver status",
						zap.Bool("status", response.Status),
						zap.String("name", driverConfig.Name),
					)
				}
			}()
		}
	}
}

func dialDriverLoop(client schemas.DriverClient) (*schemas.HealthResponse, error) {
	response, err := client.GetHealth(context.Background(), &schemas.HealthRequest{})
	if err != nil {
		return nil, err
	}

	return response, nil
}
