package logger

import (
	"log"
	"os"
)

func Logger() *log.Logger {
	return log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
}
