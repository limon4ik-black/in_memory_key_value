package main

import (
	"github.com/limon4ik-black/in_memory_key_value/internal/config"
	"github.com/limon4ik-black/in_memory_key_value/internal/logger"
	"github.com/limon4ik-black/in_memory_key_value/internal/server"
)

func main() {
	config.LoadConfig("/Users/limon4ik/Desktop/in_memory_key_value/internal/config/config.yml")
	logger.StartLog()
	logger.Log.Infow("project launch")
	server.Processing()
}
