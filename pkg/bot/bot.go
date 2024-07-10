package bot

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	handler "github.com/end1essrage/dndhelper-discord/pkg/handler"
	"github.com/sirupsen/logrus"
)

const commandLiteral string = "="

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

	b.session.AddHandler(b.onMessage) //узнать о других вариантах

	// In this example, we only care about receiving message events.
	b.session.Identify.Intents = discordgo.IntentsGuildMessages

	err := b.session.Open()
	if err != nil {
		logrus.Fatalf("error opening connection,", err)
		return err
	}

	// Wait here until CTRL-C or other term signal is received.
	logrus.Info("Bot is now running.  Press CTRL-C to exit.")
	return nil
}

func (b *Bot) onMessage(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content[0] != commandLiteral[0] { //переделать на байтовую команду
		return
	}

	//разделяем на слова, создаем команду и заполняем переменные

	items := strings.Split(m.Content, " ")
	sCommand := items[0][1:]
	params := make(map[string]string)
	value := ""

	for i := 1; i < len(items); i++ {
		// если есть -, то определять как параметр, иначе это является значением
		if string(items[i][0]) == "-" {
			flags := strings.Split(items[i], "=")
			if len(flags) > 2 {
				//err
				continue
			}
			params[flags[0][1:]] = flags[1]
		} else {
			value = items[i]
		}
	}

	if value == "" {
		s.ChannelMessageSend(m.ChannelID, "error reading command value")
	}

	message, err := b.handler.Handle(sCommand, value, params)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "error occured")
	}
	//Добавить обработку флагов и параметровб подумать над реализацией опредления данных и параметров
	//Можно нарезать строку и вычленивать значение параметров
	// фориат пока -ключ=значение
	s.ChannelMessageSend(m.ChannelID, message)
}

func (b *Bot) Stop() {
	b.session.Close()
}
