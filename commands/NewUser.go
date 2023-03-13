package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/werniq/ds-stats-bot/embed"
)

func (app *application) NewUser(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Type == 7 {
		s.ChannelMessageSendEmbed(m.ChannelID, embed.CreateEmbedMessage("Good morning!", `
			Wish you a good day! Please, check our rules: 
				Get the fuck out and never come back <3
			`, 4).Build())
	}
}
