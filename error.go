package utils

import "log"

// CheckError 检查错误.
func CheckError(err error) {
	if nil != err {
		log.Panicln(err)
	}
}
