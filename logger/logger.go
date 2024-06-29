package logger

import (
	"fmt"
	"log"
	"os"
)

func InitLogger(logPath string) *os.File {
	logFile, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return nil
	}
	defer logFile.Close()

	return logFile
}

func Logger(logPath string) {
	logger := log.New(InitLogger(logPath), "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	if len(os.Args) > 1 {
		message := os.Args[2]
		logger.Println(message)
	} else {
		fmt.Println("No message provided")
	}
}
