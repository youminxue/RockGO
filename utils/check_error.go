package utils

import (
	"errors"
	"github.com/zllangct/RockGO/logger"
	"runtime/debug"
)

func CheckError() {
	if r := recover(); r != nil {
		var str string
		switch r.(type) {
		case error:
			str = r.(error).Error()
		case string:
			str = r.(string)
		}
		err := errors.New(str + "\n" + string(debug.Stack()))
		logger.Error(err)
	}
}
