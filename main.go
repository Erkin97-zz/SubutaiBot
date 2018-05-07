package main

import (
	"log"
	"os"

	"github.com/yanzay/tbot"
)

func main() {
	os.Setenv("TELEGRAM_TOKEN", "570124949:AAESKJOMl4Rm9UtmEhw6_lP6RErsHNzYvk0")
	bot, err := tbot.NewServer(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}
	bot.Handle("/answer", "42")
	bot.ListenAndServe()
}
