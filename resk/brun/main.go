package main

import (
	"log"

	"github.com/gookit/ini"
	"resk.micro/infra"
	_ "resk.micro/resk"
)

func main() {
	files, err := ini.LoadFiles("config/dev.ini")
	if err != nil {
		log.Fatalf("load config err: %v", err)
	}
	boot := infra.New(files)
	boot.Startup()

}
