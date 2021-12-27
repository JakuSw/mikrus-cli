package main

import (
	"mikrus-cli/cmd"
	"mikrus-cli/utils"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		utils.ModuleError()
		return
	}
	if !utils.CanLogin() {
		utils.LoginError()
		return
	}

	switch os.Args[1] {
	case "info":
		cmd.GetInfo()
	case "logs":
		cmd.GetLogs()
	case "restart":
		cmd.Restart()
	case "servers":
		cmd.GetServers()
	case "ports":
		cmd.GetPorts()
	case "db":
		cmd.Db()
	case "stats":
		cmd.GetStats()
	case "exec":
		cmd.Exec()
	case "amph":
		cmd.Amphetamines()
	default:
		utils.ModuleError()
	}
}
