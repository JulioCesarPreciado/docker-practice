package logger

// Package logger provides a simple setup function to redirect log output to a file.

import (
	"log"
	"os"
)

// Setup initializes logging to a specified file path. If the file does not exist,
// it will be created. Logs will be appended if the file already exists.
// It terminates the application if the file cannot be opened.
func Setup(logPath string) {
	f, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("❌ Could not open log file: %v", err)
	}
	log.SetOutput(f)
}

// Info writes a formatted informational message to the configured log file.
// This function assumes Setup has already been called to direct output to a file.
func Info(message string) {
	log.Println(message)
}

// Warn writes a formatted warning message to the configured log file.
func Warn(message string) {
	log.Println("[WARN]", message)
}

// Error writes a formatted error message to the configured log file.
func Error(message string) {
	log.Println("[ERROR]", message)
}

// Debug writes a formatted debug message to the configured log file.
func Debug(message string) {
	log.Println("[DEBUG]", message)
}
