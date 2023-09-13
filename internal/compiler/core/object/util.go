package object

import (
	"fmt"
)

func runtimeError(format string, a ...interface{}) *Error {
	return &Error{Message: fmt.Sprintf(format, a...)}
}
