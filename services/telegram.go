package services

import (
	"fmt"
	"github.com/yanzay/tbot/v2"
	"log"
	"mega_chat/custom_types"
	"os"
)

var tg_queue []custom_types.Message
var client *tbot.Client

func CreateTelegramBot(token string) {
	bot := tbot.New(token)
	client = bot.Client()
	bot.HandleMessage("", testHandler)

	go func() {
		err := bot.Start()
		if err != nil {
			log.Fatalln("Telegram error:", err)
			return
		}
	}()
	var chatID = os.Getenv("TELEGRAM_CHAT_ID")
	for {
		if len(tg_queue) > 0 {
			var m = tg_queue[0]
			var text = fmt.Sprintf("<b>%s</b>:\n\n%s", m.Author, m.Text)
			_, err := client.SendMessage(chatID, text, tbot.OptParseModeHTML)
			if err != nil {
				log.Println("Error send tg message:", err)
			}
			if len(tg_queue) > 1 {
				tg_queue = append(tg_queue[:1], tg_queue[2:]...)
			} else {
				tg_queue = nil
			}
		}
	}
}

func testHandler(msg *tbot.Message) {
	if msg.Chat.ID == os.Getenv("TELEGRAM_CHAT_ID") {
		var userFullname = msg.From.FirstName + " " + msg.From.LastName
		var m = custom_types.Message{Text: msg.Text, Author: userFullname}
		// var m = custom_types.Message{Text: msg.Text, Author: "idk", Attachments: nil}
		vk_queue = append(vk_queue, m)
	}
}
