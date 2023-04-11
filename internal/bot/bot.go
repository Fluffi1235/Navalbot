package bot

import (
	"context"
	"naval/internal/model"
	"naval/internal/service"
	"naval/internal/sources"
	"strings"
	"sync"
)

type Bot struct {
	Source sources.Source
}

func NewBot(Source sources.Source) Bot {
	return Bot{
		Source: Source,
	}
}

func (b *Bot) RunBot(ctx context.Context, wg *sync.WaitGroup) {
	msgChan := make(chan *model.Message)

	go b.Source.Read(ctx, msgChan)

	b.HandlingMessage(msgChan)

	close(msgChan)
	wg.Done()
}

func (b *Bot) HandlingMessage(msgChan <-chan *model.Message) {
	for msg := range msgChan {
		answer := make([]string, 0, 0)
		answer = service.GerInfoDB(strings.ToLower(msg.Text))
		if len(answer) == 0 {
			b.Source.Send("Неверное название предмета", msg.ChatID)
		} else {
			for _, value := range answer {
				b.Source.Send(value, msg.ChatID)
			}
			b.Source.Send("Вот все города где продается "+msg.Text, msg.ChatID)
		}
	}
}
