package tests

import (
	"testing"

	"github.com/limon4ik-black/in_memory_key_value/internal/logger"
	"github.com/limon4ik-black/in_memory_key_value/internal/model"
	"github.com/limon4ik-black/in_memory_key_value/internal/storage"
)

func TestStorage(t *testing.T) {
	logger.StartLog()
	s := storage.NewStorage()

	t.Run("SET operation", func(t *testing.T) {

		query := model.Query{Head: "SET", Argument1: "key1", Argument2: "value1"}
		str, err := s.Distribution(query)
		if str == "" || err == nil {
			logger.Log.Errorw("unknown error")
		}
		value, err := s.Get("key1")
		if err != nil {
			t.Errorf("GET after SET failed: %v", err)
		}
		if value != "value1" {
			t.Errorf("Expected 'value1', got '%s'", value)
		}
	})

	t.Run("GET operation", func(t *testing.T) {

		value, err := s.Get("key1")
		if err != nil {
			t.Errorf("GET failed: %v", err)
		}
		if value != "value1" {
			t.Errorf("Expected 'value1', got '%s'", value)
		}

		_, err = s.Get("non_existent")
		if err == nil {
			t.Error("GET should fail for non-existent key")
		}
	})

	t.Run("DEL operation", func(t *testing.T) {

		_, err := s.Del("key1")
		if err != nil {
			t.Errorf("DEL failed: %v", err)
		}

		_, err = s.Get("key1")
		if err == nil {
			t.Error("GET should fail for deleted key")
		}

		_, err = s.Del("non_existent")
		if err == nil {
			t.Error("DEL should fail for non-existent key")
		}

	})

}
