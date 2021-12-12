package cmd

import "mikrus-cli/utils"

func GetStats() {
	response := utils.MakeRequestFor("stats")
	utils.PrintResponse(response)
}
