package bot

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	handler "github.com/end1essrage/dndhelper-discord/pkg/handler"
	t "github.com/end1essrage/dndhelper-discord/pkg/types"
	"github.com/sirupsen/logrus"
)

const commandLiteral byte = '='

type Bot struct {
	session *discordgo.Session
	handler *handler.Handler
}

func NewBot(token string, handler *handler.Handler) *Bot {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		logrus.Fatalf("error creating Discord session,", err)
	}

	return &Bot{session: session, handler: handler}
}

func (b *Bot) Start() error {
	b.session.AddHandler(b.onMessage)

	// In this example, we only care about receiving message events.
	b.session.Identify.Intents = discordgo.IntentsGuildMessages

	err := b.session.Open()
	if err != nil {
		logrus.Error("error opening connection,", err)
		return err
	}

	return nil
}

func (b *Bot) onMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	//скипает свои сообщения
	if m.Author.ID == s.State.User.ID {
		return
	}
	//скипает все сообщенияб кроме комманд
	if m.Content[0] != commandLiteral {
		return
	}

	command := b.deserializeInput(m.Content)

	message, err := b.handler.Handle(command)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "error executing command")
	}

	s.ChannelMessageSend(m.ChannelID, message)
}

func (b *Bot) deserializeInput(message string) t.BotCommand {
	items := strings.Split(message, " ")
	sCommand := items[0][1:]
	params := make(map[string]string)
	value := ""

	for i := 1; i < len(items); i++ {
		// если есть -, то определять как параметр, иначе это является значением
		if string(items[i][0]) == "-" {
			flags := strings.Split(items[i], string(commandLiteral))
			if len(flags) > 2 {
				//это значит что в одном флаге два знака равно
				continue
			}
			if len(items) < 2 {
				params[flags[0][1:]] = " "
			} else {
				params[flags[0][1:]] = items[1]
			}

		} else {
			value = items[i]
		}
	}

	return t.BotCommand{Command: sCommand, Value: value, Params: params}
}

func (b *Bot) Stop() {
	b.session.Close()
}
