package storage

import (
	"fmt"
	"sync"

	"github.com/limon4ik-black/in_memory_key_value/internal/custome_errors"
	"github.com/limon4ik-black/in_memory_key_value/internal/logger"
	"github.com/limon4ik-black/in_memory_key_value/internal/model"
)

var globalStorage = NewStorage()

type Storage struct {
	mapa  map[string]string
	mutex sync.RWMutex
}

func NewStorage() *Storage {
	return &Storage{mapa: make(map[string]string)}
}

func Distribution(query model.Query) (string, error) {
	return globalStorage.Distribution(query)
}

func (s *Storage) Distribution(query model.Query) (string, error) {
	switch query.Head {
	case "SET":
		return s.Set(query.Argument1, query.Argument2)
	case "GET":
		return s.Get(query.Argument1)
	case "DEL":
		return s.Del(query.Argument1)
	default:
		return "", nil
	}
}

func (s *Storage) Set(arg1 string, arg2 string) (string, error) {
	s.mutex.Lock()
	s.mapa[arg1] = arg2
	s.mutex.Unlock()
	logger.Log.Infow("command", "SET", "successfully")
	return "data saved successfully", nil
}

func (s *Storage) Get(arg1 string) (string, error) {
	s.mutex.RLock()
	value, ok := s.mapa[arg1]
	s.mutex.RUnlock()
	if ok {
		fmt.Println(value)
		logger.Log.Infow("command", "GET", "successfully")
		return value, nil
	}
	fmt.Println("such a key does not exist")
	logger.Log.Errorw("command", "GET", "non-existent key")
	return "such a key does not exist", custome_errors.NonExistent()

}

func (s *Storage) Del(arg1 string) (string, error) {
	s.mutex.Lock()
	_, ok := s.mapa[arg1]
	delete(s.mapa, arg1)
	s.mutex.Unlock()
	if ok {
		logger.Log.Infow("command", "DEL", "successfully")
		return "data deleted", nil
	}
	logger.Log.Errorw("command", "DEL", "non-existent key")
	return "such a key does not exist", custome_errors.NonExistent()

}
