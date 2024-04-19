package main

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"rover/drivers/schemas"
	"time"
)

func DialDrivers(config *Config, logger *zap.Logger) {
	for _, driver := range config.Drivers {
		err := dialDriver(driver, logger)
		if err != nil {
			logger.Error("failed to dial driver", zap.Error(err), zap.String("name", driver.Name))
		}
	}
}

func dialDriver(driverConfig *DriverConfig, logger *zap.Logger) error {
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
				response, err := dialDriverLoop(driverConfig, logger)
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

func dialDriverLoop(driverConfig *DriverConfig, logger *zap.Logger) (*schemas.HealthResponse, error) {
	dialer := func(ctx context.Context, addr string) (net.Conn, error) {
		var d net.Dialer
		return d.DialContext(ctx, "unix", addr)
	}
	options := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithContextDialer(dialer),
	}

	ctx, cancelFn := context.WithTimeout(context.Background(), time.Duration(driverConfig.DialTimeoutSeconds)*time.Second)
	defer cancelFn()
	conn, err := grpc.DialContext(ctx, driverConfig.SocketPath, options...)
	if err != nil {
		return nil, err
	}

	client := schemas.NewDriverClient(conn)
	response, err := client.GetHealth(context.Background(), &schemas.HealthRequest{})
	if err != nil {
		return nil, err
	}

	return response, nil
}
