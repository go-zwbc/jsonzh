package main

import (
	"fmt"

	"github.com/go-zwbc/jsonzh/sure/jsonmzh"
	"github.com/yyle88/must"
)

type Config struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	Mode string `json:"mode"`
}

func main() {
	// Create a config object
	config := Config{
		Host: "localhost",
		Port: 8080,
		Mode: "production",
	}

	// Encode to JSON bytes (panics on failure)
	data := must.Have(jsonmzh.M编码(config))
	fmt.Printf("JSON bytes: %s\n", string(data))

	// Decode from JSON bytes (panics on failure)
	decoded := must.Full(jsonmzh.U解码[Config](data))
	fmt.Printf("Decoded: Host=%s, Port=%d, Mode=%s\n", decoded.Host, decoded.Port, decoded.Mode)
}
