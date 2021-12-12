package utils

import (
	"net/url"
	"os"
)

func GetAuthData() url.Values {
	data := url.Values{
		"srv": {os.Getenv("MIKRUS_SRV")},
		"key": {os.Getenv("MIKRUS_KEY")},
	}
	return data
}
