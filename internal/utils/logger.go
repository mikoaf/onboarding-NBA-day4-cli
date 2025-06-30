package utils

import (
	"log"
	"os"
	"time"
)

type LogMessage struct {
	Time    string
	Tag     string
	Message string
}

type Logger struct {
	ch chan LogMessage
}

// Package logger provide async non blocking logging for better app performance
func NewLogger(bufferSize int) *Logger {
	logger := &Logger{
		ch: make(chan LogMessage, bufferSize), // Buffered channel
	}
	logger.init()
	return logger
}

// Async logging function that runs in a separate goroutine
func (l *Logger) init() {
	log.SetOutput(os.Stdout)
	log.SetFlags(0) // Disable the default timestamp and log prefix
	go func() {
		for logMsg := range l.ch {
			log.Printf("[%s] [%s]: %s", logMsg.Time, logMsg.Tag, logMsg.Message)
		}
	}()
}

// Commit log
func (l *Logger) Add(tag, message string) {
	currentTime := time.Now().Format("15:04:05.000") // Hour:Minute:Second.Millisecond
	go func() {
		l.ch <- LogMessage{Time: currentTime, Tag: tag, Message: message}
	}()
}
