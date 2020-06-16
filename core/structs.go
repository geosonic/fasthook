/*
 * Copyright (c) 2020. All rights reserved.
 */

package core

import "github.com/SevereCloud/vksdk/api"

const version = "1.0"

type Settings struct {
	// Ключ доступа от страницы ВК
	AccessToken string `json:"access_token"`
	// Ключ доступа от чат-менеджера
	Token string `json:"token"`

	// Привязывать ли сервер к чат-менеджеру при запуске
	AutoConnect bool `json:"auto_connect"`
	// Path после ссылки
	Path string `json:"path"`
	// Порт сервера
	Port int `json:"port"`
}

type Chats map[string]int

type Config struct {
	// Настройки бота
	Settings Settings `json:"settings"`
	// Чаты
	Chats Chats `json:"chats"`
	// ID пользователя
	ID int
	// VK
	VK *api.VK
}

func NewConfig() *Config {
	var config Config

	config.Settings.Path = "/"
	config.Settings.Port = 80

	return &config
}
