// Package app - module for working with telegram
package app

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/telegram/entity"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
	"strconv"
	"time"
)

const (
	EmojiCoin           = "\U0001FA99"
	EmojiPointRight     = "\U0001F449"
	EmojiSmile          = "\U0001F642"
	EmojiSunglasses     = "\U0001F60E"
	UpdateConfigTimeout = 60
	errorFilePathBot    = "internal/telegram/app/bot.go"
)

// bot - telegram bot
var bot *tgbotapi.BotAPI

// isStartMessage - checks that the /start message has been sent
func isStartMessage(update *tgbotapi.Update) bool {
	return update.Message != nil && update.Message.Text == "/start"
}

// delay - delay
func delay(seconds uint8) {
	time.Sleep(time.Duration(seconds) * time.Second)
}

// printSystemMessageWithDelay - displays a system message with a delay
func printSystemMessageWithDelay(chatId int64, delayInSec uint8, message string) {
	bot.Send(tgbotapi.NewMessage(chatId, message))
	delay(delayInSec)
}

// printIntro - displays a welcome message
func printIntro(chatId int64, languageCode string) {
	var welcomeMessage string
	var instructionMessage string
	instructionMessage, welcomeMessage = translationsPrintIntro(instructionMessage, languageCode, welcomeMessage)
	printSystemMessageWithDelay(chatId, 1, welcomeMessage)
	printSystemMessageWithDelay(chatId, 5, instructionMessage)
}

// StartBot - launches the telegram
func (app *App) StartBot(ctx context.Context) error {
	var err error
	// Telegram Bot
	if bot, err = tgbotapi.NewBotAPI(app.config.TelegramBotToken); err != nil {
		return err
	}
	bot.Debug = true
	app.Logger.Info("Starting Telegram service")
	app.Logger.Info("Authorized on account:", zap.String("username", bot.Self.UserName))

	// удаляет Webhook
	// https://github.com/go-telegram-bot-api/telegram-bot-api/issues/656
	// _, err = bot.MakeRequest("deleteWebhook", tgbotapi.Params{})

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = UpdateConfigTimeout
	updates := bot.GetUpdatesChan(updateConfig) // Получаем все обновления от пользователя

	var hc *entity.HubContent

	go func() {
		for {
			// Kafka
			c, err := app.kafkaReader.ReadMessage(context.Background())
			if err != nil {
				errorMessage := getErrorMessage("StartBot", "r.ReadMessage",
					errorFilePathBot)
				app.Logger.Debug(errorMessage, zap.Error(err))
				break
			}
			err = json.Unmarshal(c.Value, &hc)
			if err != nil {
				errorMessage := getErrorMessage("StartBot", "json.Unmarshal",
					errorFilePathBot)
				app.Logger.Error(errorMessage, zap.Error(err))
				continue
			}
			likedTelegramUserId, err := strconv.ParseInt(hc.LikedTelegramUserId, 10, 64)
			if err != nil {
				errorMessage := getErrorMessage("StartBot", "strconv.ParseInt",
					errorFilePathBot)
				app.Logger.Debug(errorMessage, zap.Error(err))
				break
			}
			msg := tgbotapi.NewPhoto(likedTelegramUserId, tgbotapi.FileURL(hc.UserImageUrl))
			msg.ParseMode = "HTML"
			msg.Caption = fmt.Sprintf("%s %s <a href=\"tg://resolve?domain=%s\">@%s</a>",
				hc.Message, EmojiPointRight, hc.Username, hc.Username)
			_, err = bot.Send(msg)
			if err != nil {
				errorMessage := getErrorMessage("StartBot", "telegram.Send",
					errorFilePathBot)
				app.Logger.Debug(errorMessage, zap.Error(err))
			}
		}
	}()

	go func() {
		for update := range updates {
			chatId := update.Message.Chat.ID
			userLanguageCode := update.Message.From.LanguageCode
			if isStartMessage(&update) {
				//userText := update.Message.Text // userText - сообщение, которое отправил пользователь
				//app.Logger.Info("Начало общения: ", zap.String("username", update.Message.From.UserName),
				//	zap.String("message", userText))
				printIntro(chatId, userLanguageCode)
			}
		}
	}()
	return nil
}
