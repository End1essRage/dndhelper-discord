package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	client "github.com/end1essrage/dndhelper-discord/client"

	"github.com/bwmarrin/discordgo"
)

const commandLiteral string = "="

var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {

	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	//регистрируем обработчик
	dg.AddHandler(onMessage)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

func onMessage(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content[0] != commandLiteral[0] { //переделать на байтовую команду
		return
	}

	//разделяем на слова, создаем команду и заполняем переменные

	items := strings.Split(m.Content, " ")
	sCommand := items[0][1:]

	fmt.Print(items[1])

	switch sCommand {
	case "spellinfo":
		s.ChannelMessageSend(m.ChannelID, handleSpellInfoCommand(items[1]))
		break
	case "help":
		break
	default:
		break
	}
}

func handleSpellInfoCommand(spellName string) string {

	var sb strings.Builder

	spell, err := getSpellInfo(spellName)
	if err != nil {
		sb.WriteString("ERROR OCCURED: ")
		sb.WriteString(err.Error())
	}

	sb.WriteString("Spell Name : " + spell.Name)
	sb.WriteString("\n")

	sb.WriteString("Description: ")
	for i := 0; i < len(spell.Desc); i++ {
		sb.WriteString(spell.Desc[i])
		sb.WriteString("\n")
	}

	sb.WriteString("Damage is " + spell.Damage.DamageType.Name + "\n")
	for key, value := range spell.Damage.DamageAtSlotLevel {
		sb.WriteString(key + " - " + value + "\n")
	}

	fmt.Print(sb.String())

	return sb.String()
}

func getSpellInfo(spellName string) (*client.Spell, error) {
	//отправить запрос на апишку
	client := client.NewClient("localhost:8080")
	//десериализовать в объект
	spell, err := client.GetSpellInfo(spellName)
	if err != nil {
		//
	}

	//отправить ответ
	return spell, nil
}
