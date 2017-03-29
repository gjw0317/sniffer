package base

import (
	"os"
)

func IsFileExists(name string) bool {
	_, err := os.Stat(name)
	return err == nil || os.IsExist(err)
}

func IsDirExists(name string) bool {
	d, err := os.Stat(name)
	return (err == nil || os.IsExist(err)) && d.IsDir()
}
