package cmd

import (
	"encoding/json"
	"fmt"
	"mikrus-cli/utils"
)

func Db() {
	response := utils.MakeRequestFor("db")
	var db Database
	json.Unmarshal([]byte(response), &db)
	fmt.Println("Mysql:", db.Mysql)
	fmt.Println("Psql:", db.Psql)
	fmt.Println("Mongo:", db.Mongo)
}

type Database struct {
	Mysql string `json:"mysql"`
	Psql  string `json:"psql"`
	Mongo string `json:"mongo"`
}
