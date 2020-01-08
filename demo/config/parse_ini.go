package main

import (
	"fmt"
	"log"

	"github.com/gookit/ini"
)

func main() {
	//config.WithOptions(config.ParseEnv)
	c, err := ini.LoadFiles("config/dev.ini")
	//config.AddDriver(ini.Driver)
	//err := config.LoadFiles("config/dev.ini")
	if err != nil {
		log.Fatalf("load config err: %v", err)
	}

	appname, _ := c.String("appinfo.name")
	//appname, _ := config.String("appinfo.name")
	fmt.Println(appname)
}
