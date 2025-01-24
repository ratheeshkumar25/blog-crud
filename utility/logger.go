package utility

import (
	"log"
	"os"
)

func InitLogger() *log.Logger {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}

	logger := log.New(file, "LOG: ", log.Ldate|log.Ltime|log.Lshortfile)
	return logger
}
