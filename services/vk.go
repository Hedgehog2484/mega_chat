package services

import (
	"fmt"
	"github.com/nikepan/govkbot/v2"
	"log"
	"mega_chat/custom_types"
	"os"
	"strconv"
)

var vk_queue []custom_types.Message

func CreateVkBot(token string) {
	govkbot.HandleMessage("", messageHandler)

	govkbot.SetAPI(token, "", "5.89")
	me, _ := govkbot.API.Me()
	log.Println(me.FullName())

	go func() {
		err := govkbot.Listen(token, "", "5.89", 819438293)
		if err != nil {
			log.Fatalln("VK error:", err)
			return
		}
	}()
	var chatID, _ = strconv.Atoi(os.Getenv("VK_CHAT_ID"))
	for {
		if len(vk_queue) > 0 {
			var m = vk_queue[0]
			var text = fmt.Sprintf("%s:\n\n%s", m.Author, m.Text)
			_, err := govkbot.API.SendChatMessage(chatID, text)
			if err != nil {
				return
			}
			if len(vk_queue) > 1 {
				vk_queue = append(vk_queue[:1], vk_queue[2:]...)
			} else {
				vk_queue = nil
			}
		}
	}
}

func messageHandler(msg *govkbot.Message) (reply string) {
	VkChatId, _ := strconv.Atoi(os.Getenv("VK_CHAT_ID"))
	me, _ := govkbot.API.Me()
	if msg.ChatID == VkChatId && msg.UserID != me.ID {
		var user, err = govkbot.API.User(msg.UserID)
		if err != nil {
			log.Fatalln("Errror get VK user:", err)
		}
		var message = custom_types.Message{Text: msg.Body, Author: user.FullName(), Attachments: nil}
		tg_queue = append(tg_queue, message)
		err = msg.MarkAsRead()
		if err != nil {
			return ""
		}
	}
	return ""
}
