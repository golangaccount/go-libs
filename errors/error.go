package errors

import (
	"errors"
)

func New(err string) error {
	return errors.New(err)
}

func Error(err error, str string) error {
	if err != nil {
		return err
	} else {
		return New(str)
	}
}
