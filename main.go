package main

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

func main() {

	inidata, err := ini.Load("config.ini")
	if err != nil {
		fmt.Printf("Fail to read the config file: %v", err)
		os.Exit(1)
	}

	env := inidata.Section("environment")
	sistem := inidata.Section("sistem")
	database := inidata.Section("database")

	userdb := database.Key("user").String()
	passdb := database.Key("password").String()
	ipdb := database.Key("ip").String()
	portdb, _ := database.Key("port").Int()

	rootpass := sistem.Key("rootpassword").String()

	ip := env.Key("ip").String()
	port, _ := env.Key("port").Int()

	db := initDB(userdb, passdb, ipdb, portdb)
	initRouters(ip, port, rootpass, db)
}
