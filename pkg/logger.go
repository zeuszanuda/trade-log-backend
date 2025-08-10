package pkg

import (
	"log"
)

type Logger struct {
	env string
}

func NewLogger(env string) *Logger {
	return &Logger{env: env}
}

func (l *Logger) Info(msg string) {
	log.Println("[INFO]", msg)
}

func (l *Logger) Error(msg string) {
	log.Println("[ERROR]", msg)
}
