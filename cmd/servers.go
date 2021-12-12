package cmd

import "mikrus-cli/utils"

func GetServers() {
	response := utils.MakeRequestFor("serwery")
	utils.PrintResponse(response)
}
