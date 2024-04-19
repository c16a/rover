package main

type Config struct {
	Drivers []*DriverConfig
}

type DriverConfig struct {
	Name       string
	SocketPath string
}
