package tests

import (
	"testing"

	"github.com/limon4ik-black/in_memory_key_value/internal/compute"
	"github.com/limon4ik-black/in_memory_key_value/internal/errors"
	"github.com/limon4ik-black/in_memory_key_value/internal/logger"
)

func compareErrors(err1, err2 error) bool {
	if err1 == nil && err2 == nil {
		return true
	}
	if err1 == nil || err2 == nil {
		return false
	}
	return err1.Error() == err2.Error()
}

func TestReception(t *testing.T) {

	logger.StartLog()

	testTable := []struct {
		command  string
		expected error
	}{
		{
			command:  "SET arg1 arg2",
			expected: nil,
		},
		{
			command:  "GET arg1",
			expected: nil,
		},
		{
			command:  "DEL arg1",
			expected: nil,
		},
		{
			command:  "SET",
			expected: errors.IncorrectNOA(),
		},
		{
			command:  "GET",
			expected: errors.IncorrectNOA(),
		},
		{
			command:  "DEL",
			expected: errors.IncorrectNOA(),
		},
		{
			command:  "set arg1 arg2",
			expected: errors.IncorrectCommandWord(),
		},
		{
			command:  "get arg1",
			expected: errors.IncorrectCommandWord(),
		},
		{
			command:  "del arg1",
			expected: errors.IncorrectCommandWord(),
		},
		{
			command:  "GET arg1()",
			expected: errors.IncorrectSymbols(),
		},
	}

	for _, testcase := range testTable {
		result := compute.Reception(testcase.command)

		if !compareErrors(result, testcase.expected) {
			t.Errorf("Incorrect result for command '%s'. Expect %v, got %v",
				testcase.command, testcase.expected, result)
		}
	}
}
