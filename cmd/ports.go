package cmd

import "mikrus-cli/utils"

func GetPorts() {
	response := utils.MakeRequestFor("porty")
	utils.PrintResponse(response)
}
