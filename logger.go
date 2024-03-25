package easyget

import (
	"fmt"
	"io"
	"log"
	"os"
)

const LOG_FILE_VAR = "EASYGET_LOG_FILE"

func NewDefaultLogger(file string) ILogger {
	var writer io.Writer

	f, err := os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		writer = os.Stderr
	} else {
		writer = f
	}

	l := log.New(writer, "[easyget]", log.LstdFlags)

	return &DefaultLogger{l: l}
}

type DefaultLogger struct {
	l *log.Logger
}

func (l *DefaultLogger) Errorf(format string, args ...interface{}) {
	l.l.Printf("[Error] %s\n", fmt.Sprintf(format, args...))
}

func (l *DefaultLogger) Debugf(format string, args ...interface{}) {
	l.l.Printf("[Debug] %s\n", fmt.Sprintf(format, args...))
}

type ILogger interface {
	Errorf(format string, args ...interface{})
	Debugf(format string, args ...interface{})
}

var l ILogger

func init() {
	fp := os.Getenv(LOG_FILE_VAR)
	if fp == "" {
		fp = "/var/log/easyget.log"
	}

	l = NewDefaultLogger(fp)
}

func SetLogger(tl ILogger) {
	l = tl
}
