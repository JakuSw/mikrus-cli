package cmd

import (
	"encoding/json"
	"fmt"
	"mikrus-cli/utils"
)

func Restart() {
	response := utils.MakeRequestFor("restart")
	var restartResult RestartResult
	json.Unmarshal([]byte(response), &restartResult)
	fmt.Println("Msg:", restartResult.Msg)
	fmt.Println("TaskId:", restartResult.TaskId)

}

type RestartResult struct {
	Msg    string `json:"msg"`
	TaskId int    `json:"task_id"`
}
