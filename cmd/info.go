package cmd

import (
	"mikrus-cli/utils"
)

func GetInfo() {
	response := utils.MakeRequestFor("info")
	utils.PrintResponse(response)
}
