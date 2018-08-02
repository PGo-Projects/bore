package utils

import (
	"os"
)

func CreateDirIfNotExist(path string) (err error) {
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		return os.MkdirAll(path, 0755)
	}
	if err != nil {
		return err
	}
	return nil
}
