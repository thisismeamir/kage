package init_methods

import (
	"io"
	"log"
	"os"
)

func InitializeLogger(logFilePath string) {
	// Open the log file (create it if it doesn't exist)
	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}

	// Create a multi-writer to write to both stdout (terminal) and the log file
	multiWriter := io.MultiWriter(os.Stdout, logFile)

	// Set the log output to the multi-writer
	log.SetOutput(multiWriter)

	// Optional: Customize log format (e.g., add date, time)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
