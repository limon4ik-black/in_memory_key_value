package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/limon4ik-black/in_memory_key_value/internal/logger"
)

func main() {
	conn, err := net.Dial("tcp", ":3223")
	if err != nil {
		fmt.Println("Ошибка подключения:", err)
		return
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
			fmt.Println("Ошибка ввода:", err)
			continue
		}
		text = strings.TrimSpace(text)

		_, err = conn.Write([]byte(text))
		if err != nil {
			fmt.Println("Ошибка отправки:", err)
			return
		}

		buff := make([]byte, 1024)
		n, err := conn.Read(buff)
		if err != nil {
			fmt.Println("Ошибка чтения:", err)
			break
		}
		fmt.Println("Ответ сервера:", string(buff[:n]))
	}
}
