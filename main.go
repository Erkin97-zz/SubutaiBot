package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/yanzay/tbot"
)

func main() {
	bot, err := tbot.NewServer(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}
	bot.HandleFunc("/start", startHandler)
	bot.Handle("/answer", "Erkin is our master")
	bot.HandleFunc("/timer {seconds}", timerHandler)
	bot.ListenAndServe()
}

func timerHandler(m *tbot.Message) {
	log.Printf("Handling timer")
	// m.Vars contains all variables, parsed during routing
	secondsStr := m.Vars["seconds"]
	// Convert string variable to integer seconds value
	seconds, err := strconv.Atoi(secondsStr)
	if err != nil {
		m.Reply("Invalid number of seconds")
		return
	}
	m.Replyf("Timer for %d seconds started", seconds)
	time.Sleep(time.Duration(seconds) * time.Second)
	m.Reply("Time out!")
}

func startHandler(m *tbot.Message) {
	buttons := [][]string{
		{"/answer", "/help"},
	}
	m.ReplyKeyboard("Kadir loh", buttons)
}
