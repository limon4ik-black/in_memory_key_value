package main

import (
	"github.com/limon4ik-black/in_memory_key_value/internal/logger"
	"github.com/limon4ik-black/in_memory_key_value/internal/server"
)

func main() {
	logger.StartLog()
	logger.Log.Infow("project launch")
	server.Processing()
	// var wg sync.WaitGroup

	// channel := make(chan string, 100)

	// workerCount := 10

	// for i := 0; i < workerCount; i++ {
	// 	wg.Add(1)
	// 	go func() {
	// 		defer wg.Done()
	// 		for command := range channel {
	// 			logger.Log.Infow("request received by worker", "command", command)
	// 			err := compute.Reception(command)
	// 			if err != nil {
	// 				fmt.Println("wrong request")
	// 				logger.Log.Errorw("wrong request: ", "command", command)
	// 			}
	// 		}
	// 	}()
	// }

	// scanner := bufio.NewScanner(os.Stdin)
	// for {
	// 	if !scanner.Scan() {
	// 		break
	// 	}
	// 	command := scanner.Text()
	// 	channel <- command
	// }

	// close(channel)
	// wg.Wait()
}
