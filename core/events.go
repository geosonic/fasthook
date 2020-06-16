/*
 * Copyright (c) 2020. All rights reserved.
 */

package core

import "encoding/json"

//go:generate easyjson -all events.go

// События
// "type":
// "confirm" - подтверждение сервера
// "invite", "ban_expired" - приглашение пользователя в беседу
// "delete_for_all" - удаление сообщений для всех
// "message_pin" - закрепление сообщения
// "photo_update" - установка изображения
//easyjson:json
type Event struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

// "type": "invite" и "type": "ban_expired"
// Приглашение пользователя в беседу
// А также возвращение после его бана
//easyjson:json
type Invite struct {
	Chat string `json:"chat"`
	User int    `json:"user"`
}

// "type": "delete_for_all"
// Удаление сообщений для всех
//easyjson:json
type DeleteForAll struct {
	Chat                   string `json:"chat"`
	ConversationMessageIDs []int  `json:"conversation_message_ids"`
}

// "type": "message_pin"
// Закрепление сообщения
//easyjson:json
type MessagePin struct {
	Chat                  string `json:"chat"`
	ConversationMessageID int    `json:"conversation_message_id"`
}

// "type": "photo_update"
// Смена фото
//easyjson:json
type PhotoUpdate struct {
	Chat  string `json:"chat"`
	Photo string `json:"photo"`
}
