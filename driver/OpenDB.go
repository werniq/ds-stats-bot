package driver

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/werniq/ds-stats-bot/logger"
	"os"
)

func OpenDb() (*sql.DB, error) {
	DatabaseDsn := os.Getenv("DATABASE_DSN")

	db, err := sql.Open("postgres", DatabaseDsn)
	if err != nil {
		logger.Logger().Printf("Error creating new database connection: %v\n", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		logger.Logger().Printf("Error pinging database connection: %v\n", err)
		return nil, err
	}
	return db, nil
}
