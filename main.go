/*
 * Copyright (c) 2020. All rights reserved.
 */

package main

import (
	"encoding/json"
	"fasthook/core"
	"io/ioutil"
	"log"

	"github.com/SevereCloud/vksdk/api"
)

func main() {
	// Инициализация переменной конфига
	config := core.NewConfig()

	// Чтение файла конфига
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatalln(err)
	}

	// Разбор JSON конфига
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatalln(err)
	}

	// Присваивание объекта ВК к конфигу
	config.VK = api.NewVK(config.Settings.AccessToken)

	// Проверка токена + получение своего ID
	res, err := config.VK.UsersGet(nil)
	if err != nil {
		log.Fatalln(err)
	}

	// Присваивание своего ID конфигу
	config.ID = res[0].ID

	// Запуск веб-сервера
	core.Start(config)
}
