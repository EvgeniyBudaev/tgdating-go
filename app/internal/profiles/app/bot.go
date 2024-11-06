// Package app - module for working with telegram bot
package app

import (
	"context"
	"fmt"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/entity"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
	"time"
)

const (
	EmojiCoin           = "\U0001FA99"
	EmojiPointRight     = "\U0001F449"
	EmojiSmile          = "\U0001F642"
	EmojiSunglasses     = "\U0001F60E"
	UpdateConfigTimeout = 60
	errorFilePathBot    = "internal/app/http.go"
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
func printIntro(chatId int64) {
	printSystemMessageWithDelay(chatId, 1, "Привет! "+EmojiSunglasses)
	printSystemMessageWithDelay(chatId, 5, "Нажми на кнопку App,"+
		" чтобы перейти на главную страницу приложения")
}

// StartBot - launches the bot
func (app *App) StartBot(ctx context.Context, msgChan <-chan *entity.HubContent) error {
	var err error
	// Telegram Bot
	if bot, err = tgbotapi.NewBotAPI(app.config.TelegramBotToken); err != nil {
		return err
	}
	bot.Debug = true
	app.Logger.Info("Authorized on account:", zap.String("username", bot.Self.UserName))
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = UpdateConfigTimeout
	updates := bot.GetUpdatesChan(updateConfig) // Получаем все обновления от пользователя

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case c, ok := <-msgChan:
				if !ok {
					return
				}
				msg := tgbotapi.NewPhoto(int64(c.UserId), tgbotapi.FileURL(c.UserImageUrl))
				msg.ParseMode = "HTML"
				msg.Caption = fmt.Sprintf("%s %s <a href=\"tg://resolve?domain=%s\">@%s</a>",
					c.Message, EmojiPointRight, c.Username, c.Username)
				_, err = bot.Send(msg)
				if err != nil {
					errorMessage := getErrorMessage("StartBot", "bot.Send",
						errorFilePathBot)
					app.Logger.Debug(errorMessage, zap.Error(err))
				}
			}
		}
	}()

	for update := range updates {
		chatId := update.Message.Chat.ID
		if isStartMessage(&update) {
			userText := update.Message.Text // userText - сообщение, которое отправил пользователь
			app.Logger.Info("Начало общения: ", zap.String("username", update.Message.From.UserName),
				zap.String("message", userText))
			printIntro(chatId)
		}

		//if update.Message != nil {
		//	log.Printf("[%s] %s", update.Message.From.UserName, userText)
		//	msg := tgbotapi.NewMessage(chatId, userText)
		//	msg.ReplyToMessageID = update.Message.MessageID
		//	bot.Send(msg)
		//}
	}
	return nil
}
