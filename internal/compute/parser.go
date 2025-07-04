package compute

import (
	"strings"

	"github.com/limon4ik-black/in_memory_key_value/internal/custome_errors"
	"github.com/limon4ik-black/in_memory_key_value/internal/logger"
	"github.com/limon4ik-black/in_memory_key_value/internal/model"
	"github.com/limon4ik-black/in_memory_key_value/internal/storage"
)

func Reception(command string) (string, error) {
	words := strings.Fields(command)

	if command == "" {
		return custome_errors.QueryIsEmpty().Error(), custome_errors.QueryIsEmpty()
	}

	if words[0] != "SET" && words[0] != "DEL" && words[0] != "GET" {
		logger.Log.Errorw("incorresct command word in request", "command", command)
		return custome_errors.IncorrectCommandWord().Error(), custome_errors.IncorrectCommandWord()
	}

	if words[0] == "SET" && len(words) != 3 {
		logger.Log.Errorw("incorrect number of arguments in the query", "commmand", command)
		return custome_errors.IncorrectNOA().Error(), custome_errors.IncorrectNOA()
	}

	if words[0] == "GET" && len(words) != 2 {
		logger.Log.Errorw("incorrect number of arguments in the query", "commmand", command)
		return custome_errors.IncorrectNOA().Error(), custome_errors.IncorrectNOA()
	}

	if words[0] == "DEL" && len(words) != 2 {
		logger.Log.Errorw("incorrect number of arguments in the query", "commmand", command)
		return custome_errors.IncorrectNOA().Error(), custome_errors.IncorrectNOA()
	}

	for i := 1; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ {
			if !(words[i][j] >= 'A' && words[i][j] <= 'Z') && !(words[i][j] >= 'a' && words[i][j] <= 'z') &&
				words[i][j] != '*' && words[i][j] != '/' && words[i][j] != '_' && !(words[i][j] >= '0' && words[i][j] <= '9') {
				logger.Log.Errorw("incorrect symbols in request", "command", command)
				return custome_errors.IncorrectSymbols().Error(), custome_errors.IncorrectSymbols()
			}
		}
	}

	return Parse(command)
}

func Parse(command string) (string, error) {
	words := strings.Fields(command)
	var query model.Query
	if len(words) == 3 {
		query = model.Query{Head: words[0], Argument1: words[1], Argument2: words[2]}
	}
	if len(words) == 2 {
		query = model.Query{Head: words[0], Argument1: words[1]}
	}
	return storage.Distribution(query)
}
