package main

import (
	"github.com/joho/godotenv"
	"log"
	"mega_chat/services"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("Error:", err)
	}

	go services.CreateTelegramBot(os.Getenv("TELEGRAM_TOKEN"))
	var VkToken = os.Getenv("VK_TOKEN")
	services.CreateVkBot(VkToken)
}
