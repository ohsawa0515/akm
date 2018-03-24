package utils

import (
	"os"
	"os/user"
	"runtime"
)

func GetHomeDir() string {
	usr, err := user.Current()
	if err != nil || len(usr.HomeDir) > 0 {
		return usr.HomeDir
	}
	if runtime.GOOS == "windows" {
		return os.Getenv("HOMEPATH")
	}
	return os.Getenv("HOME")
}

func IsExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
