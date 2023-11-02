package main

import (
	"example.com/snake_go/pkg"
	"log"
)

func main() {
	config, err := pkg.InitConfig()
	if err != nil {
		log.Fatal("ошибка чтение конфига: " + err.Error())
	}

	pkg.Run(config)
}
