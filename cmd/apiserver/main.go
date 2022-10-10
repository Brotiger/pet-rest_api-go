package main

import (
	"flag"
	"log"

	"github.com/Brotiger/rest_api_gopher.git/internal/app/apiserver"
	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

func init() {
	//Флаг при запуске бинарника.
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	//Берем значение из флага и помещаем в переменную
	flag.Parse()

	//Созданиеконфига с дефолтными значениями
	config := apiserver.NewConfig()

	//Декодируем toml файл по тегам
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
