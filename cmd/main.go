package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/werniq/ds-stats-bot/commands"
	"github.com/werniq/ds-stats-bot/logger"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		logger.Logger().Printf("Error starting new discord session: %v", err)
		return
	}

	bot := commands.HandleCommands()
	err = bot.Open()
	if err != nil {
		logger.Logger().Printf("Error opening Discord Session: %v\n", err)
		return
	}
	fmt.Println("Bot is currently running. CTRL-C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
