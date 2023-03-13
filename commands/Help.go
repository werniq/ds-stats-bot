package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/werniq/ds-stats-bot/embed"
)

func (app *application) Help(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == BotPrefix+"help" {
		text := `
			Hey! My name is Семен) and I provide some statistic for discord channels.
			For ideas, please contact creator: Rama(V)#5065
			Commands:
				.stats <-word-> - Gives data of usage of this word
		`
		_, err := s.ChannelMessageSendEmbed(m.ChannelID, embed.CreateEmbedMessage("Here is some information you asked for <3", text, 3).Build())
		if err != nil {
			errorLog.Println("Error sending embed message: %v", err)
		}
	}
}
