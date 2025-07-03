package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/limon4ik-black/in_memory_key_value/internal/config"
	"github.com/limon4ik-black/in_memory_key_value/internal/logger"
)

var buff = make([]byte, 1024)

func main() {
	configPath := flag.String("config-file", "", "Path to config file (YAML)")
	flag.Parse()

	logger.StartLog()

	if *configPath == "" {
		fmt.Println("Usage: ./server --config-file=path/to/config.yml")
		os.Exit(1)
	}
	config.LoadConfig(*configPath)

	address := config.AppConfig.Network.Address
	conn, err := net.Dial("tcp", address)
	if err != nil {
		logger.Log.Errorw("connection error: %v", err)
		os.Exit(1)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			logger.Log.Errorw("failed to close conn: %v", err)
		}
	}()

	reader := bufio.NewReader(os.Stdin)

	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			logger.Log.Errorw("input error: %v", err)
			continue
		}
		text = strings.TrimSpace(text)
		if text == "" {
			logger.Log.Errorw("empty input error: %v", err)
			continue
		}
		_, err = conn.Write([]byte(text))
		if err != nil {
			logger.Log.Errorw("sending error: %v", err)
			os.Exit(1)
		}

		//buff := make([]byte, 1024)
		n, err := conn.Read(buff)
		if err != nil {
			logger.Log.Errorw("reading error: %v", err)
			break
		}
		fmt.Println("Server response:", string(buff[:n]))
	}
}
