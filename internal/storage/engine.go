package storage

import (
	"fmt"
	"in_memory_key_value/internal/model"
	"in_memory_key_value/logger"
)

var mapa = make(map[string]string)

func Distribution(query model.Query) {
	if query.Head == "SET" {
		HandleSet(query.Argument1, query.Argument2)
	}

	if query.Head == "GET" {
		HandleGet(query.Argument1)
	}

	if query.Head == "DEL" {
		HandleDel(query.Argument1)
	}
}

func HandleSet(arg1 string, arg2 string) {
	mapa[arg1] = arg2
	fmt.Println("Данные успешно сохранены")
	logger.Log.Infow("Команда", "SET", "Успешно")
}

func HandleGet(arg1 string) {
	value, ok := mapa[arg1]
	if ok {
		fmt.Println(value)
		logger.Log.Infow("Команда", "GET", "Успешно")
	} else {
		fmt.Println("Такого ключа не существует")
		logger.Log.Errorw("Команда", "GET", "Несуществующий ключ")
	}
}

func HandleDel(arg1 string) {
	delete(mapa, arg1)
	fmt.Println("Данные удалены")
	logger.Log.Infow("Команда", "DEL", "Успешно")
}
