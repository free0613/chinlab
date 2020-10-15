package main

import (
	"github.com/gookit/ini"
	"log"
	"resk.micro/infra"
)

func main() {

	files, err := ini.LoadFiles("config/dev.ini")

	if err != nil {
		log.Fatalf("load config err: %v", err)
	}

	boot := infra.New(files)
	boot.Start()
}
