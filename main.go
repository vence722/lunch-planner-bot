package main

import (
	"log"
	"lunch-planner-bot/constant"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(constant.TG_BOT_API_KEY)
	if err != nil {
		panic(err)
	}
	bot.Debug = false

	log.Printf("Bot started on account %s\n", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			if err := MessageHandler(bot, update.Message.From.ID, update.Message.Text); err != nil {
				log.Println("[ERROR] handler message error: " + err.Error())
			}
		}
	}
}
