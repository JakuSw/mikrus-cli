package cmd

import (
	"encoding/json"
	"fmt"
	"mikrus-cli/utils"
	"os"
)

func GetLogs() {
	if len(os.Args) != 3 {
		response := utils.MakeRequestFor("logs")
		var logs []LogMessage
		json.Unmarshal([]byte(response), &logs)
		for _, log := range logs {
			PrintLog(log)
		}
		return
	}
	target := "logs/" + os.Args[2]
	response := utils.MakeRequestFor(target)
	var logMessage LogMessage
	json.Unmarshal([]byte(response), &logMessage)
	PrintLog(logMessage)
}

func PrintLog(message LogMessage) {
	fmt.Println("Id:", message.Id)
	fmt.Println("ServerId:", message.ServerId)
	fmt.Println("Task:", message.Task)
	fmt.Println("WhenCreated:", message.WhenCreated)
	fmt.Println("WhenDone:", message.WhenDone)
	fmt.Println("Output:", message.Output)
}

type LogMessage struct {
	Id          string `json:"id"`
	ServerId    string `json:"server_id"`
	Task        string `json:"task"`
	WhenCreated string `json:"when_created"`
	WhenDone    string `json:"when_done"`
	Output      string `json:"output"`
}
