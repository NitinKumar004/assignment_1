package main

import (
	"fmt"
)

// Logger Interface
type Logger interface {
	Log(message string)
}

// Logger Types
type ConsoleLogger struct{}
type FileLogger struct {
	logs []string
}
type RemoteLogger struct{}

// Implement Log() for each logger type

func (l ConsoleLogger) Log(message string) {
	fmt.Println("Console :", message)
}

func (l *FileLogger) Log(message string) {
	l.logs = append(l.logs, message)
	fmt.Println("File :", message)
}

func (l RemoteLogger) Log(message string) {
	fmt.Println("Remote :", message)
}

// LogAll function
func LogAll(loggers []Logger, message string) {
	for _, logger := range loggers {
		logger.Log(message)
	}
}

// Main
func main() {
	// Create logger instances
	console := ConsoleLogger{}
	file := &FileLogger{} // pointer because method has *FileLogger receiver
	remote := RemoteLogger{}

	// Group all as Logger interface
	loggers := []Logger{console, file, remote}

	// Log message to all
	LogAll(loggers, "hello")
}
