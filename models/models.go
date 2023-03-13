package models

import (
	"database/sql"
	"github.com/werniq/ds-stats-bot/logger"
	"strings"
	"time"
)

type Message struct {
	ID        string    `json:"id"`
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
	Username  string    `json:"username"`
	UserID    string    `json:"user_id"`
	Avatar    string    `json:"avatar"`
}

type DbModel struct {
	DB *sql.DB
}

func (db *DbModel) GetStatsForWord(word string) (int, string, error) {
	statistic := make(map[string]int)
	var ids []string

	stmt := `select * from messages`

	rows, err := db.DB.Query(stmt)
	if err != nil {
		logger.Logger().Printf("Error querying statement: %v\n", err)
		return 0, "", err
	}

	for rows.Next() {
		var id int
		var message string
		var messageID string
		var author_username string
		var author_id string
		var avatar string
		if err = rows.Scan(&id, &message, &messageID, &author_id, &author_username, &avatar); err != nil {
			logger.Logger().Printf("Error scanning values: %v\n", err)
			return 0, "", err
		}

		if strings.Contains(message, word) {
			statistic[author_id]++
			ids = append(ids, author_id)
		}
	}
	max := statistic[ids[0]]
	user_id := ids[0]
	for i := 0; i <= len(ids)-1; i++ {
		if statistic[ids[i]] > max {
			max = statistic[ids[i]]
			user_id = ids[i]
		}
	}

	return max, user_id, nil
}
