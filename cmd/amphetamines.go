package cmd

import "mikrus-cli/utils"

func Amphetamines() {
	response := utils.MakeRequestFor("amfetamina")
	utils.PrintResponse(response)
}
