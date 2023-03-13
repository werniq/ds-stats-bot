package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/werniq/ds-stats-bot/embed"
	"github.com/werniq/ds-stats-bot/logger"
	"strings"
)

func (app *application) Stats(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	if m.Content == "" {
		return
	}

	args := strings.Split(strings.TrimPrefix(m.Content, "."), " ")
	command := args[0]

	args = args[1:]
	var statistic int
	var userId string
	var err error

	if command == "stats" {
		statistic, userId, err = app.db.GetStatsForWord(args[0])
		if err != nil {
			logger.Logger().Printf("Error getting statistic for word: %v\n", err)
			return
		}
		fmt.Printf("User %s have typed word %s %d times\n", userId, args[0], statistic)
		user, err := s.User(userId)
		if err != nil {
			logger.Logger().Printf("Error retrieving user from session: %v\n", err)
			return
		}
		s.ChannelMessageSendEmbed(m.ChannelID, embed.CreateEmbedMessage(
			fmt.Sprintf("Statis for usage of word: %s", args[0]), fmt.Sprintf(`
				User %v used the word %s for %d times!
			`,
				user.Mention(),
				args[0],
				statistic),
			1).Build())
	}
}
