package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/limon4ik-black/in_memory_key_value/internal/config"
	"github.com/limon4ik-black/in_memory_key_value/internal/logger"
	"github.com/limon4ik-black/in_memory_key_value/internal/server"
)

func main() {
	logger.StartLog()
	configPath := flag.String("config-file", "", "Path to config file (YAML)")
	flag.Parse()

	if *configPath == "" {
		fmt.Println("Usage: ./server --config-file=path/to/config.yml")
		os.Exit(1)
	}
	config.LoadConfig(*configPath)

	logger.Log.Infow("project launch")
	server.Processing()
}
