package compute

import (
	"in_memory_key_value/internal/model"
	"in_memory_key_value/internal/storage"
	"in_memory_key_value/logger"
	"strings"
)

func Reception(command string) bool {
	words := strings.Fields(command)

	if len(words) != 3 && len(words) != 2 {
		logger.Log.Errorw("Неверное кол-во слов в запросе", "command", command)
		return false
	}

	if words[0] == "SET" && len(words) != 3 {
		logger.Log.Errorw("Неверное кол-во аргументов в запросе", "commmand", command)
		return false
	}

	if words[0] == "GET" && len(words) != 2 {
		logger.Log.Errorw("Неверное кол-во аргументов в запросе", "commmand", command)
		return false
	}

	if words[0] == "DEL" && len(words) != 2 {
		logger.Log.Errorw("Неверное кол-во аргументов в запросе", "commmand", command)
		return false
	}

	if words[0] != "SET" && words[0] != "DEL" && words[0] != "GET" {
		logger.Log.Errorw("Некорректное командное слово в запросе", "command", command)
		return false
	}

	for i := 1; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ {
			if !(words[i][j] >= 'A' && words[i][j] <= 'Z') && !(words[i][j] >= 'a' && words[i][j] <= 'z') &&
				words[i][j] != '*' && words[i][j] != '/' && words[i][j] != '_' && !(words[i][j] >= '0' && words[i][j] <= '9') {
				logger.Log.Errorw("Некорректные символы в запросе", "command", command)
				return false
			}
		}
	}
	//fmt.Println("Parsed:", words, len(words))
	Parse(command)
	return true
}

func Parse(command string) {
	words := strings.Fields(command)
	var query model.Query
	if len(words) == 3 {
		query = model.Query{Head: words[0], Argument1: words[1], Argument2: words[2]}
	}
	if len(words) == 2 {
		query = model.Query{Head: words[0], Argument1: words[1]}
	}
	logger.Log.Infow("Запрос преобразован в структуру", "head", query.Head, "arg1", query.Argument1, "arg2", query.Argument2)
	storage.Distribution(query)
}
