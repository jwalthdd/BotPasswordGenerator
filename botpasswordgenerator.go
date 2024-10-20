package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

type App struct {
	bot *tgbotapi.BotAPI
}

const (
	min_length = 1
	max_length = 100
)

func NewBot(token string) (*App, error) {
	bot, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		return nil, err
	}

	bot.Debug = false

	return &App{bot: bot}, nil
}

func (app *App) Start() error {

	log.Printf("Authorized on account %s", app.bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := app.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			length, err := strconv.Atoi(update.Message.Text)

			// Non-numeric or outside the allowed range
			if err != nil || length < min_length || length > max_length {
				continue
			}

			password := generatePassword(length)

			log.Printf("[%s][%s] [ %s ]", update.Message.From.FirstName, update.Message.Text, password)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, password)
			msg.ReplyToMessageID = update.Message.MessageID
			app.bot.Send(msg)
		}
	}

	return nil
}

func (app *App) Debug(debug bool) {
	app.bot.Debug = debug
}
