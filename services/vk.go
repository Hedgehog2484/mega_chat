package services

import (
	"github.com/nikepan/govkbot/v2"
	"log"
	"mega_chat/custom_types"
	"os"
	"strconv"
)

func CreateVkBot(token string) {
	govkbot.HandleMessage("", messageHandler)

	govkbot.SetAPI(token, "", "5.89")
	me, _ := govkbot.API.Me()
	log.Println(me.FullName())
	err := govkbot.Listen(token, "", "5.89", 819438293)
	if err != nil {
		log.Fatalln("VK error:", err)
	}
}

func messageHandler(msg *govkbot.Message) (reply string) {
	VkChatId, _ := strconv.Atoi(os.Getenv("VK_CHAT_ID"))
	me, _ := govkbot.API.Me()
	if msg.ChatID == VkChatId && msg.UserID != me.ID {
		log.Println(msg.UserID)
		var user, err = govkbot.API.User(msg.UserID)
		if err != nil {
			log.Fatalln("Errror get VK user:", err)
		}
		var message = custom_types.Message{Text: msg.Body, Author: user.FullName(), Attachments: nil}
		queue = append(queue, message)
		err = msg.MarkAsRead()
		if err != nil {
			return ""
		}
	}
	return ""
}
