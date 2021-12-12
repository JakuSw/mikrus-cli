package cmd

import (
	"mikrus-cli/utils"
	"os"
)

func GetLogs() {
	if len(os.Args) != 3 {
		response := utils.MakeRequestFor("logs")
		utils.PrintResponse(response)
		return
	}
	target := "logs/" + os.Args[2]
	response := utils.MakeRequestFor(target)
	utils.PrintResponse(response)
}
