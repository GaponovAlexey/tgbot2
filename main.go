package main

import (
	"fmt"
	"log"
	"os"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"

)



func main() {
	err := godotenv.Load() // читать .env
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client()
}

func client() {
	var numericKeyboard = tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonURL("1.com", "http://1.com"),
			tg.NewInlineKeyboardButtonData("2", "2"),
			tg.NewInlineKeyboardButtonData("3", "3"),
		),
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData("4", "4"),
			tg.NewInlineKeyboardButtonData("5", "5"),
			tg.NewInlineKeyboardButtonData("6", "6"),
		),
	)

	bot, err := tg.NewBotAPI(os.Getenv("TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	u := tg.NewUpdate(0)
	u.Timeout = 100

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		text := update.Message.Text
		chatId := update.Message.Chat.ID
		hi := "hi"
		you := "you"
		open := "open"

		
		msg := tg.NewMessage(chatId, text)

		//command
		switch text {
		case hi:
			bot.Send(tg.NewMessage(chatId, "hi my friend"))
		case open:
			msg.ReplyMarkup = numericKeyboard
		case you:
			bot.Send(tg.NewMessage(chatId, "you you you bro"))
		default:
			bot.Send(tg.NewMessage(chatId, "i don't understand you human"))
		}
		///da 

		
	}

}

// func (s Struct) send() {
// 	bot.Send(tg.NewMessage(chatId, s))
// }
