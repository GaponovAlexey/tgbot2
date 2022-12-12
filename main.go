package main

import (
	"log"
	"os"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
  "github.com/joho/godotenv"

)

type Struct struct {
	s string
}

func main() {

  client()
}

func client() {
	err := godotenv.Load() // читать .env
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	bot, err := tg.NewBotAPI(os.Getenv("TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	u := tg.NewUpdate(0)
	u.Timeout = 50000
  
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		text := update.Message.Text
		chatId := update.Message.Chat.ID
		hi := "hi"
		you := "you"
		switch text {
		case hi:
			bot.Send(tg.NewMessage(chatId, "hi my friend"))
		case you:
			bot.Send(tg.NewMessage(chatId, "you you you bro"))
		default:
			bot.Send(tg.NewMessage(chatId, "i don't understand you human"))
		}
	}

}

// func (s Struct) send() {
// 	bot.Send(tg.NewMessage(chatId, s))
// }
