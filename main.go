package main

import (
	"fmt"
	"os"

	"github.com/steelthedev/go-logger/logger"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: logger <logPath> <message>")
		return
	}
	logPath := os.Args[1]
	logger.Logger(logPath)
}
