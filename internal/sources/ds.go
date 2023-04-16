package sources

import (
	"context"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"naval/internal/model"
	"strconv"
)

type Source interface {
	Read(ctx context.Context, msgChan chan<- *model.Message)
	Send(msg string, clientID int)
	MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate)
}

type DS struct {
	Chan chan *discordgo.MessageCreate
	bot  *discordgo.Session
}

func NewDs(token string) Source {
	bot, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Ошибка создания Discord клиента: ", err)
	}

	ds := &DS{
		bot: bot,
	}

	ds.Chan = make(chan *discordgo.MessageCreate)

	bot.AddHandler(ds.MessageCreate)
	err = bot.Open()
	if err != nil {
		fmt.Println("Ошибка открытия соединения")
	}
	fmt.Println("Соединение открыто")
	return ds
}

func (ds *DS) MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Проверяем, является ли автор сообщения ботом
	if m.Author.ID == s.State.User.ID {
		return
	}
	ds.Chan <- m
}

func (ds *DS) Read(ctx context.Context, msgChan chan<- *model.Message) {
	for update := range ds.Chan {
		print()
		select {
		case <-ctx.Done():
			return
		default:
		}

		if update.Message != nil { // If we got a message
			id, _ := strconv.Atoi(update.ChannelID)
			msg := &model.Message{
				Text:   update.Content,
				ChatID: id,
			}
			_ = msg

			msgChan <- msg
		}
	}
}

func (ds *DS) Send(msg string, clientID int) {
	Id := strconv.Itoa(clientID)
	_, err := ds.bot.ChannelMessageSend(Id, msg)
	if err != nil {
		log.Printf("%v", err)
	}
}
