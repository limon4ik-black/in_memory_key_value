package main

import (
	"github.com/limon4ik-black/in_memory_key_value/internal/logger"
	"github.com/limon4ik-black/in_memory_key_value/internal/server"
)

func main() {
	logger.StartLog()
	logger.Log.Infow("project launch")
	server.Processing()
}
