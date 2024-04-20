package main

type Config struct {
	Name       string `json:"name"`
	SocketPath string `json:"socket_path"`
	JreHome    string `json:"jre_home"`
}
