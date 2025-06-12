package main

import (
	"bufio"
	"fmt"
	"in_memory_key_value/internal/compute"
	"in_memory_key_value/logger"
	"os"
)

func main() {
	logger.StartLog()
	logger.Log.Infow("Проект запущен")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		command := scanner.Text()
		logger.Log.Infow("Запрос: ", "command", command)
		if !compute.Reception(command) {
			fmt.Println("Неверный запрос")
			logger.Log.Errorw("Неверный запрос: ", "command", command)
		}
	}
}
