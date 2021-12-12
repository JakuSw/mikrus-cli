package cmd

import "mikrus-cli/utils"

func Restart() {
	response := utils.MakeRequestFor("restart")
	utils.PrintResponse(response)
}
