package services

import (
	"fmt"
	"github.com/yanzay/tbot/v2"
	"log"
	"mega_chat/custom_types"
	"os"
)

var queue []custom_types.Message
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
		if len(queue) > 0 {
			var m = queue[0]
			var text = fmt.Sprintf("<b>%s</b>:\n\n%s", m.Author, m.Text)
			_, err := client.SendMessage(chatID, text, tbot.OptParseModeHTML)
			if err != nil {
				log.Println("Error send tg message:", err)
			}
			if len(queue) > 1 {
				queue = append(queue[:1], queue[2:]...)
			} else {
				queue = nil
			}
		}
	}
}

func testHandler(msg *tbot.Message) {
	if msg.Chat.ID == os.Getenv("TELEGRAM_CHAT_ID") {
		// TODO: Пересылка сообщения в вк.
	}
}
