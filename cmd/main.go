package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	c "github.com/end1essrage/dndhelper-discord/pkg"
	bot "github.com/end1essrage/dndhelper-discord/pkg/bot"
	client "github.com/end1essrage/dndhelper-discord/pkg/client"
	handler "github.com/end1essrage/dndhelper-discord/pkg/handler"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	Token string
	Env   string
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	//Environment handling
	Env = os.Getenv("ENVIRONMENT")

	if Env == c.ENV_LOCAL {
		if err := godotenv.Load(); err != nil {
			logrus.Fatalf("error while reading environment %s", err.Error())
		}
	}

	Env = os.Getenv("ENVIRONMENT")

	if Env == "" {
		logrus.Fatal("cant set environment")
	}

	logrus.Info("ENVIRONMENT IS " + Env)

	setToken()
	//Config Handling

	if err := initConfig(); err != nil {
		logrus.Fatalf("error while reading config %s", err.Error())
	}
}

func main() {
	client := client.NewClient(viper.GetString("api_host"))
	handler := handler.NewHandler(client, Env)
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
	v := viper.New()

	if Env == c.ENV_LOCAL {
		v.SetConfigName("config_local")
	}
	if Env == c.ENV_DEV {
		v.SetConfigName("config_pod")
	}

	v.SetConfigType("yml")
	v.AddConfigPath(".")
	v.AddConfigPath("./configs")

	return v.ReadInConfig()
}

func setToken() {
	if Env == c.ENV_LOCAL {
		flag.StringVar(&Token, "t", "", "Bot Token")
		flag.Parse()
	}

	if Env == c.ENV_DEV {
		Token = os.Getenv("TOKEN")
	}
}
