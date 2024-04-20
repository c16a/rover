package main

import (
	"flag"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"os"
	"rover/drivers/schemas"
	"rover/utils"
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

	c, err := utils.ParseJsonFile[Config](configFileName)
	if err != nil {
		panic(err)
	}

	err = os.Remove(c.SocketPath)

	listener, err := net.Listen("unix", c.SocketPath)
	if err != nil {
		panic(err)
	}
	var opts []grpc.ServerOption

	handler := NewHandler(c)

	grpcServer := grpc.NewServer(opts...)
	schemas.RegisterDriverServer(grpcServer, handler)

	logger.Info("starting driver", zap.String("socketPath", c.SocketPath), zap.String("name", c.Name))
	grpcServer.Serve(listener)
}

func setupLogger() (*zap.Logger, error) {
	return zap.NewProduction()
}
