package sources

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"naval/internal/model"
	"naval/resources"
)

type Source interface {
	Read(ctx context.Context, msgChan chan<- *model.Message)
	Send(msg string, clientID int64)
}

type TG struct {
	Chan   tgbotapi.UpdatesChannel
	bot    *tgbotapi.BotAPI
	CharID int64
}

func NewTG() Source {
	bot, err := tgbotapi.NewBotAPI(resources.Token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	tg := &TG{
		Chan: updates,
		bot:  bot,
	}

	return tg
}

func (tg *TG) Read(ctx context.Context, msgChan chan<- *model.Message) {
	for update := range tg.Chan {
		print()

		select {
		case <-ctx.Done():
			return
		default:
		}

		if update.Message != nil { // If we got a message
			msg := &model.Message{
				Text:   update.Message.Text,
				ChatID: update.Message.Chat.ID,
			}
			_ = msg

			msgChan <- msg
		}
	}
}

func (tg *TG) Send(msg string, clientID int64) {
	tgMsg := tgbotapi.NewMessage(clientID, msg)

	_, err := tg.bot.Send(tgMsg)
	if err != nil {
		log.Printf("%v", err)
	}
}
