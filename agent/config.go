package main

type Config struct {
	BindAddr      string          `json:"bind_addr"`
	AdvertiseAddr string          `json:"advertise_addr"`
	Drivers       []*DriverConfig `json:"drivers"`
}

type DriverConfig struct {
	Name                string `json:"name"`
	DialTimeoutSeconds  uint   `json:"dial_timeout_seconds"`
	DialIntervalSeconds uint   `json:"dial_interval_seconds"`
	SocketPath          string `json:"socket_path"`
}
