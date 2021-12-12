package utils

import (
	"io"
	"log"
	"net/http"
)

func MakeRequestFor(module string) string {
	url := "https://api.mikr.us/" + module
	resp, err := http.PostForm(url, GetAuthData())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	return bodyString
}
