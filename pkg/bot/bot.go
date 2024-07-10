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

	message, err := b.handler.Handle(sCommand, items[1][:])
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "error occured")
	}
	//Добавить обработку флагов и параметровб подумать над реализацией опредления данных и параметРОВ,,

	s.ChannelMessageSend(m.ChannelID, message)
}

func (b *Bot) Stop() {
	b.session.Close()
}
