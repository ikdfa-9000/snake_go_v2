package main

import (
	"example.com/snake_go/pkg"
	"example.com/snake_go/pkg/cfg"
	"log"
)

func main() {
	config, err := cfg.InitJsonConfig("")

	if err != nil {
		log.Fatal("ошибка чтение конфига: " + err.Error())
	}

	pkg.Run(config)
}
