package cmd

import (
	"encoding/json"
	"fmt"
	"mikrus-cli/utils"
)

func GetInfo() {
	response := utils.MakeRequestFor("info")
	var serverInfo ServerInfo
	json.Unmarshal([]byte(response), &serverInfo)
	fmt.Println("ServerId:", serverInfo.ServerId)
	fmt.Println("ServerName:", serverInfo.ServerName)
	fmt.Println("Expires:", serverInfo.Expires)
	fmt.Println("ExpiresCytrus:", serverInfo.ExpiresCytrus)
	fmt.Println("ExpiresStorage:", serverInfo.ExpiresStorage)
	fmt.Println("ParamRam:", serverInfo.ParamRam)
	fmt.Println("ParamDisk:", serverInfo.ParamDisk)
	fmt.Println("LastlogPanel:", serverInfo.LastlogPanel)

}

type ServerInfo struct {
	ServerId       string `json:"server_id"`
	ServerName     string `json:"server_name"`
	Expires        string `json:"expires"`
	ExpiresCytrus  string `json:"expires_cytrus"`
	ExpiresStorage string `json:"expires_storage"`
	ParamRam       string `json:"param_ram"`
	ParamDisk      string `json:"param_disk"`
	LastlogPanel   string `json:"lastlog_panel"`
}
