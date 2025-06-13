package storage

import (
	"fmt"
	"sync"

	"github.com/limon4ik-black/in_memory_key_value/internal/errors"
	"github.com/limon4ik-black/in_memory_key_value/internal/logger"
	"github.com/limon4ik-black/in_memory_key_value/internal/model"
)

var mapa = make(map[string]string)
var mutex sync.RWMutex

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

func HandleSet(arg1 string, arg2 string) { //хуй пойми, что вернуть
	mutex.Lock()
	mapa[arg1] = arg2
	mutex.Unlock()
	fmt.Println("data saved successfully")
	logger.Log.Infow("command", "SET", "successfully")
}

func HandleGet(arg1 string) (string, error) {
	mutex.RLock()
	value, ok := mapa[arg1]
	mutex.RUnlock()
	if ok {
		fmt.Println(value)
		logger.Log.Infow("command", "GET", "successfully")
		return value, nil
	} else {
		fmt.Println("such a key does not exist")
		logger.Log.Errorw("command", "GET", "non-existent key")
		return "", errors.NonExistent()
	}
}

func HandleDel(arg1 string) (bool, error) {
	mutex.Lock()
	_, ok := mapa[arg1]
	delete(mapa, arg1)
	mutex.Unlock()
	if ok {
		fmt.Println("data deleted")
		logger.Log.Infow("command", "DEL", "ssuccessfully")
		return ok, nil
	} else {
		fmt.Println("such a key does not exist")
		logger.Log.Errorw("command", "DEL", "non-existent key")
		return ok, errors.NonExistent()
	}
}
