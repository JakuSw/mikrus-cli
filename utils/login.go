package utils

import (
	"os"
)

func CanLogin() bool {
	srv, _ := os.LookupEnv("MIKRUS_SRV")
	key, _ := os.LookupEnv("MIKRUS_KEY")
	if srv == "" || key == "" {
		return false
	}
	return true
}
