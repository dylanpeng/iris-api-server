package exception

import (
	"fmt"
	"juggernaut/common"
	"juggernaut/lib/logger"
)

type Exception struct {
	code    int64
	message string
}

func (e *Exception) GetCode() int64 {
	return e.code
}

func (e *Exception) GetMessage() string {
	return e.message
}

func (e *Exception) String() string {
	return fmt.Sprintf("(%d) %s", e.code, e.message)
}

func Desc(code int64) string {
	if e, ok := Desces[code]; ok {
		return e
	}

	return "server internal error"
}

func New(code int64, args ...interface{}) *Exception {
	if len(args) > 0 {
		if err, ok := args[0].(error); ok {
			common.Logger.Log(logger.ErrorLevel, "Error: %s | Args: %+v", err, args[1:])
		}
	}

	return &Exception{code: code, message: Desc(code)}
}
