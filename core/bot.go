/*
 * Copyright (c) 2020. All rights reserved.
 */

package core

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"image/jpeg"
	"log"
	"strconv"
	"strings"

	"github.com/valyala/fasthttp"
)

func handle(ctx *fasthttp.RequestCtx, config *Config) {
	vk := config.VK
	body := ctx.PostBody()

	var event Event
	err := event.UnmarshalJSON(body)
	if err != nil {
		log.Println(err)
		return
	}

	switch event.Type {
	case "confirm":
		hash := md5.New()
		_, _ = hash.Write([]byte(strconv.Itoa(config.ID) + config.Settings.Token))
		confirm := hex.EncodeToString(hash.Sum(nil))
		_, _ = ctx.WriteString(confirm)
	case "invite", "ban_expired":
		var obj Invite
		err = obj.UnmarshalJSON(event.Data)
		if err != nil {
			log.Println(err)
			return
		}

		chatID, ok := config.Chats[obj.Chat]
		if !ok {
			log.Println("not found chat")
			return
		}

		code := fmt.Sprintf(`if (API.friends.areFriends({user_ids: %d})[0].friend_status == 3){return API.messages.addChatUser({chat_id: %d, user_id: %d});}return 0;`, obj.User, chatID, obj.User)
		err = vk.Execute(code, nil)
		if err != nil {
			log.Println(err)
			return
		}
	case "delete_for_all":
		var obj DeleteForAll
		err = obj.UnmarshalJSON(event.Data)
		if err != nil {
			log.Println(err)
			return
		}

		chatID, ok := config.Chats[obj.Chat]
		if !ok {
			log.Println("not found chat")
			return
		}

		code := fmt.Sprintf(`return API.messages.delete({delete_for_all: 1, message_ids: API.messages.getByConversationMessageId({peer_id: %d, conversation_message_ids: [%s]}).items@.id});`, chatID+2e9, toArray(obj.ConversationMessageIDs))
		err = vk.Execute(code, nil)
		if err != nil {
			log.Println(err)
			return
		}
	case "message_pin":
		var obj MessagePin
		err = obj.UnmarshalJSON(event.Data)
		if err != nil {
			log.Println(err)
			return
		}

		chatID, ok := config.Chats[obj.Chat]
		if !ok {
			log.Println("not found chat")
			return
		}

		code := fmt.Sprintf(`return API.messages.pin({peer_id: %d, message_id: API.messages.getByConversationMessageId({peer_id: %d, conversation_message_ids: [%d]}).items@.id[0]});`, chatID+2e9, chatID+2e9, obj.ConversationMessageID)
		err = vk.Execute(code, nil)
		if err != nil {
			log.Println(err)
			return
		}
	case "photo_update":
		var obj PhotoUpdate
		err = obj.UnmarshalJSON(event.Data)
		if err != nil {
			log.Println(err)
			return
		}

		chatID, ok := config.Chats[obj.Chat]
		if !ok {
			log.Println("not found chat")
			return
		}

		reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(obj.Photo))
		m, err := jpeg.Decode(reader)
		if err != nil {
			log.Println(err)
			return
		}

		buf := &bytes.Buffer{}
		err = jpeg.Encode(buf, m, nil)
		if err != nil {
			log.Println(err)
			return
		}

		_, err = vk.UploadChatPhoto(chatID, buf)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
