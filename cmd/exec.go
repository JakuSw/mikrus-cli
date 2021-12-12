package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mikrus-cli/utils"
	"net/http"
	"net/url"
	"os"
)

func Exec() {
	if len(os.Args) < 3 {
		utils.NoCommandProvidedError()
		return
	}
	var command string

	for _, s := range os.Args[2:] {
		command = command + " " + s
	}

	data := url.Values{
		"srv": {os.Getenv("MIKRUS_SRV")},
		"key": {os.Getenv("MIKRUS_KEY")},
		"cmd": {command},
	}
	url := "https://api.mikr.us/exec"
	resp, err := http.PostForm(url, data)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var output ExecOutput
	json.Unmarshal([]byte(string(bodyBytes)), &output)
	fmt.Println("Output:")
	fmt.Println(output.Output)
}

type ExecOutput struct {
	Output string `json:"output"`
}
