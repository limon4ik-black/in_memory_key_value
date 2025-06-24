package server

import (
	"fmt"
	"net"

	"github.com/limon4ik-black/in_memory_key_value/internal/compute"
	"github.com/limon4ik-black/in_memory_key_value/internal/logger"
)

func Processing() {
	listener, err := net.Listen("tcp", ":3223")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := listener.Close(); err != nil {
			logger.Log.Errorw("failed to close conn: %v", err)
		}
	}()

	workerCount := 10
	channel := make(chan net.Conn)

	for i := 0; i < workerCount; i++ {
		go workers(channel)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		channel <- conn
	}

}

func workers(connChan <-chan net.Conn) {
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
		input := make([]byte, (1024 * 4))
		n, err := conn.Read(input)
		if n == 0 || err != nil {
			fmt.Println(err)
			break
		}
		query := string(input[0:n])

		target, _ := compute.Reception(query)

		fmt.Println(query, "-", target)

		_, errs := conn.Write([]byte(target))
		if errs != nil {
			logger.Log.Errorw("failed to write to conn: %v", err)

		}

	}
}
