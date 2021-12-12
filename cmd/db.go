package cmd

import "mikrus-cli/utils"

func Db() {
	response := utils.MakeRequestFor("db")
	utils.PrintResponse(response)
}
