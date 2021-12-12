package cmd

import (
	"encoding/json"
	"fmt"
	"mikrus-cli/utils"
)

func GetServers() {
	response := utils.MakeRequestFor("serwery")
	var servers []Server
	json.Unmarshal([]byte(response), &servers)
	for _, server := range servers {
		PrintServer(server)
	}
}

func PrintServer(server Server) {
	fmt.Println("ServerId:", server.ServerId)
	fmt.Println("ServerName:", server.ServerName)
	fmt.Println("Expires:", server.Expires)
	fmt.Println("ParamRam:", server.ParamRam)
	fmt.Println("ParamDisk:", server.ParamDisk)

}

type Server struct {
	ServerId   string `json:"server_id"`
	ServerName string `json:"server_name"`
	Expires    string `json:"expires"`
	ParamRam   string `json:"param_ram"`
	ParamDisk  string `json:"param_disk"`
}
