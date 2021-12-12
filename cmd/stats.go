package cmd

import (
	"encoding/json"
	"fmt"
	"mikrus-cli/utils"
)

func GetStats() {
	response := utils.MakeRequestFor("stats")
	var stats Stats
	json.Unmarshal([]byte(response), &stats)
	fmt.Println("Free:", stats.Free)
	fmt.Println("Df:", stats.Df)
	fmt.Println("Uptime:", stats.Uptime)
	fmt.Println("Ps:", stats.Ps)
}

type Stats struct {
	Free   string `json:"free"`
	Df     string `json:"df"`
	Uptime string `json:"uptime"`
	Ps     string `json:"ps"`
}
