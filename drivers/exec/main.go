package main

import (
	"flag"
	"google.golang.org/grpc"
	"net"
	"rover/drivers/schemas"
	"rover/utils"
)

var configFileName string

func init() {
	flag.StringVar(&configFileName, "config", "config.json", "configuration file")
}

func main() {
	flag.Parse()

	c, err := utils.ParseJsonFile[Config](configFileName)
	if err != nil {
		panic(err)
	}

	listener, err := net.Listen("unix", c.SocketPath)
	if err != nil {
		panic(err)
	}
	var opts []grpc.ServerOption

	handler := NewHandler()

	grpcServer := grpc.NewServer(opts...)
	schemas.RegisterDriverServer(grpcServer, handler)

	grpcServer.Serve(listener)
}
