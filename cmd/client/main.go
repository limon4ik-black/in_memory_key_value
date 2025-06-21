package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", ":3223")
	if err != nil {
		fmt.Println("Ошибка подключения:", err)
		return
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка ввода:", err)
			continue
		}
		text = strings.TrimSpace(text) // убираем \n

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
