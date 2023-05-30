package utils

import (
	"errors"
	"log"
)

func LogErrorWithPanic(err error) {
	if err != nil {
		log.Panicf("error : %s", err)
	}
}

func ErrorWithReturn(message string) error {
	return errors.New(message)
}
