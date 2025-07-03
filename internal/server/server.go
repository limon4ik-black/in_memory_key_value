package server

import (
	"fmt"
	"net"
	"os"

	"github.com/limon4ik-black/in_memory_key_value/internal/compute"
	"github.com/limon4ik-black/in_memory_key_value/internal/config"
	"github.com/limon4ik-black/in_memory_key_value/internal/logger"
	"github.com/limon4ik-black/in_memory_key_value/internal/wal"
)

var (
	input = make([]byte, 1024*4)
	w     *wal.WAL
)

func Processing() {
	var err error

	w, err = wal.InitWal("./internal/wal/wals", 1024*1024)
	if err != nil {
		logger.Log.Errorw("Error create first wal: %v", err)
		os.Exit(1)
	}
	w.Load()
	defer func() {
		if w != nil {
			if err := w.Close(); err != nil {
				logger.Log.Errorw("failed to close WAL: %v", err)
			}
		}
	}()

	address := config.AppConfig.Network.Address
	listener, err := net.Listen("tcp", address)
	if err != nil {
		logger.Log.Errorw("connection error: %v", err)
		os.Exit(1)
	}
	defer func() {
		if err := listener.Close(); err != nil {
			logger.Log.Errorw("failed to close conn: %v", err)
		}
	}()

	workerCount := 10
	channel := make(chan net.Conn, 1)

	for i := 0; i < workerCount; i++ {
		go StartWorkerPool(channel)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		channel <- conn
	}
}

func StartWorkerPool(connChan <-chan net.Conn) {
	for conn := range connChan {
		HandleConnections(conn)
	}
}

func HandleConnections(conn net.Conn) {
	defer func() {
		if err := conn.Close(); err != nil {
			logger.Log.Errorw("failed to close conn: %v", err)
		}
	}()

	for {
		n, err := conn.Read(input)
		if n == 0 || err != nil {
			logger.Log.Errorw("reading error: %v", err)
			break
		}
		query := string(input[0:n])

		if w != nil {
			if err := w.WriteToWal(query); err != nil {
				logger.Log.Errorw("failed to write to WAL: %v", err)
				os.Exit(1)
			}
		}

		target, _ := compute.Reception(query)

		fmt.Println(query, "-", target)

		_, errs := conn.Write([]byte(target))
		if errs != nil {
			logger.Log.Errorw("failed to write to conn: %v", errs)
		}
	}
}
