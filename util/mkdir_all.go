package util

import (
	"errors"
	"os"
)

func MkdirAllIfNotExist(name string) error {
	_, err := os.Stat(name)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return err
		}

		if err := os.MkdirAll(name, os.ModePerm); err != nil {
			return err
		}
	}

	return nil
}
