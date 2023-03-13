package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/werniq/ds-stats-bot/logger"
	"time"
)

type Message struct {
	ID        string
	Content   string
	Author    Author
	Timestamp time.Time
}

type Author struct {
	ID       int    `json:"id"`
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

func (app *application) TruncateDatabase(m *discordgo.MessageCreate, s *discordgo.Session) {
	stmt := `truncate table messages`

	row := app.db.DB.QueryRow(stmt)
	if row.Err() != nil {
		logger.Logger().Printf("Error truncating table: %v\n", row.Err())
		return
	}
	fmt.Printf("Truncate successfull\n")
}

func (app *application) SaveMessage(_ *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == "" {
		return
	}

	if m.Author.Bot {
		return
	}

	stmt := `
				INSERT INTO 
				    	messages(message, message_id, author_id, author_username, avatar)
				VALUES ($1, $2, $3, $4, $5)
			`

	row := app.db.DB.QueryRow(stmt, m.Content, m.ID, m.Author.ID, m.Author.Username, m.Author.Avatar)

	if row.Err() != nil {
		logger.Logger().Printf("Error inserting message to table: %v\n", row.Err())
		return
	}
}
