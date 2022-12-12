package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"

)

var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("1.com", "http://1.com"),
		tgbotapi.NewInlineKeyboardButtonData("2", "2"),
		tgbotapi.NewInlineKeyboardButtonData("3", "3"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("4", "4"),
		tgbotapi.NewInlineKeyboardButtonData("5", "5"),
		tgbotapi.NewInlineKeyboardButtonData("6", "6"),
	),
)

func main() {
	err := godotenv.Load() // читать .env
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client()
}

func client() {

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		text := update.Message.Text
		chatId := update.Message.Chat.ID

		hi := "hi"
		you := "you"

		if update.Message != nil {
			msg := tgbotapi.NewMessage(chatId, update.Message.Text)

			switch text {
			case hi:
				bot.Send(tgbotapi.NewMessage(chatId, "hi my friend"))
			case you:
				bot.Send(tgbotapi.NewMessage(chatId, "end you you"))
			case "open":
				msg.ReplyMarkup = numericKeyboard
				bot.Send(msg)
			default:
				bot.Send(tgbotapi.NewMessage(chatId, "i don't understand you human"))

			}

		} else if update.CallbackQuery != nil {
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := bot.Request(callback); err != nil {
				panic(err)
			}

			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
		}
	}
}
