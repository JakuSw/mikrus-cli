package cmd

import (
	"encoding/json"
	"fmt"
	"mikrus-cli/utils"
)

func Amphetamines() {
	response := utils.MakeRequestFor("amfetamina")
	var amphetamine Amphetamine
	json.Unmarshal([]byte(response), &amphetamine)
	fmt.Println("Msg:", amphetamine.Msg)
	fmt.Println("TaskId:", amphetamine.TaskId)
}

type Amphetamine struct {
	Msg    string `json:"msg"`
	TaskId int    `json:"task_id"`
}
