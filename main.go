package main

import (
	"fmt"

	"github.com/Rogue1-cmd/go-discord-bot/bot"
	"github.com/Rogue1-cmd/go-discord-bot/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error)
		return
	}

	bot.Start()

	<-make(chan chan struct{})
	return
}
