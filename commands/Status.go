package commands

import "github.com/bwmarrin/discordgo"

func (app *application) ChangeStatus(s *discordgo.Session, event *discordgo.Event) {
	s.UpdateListeningStatus("Bullshit")
}
