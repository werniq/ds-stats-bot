package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/werniq/ds-stats-bot/driver"
	"github.com/werniq/ds-stats-bot/logger"
	"github.com/werniq/ds-stats-bot/models"
	"log"
	"os"
)

type application struct {
	bot *discordgo.Session
	db  *models.DbModel
}

var (
	BotPrefix = "."
	errorLog  = log.New(os.Stdout, "ERROR\t", log.Lshortfile|log.Ldate|log.Ltime)
)

func HandleCommands() *discordgo.Session {
	var app application
	bot, err := discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))

	if err != nil {
		logger.Logger().Printf("Error creating new discordgo session: %v\n", err)
		return nil
	}

	app.bot = bot
	db, err := driver.OpenDb()
	if err != nil {
		logger.Logger().Printf("Error opening database connection: %v", err)
		return nil
	}

	app.db = &models.DbModel{
		DB: db,
	}

	bot.Identify.Intents = discordgo.IntentsGuildMessages
	bot.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsGuildMembers | discordgo.IntentsGuildPresences

	// saves each message to database
	bot.AddHandler(app.SaveMessage)
	bot.AddHandler(app.Help)
	bot.AddHandler(app.NewUser)
	bot.AddHandler(app.Stats)
	bot.AddHandler(app.ChangeStatus)

	return bot
}
