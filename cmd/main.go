package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	bot "github.com/end1essrage/dndhelper-discord/pkg/bot"
	client "github.com/end1essrage/dndhelper-discord/pkg/client"
	handler "github.com/end1essrage/dndhelper-discord/pkg/handler"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	//конфигурируем приложение
	logrus.SetFormatter(&logrus.JSONFormatter{})

	if err := initConfig(); err != nil {
		logrus.Fatalf("error while reading config %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		//non fatal
		logrus.Error("error while reading environment %s", err.Error())
	}

	token := os.Getenv("TOKEN")
	if token != "" {
		Token = token
	}

	client := client.NewClient(viper.GetString("api_host"))
	handler := handler.NewHandler(client)
	bot := bot.NewBot(Token, handler)
	err := bot.Start()
	if err != nil {
		logrus.Fatalf("cant start bot %s", err.Error())
	}

	logrus.Info("Bot is now running.  Press CTRL-C to exit.")
	// Wait here until CTRL-C or other term signal is received.
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	bot.Stop()
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
