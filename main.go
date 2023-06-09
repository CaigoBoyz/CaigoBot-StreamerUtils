package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/caigoboyz/caigobot/streamerutils/bot"
	"github.com/caigoboyz/caigobot/streamerutils/commands"
)

func main() {
	if err := bot.Session.Open(); err != nil {
		log.Fatalf("Cannot enstablish connection: %e", err)
	}
	commands.RegisterCommands(bot.Session)

	defer bot.Session.Close()

	fmt.Println("Bot is running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	if commands.RemoveCommandsCfg {
		commands.RemoveCommands(bot.Session)
	}

	fmt.Println("Gracefully closed the application.")
	os.Exit(0)
}
