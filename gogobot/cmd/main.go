package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main(){
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
    if err != nil {
        log.Panic(err)
    }

    bot.Debug = true
	u := tgbotapi.NewUpdate(0)
    u.Timeout = 60

    updates := bot.GetUpdatesChan(u)

    for update := range updates {
        if update.Message == nil { 
            continue
        }

        if !update.Message.IsCommand() { 
            continue
        }

        msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

        switch update.Message.Command() {
		case "start":
			msg.Text = "Welcome to QuanGo - automatic Go game's score calculation system"
			msg.Text = "You can proceed by"

        case "help":
            msg.Text = "I understand /sayhi and /status."
        case "sayhi":
            msg.Text = "Hi :)"
        case "status":
            msg.Text = "I'm ok."
        default:
            msg.Text = "I don't know that command"
        }

        if _, err := bot.Send(msg); err != nil {
            log.Panic(err)
        }
    }
}