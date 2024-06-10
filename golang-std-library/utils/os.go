package utils

import "os"

func GetArgs() []string {
	return os.Args
}

func GetHostname() (string, error) {
	return os.Hostname()
}
