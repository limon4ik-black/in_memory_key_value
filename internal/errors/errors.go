package errors

import (
	"fmt"
)

type ErrorLI struct {
	Code int
	Msg  string
}

func (e *ErrorLI) Error() string {
	return fmt.Sprintf("code %d: %s", e.Code, e.Msg)
}

func DoSomething() error {
	return &ErrorLI{Code: 001, Msg: "XZ"}
}

func NonExistent() error {
	return &ErrorLI{Code: 999, Msg: "non-existent key"}
}

func IncorrectNOW() error {
	return &ErrorLI{Code: 998, Msg: "non-existent key"}
}

func IncorrectNOA() error {
	return &ErrorLI{Code: 997, Msg: "non-existent key"}
}

func IncorrectCommandWord() error {
	return &ErrorLI{Code: 996, Msg: "non-existent key"}
}

func IncorrectSymbols() error {
	return &ErrorLI{Code: 995, Msg: "non-existent key"}
}
