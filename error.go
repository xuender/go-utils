package utils

import "log"

// CheckError check error.
func CheckError(err error) {
	if nil != err {
		log.Panicln(err)
	}
}
