package main

import (
	"fmt"
	"lunch-planner-bot/constant"
	"lunch-planner-bot/store"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func MessageHandler(bot *tgbotapi.BotAPI, chatID int64, msg string) error {
	replyMessage := replyMessageFunc(bot, chatID)

	tokens := strings.Split(msg, " ")

	switch tokens[0] {
	case "/add":
		if len(tokens) < 2 {
			return replyMessage(constant.MESSAGE_USAGE)
		}
		restaurant := tokens[1]
		store.Restaurants.Add(restaurant)
		return replyMessage(constant.MESSAGE_RESTAURANT_ADDED + restaurant)
	case "/showall":
		allRestaurants := store.Restaurants.ListAll()
		return replyMessage(constant.MESSAGE_SHOW_ALL_RESTAURANTS + "\n" + strings.Join(allRestaurants, "\n"))
	case "/plan":
		plannedRestaurant, err := store.Restaurants.Plan(5)
		if err != nil {
			return replyMessage(constant.MESSAGE_PLAN_ERROR)
		}
		planned := ([5]string)(plannedRestaurant)
		mondayDate, fridayDate := getCurrentWeek()
		return replyMessage(fmt.Sprintf(
			constant.MESSAGE_LUNCH_PLAN_WEEKLY,
			mondayDate.Format(constant.DEFAULT_DATE_FORMAT),
			fridayDate.Format(constant.DEFAULT_DATE_FORMAT),
			planned[0],
			planned[1],
			planned[2],
			planned[3],
			planned[4],
		))
	default:
		return replyMessage(constant.MESSAGE_USAGE)
	}
}

func replyMessageFunc(bot *tgbotapi.BotAPI, chatID int64) func(string) error {
	return func(msg string) error {
		reply := tgbotapi.NewMessage(chatID, msg)
		if _, err := bot.Send(reply); err != nil {
			return err
		}
		return nil
	}
}

func getCurrentWeek() (monday time.Time, friday time.Time) {
	wd := time.Now().Weekday()
	monday = time.Now()
	if wd == time.Saturday {
		monday = monday.AddDate(0, 0, 2)
	}
	if wd == time.Sunday {
		monday = monday.AddDate(0, 0, 1)
	}
	if wd >= time.Monday && wd <= time.Friday {
		monday = monday.AddDate(0, 0, int(time.Monday-wd))
	}
	friday = monday.AddDate(0, 0, 4)
	return
}
