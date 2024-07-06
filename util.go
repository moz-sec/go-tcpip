package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"unicode"
)

const (
	ERROR = iota + 1
	WARNING
	INFO
	DEBUG
)

func SetLogLevel() int {
	logLevel := os.Getenv("LOG_LEVEL")
	switch logLevel {
	case "INFO", "info":
		return INFO
	case "DEBUG", "debug":
		return DEBUG
	default:
		return INFO
	}
}

type BuiltinLogger struct {
	logger *log.Logger
	level  int
}

func NewBuiltinLogger() *BuiltinLogger {
	return &BuiltinLogger{
		logger: log.Default(),
		level:  SetLogLevel(),
	}
}

func (l *BuiltinLogger) Debug(format string, args ...interface{}) {
	if l.level >= DEBUG {
		l.logger.SetOutput(os.Stdout)
		l.logger.SetFlags(log.Ltime)
		l.logger.SetPrefix("[D] ")

		_, file, line, ok := runtime.Caller(1)
		if ok {
			caller := fmt.Sprintf(" (%s:%d)", file, line)
			l.logger.Printf(format+caller, args...)
		} else {
			l.logger.Printf(format, args...)
		}
	}
}

func (l *BuiltinLogger) Info(format string, args ...interface{}) {
	if l.level >= INFO {
		prefix := fmt.Sprintf("[%s]", "INFO")
		l.logger.SetOutput(os.Stdout)
		l.logger.SetPrefix(prefix)
		l.logger.SetFlags(log.Ldate | log.Ltime)
		l.logger.Printf(format, args...)
	}
}

func HexDump(fp *os.File, data []byte) {
	fmt.Fprintf(fp, "+------+-------------------------------------------------+------------------+\n")
	for offset := 0; offset < len(data); offset += 16 {
		fmt.Fprintf(fp, "| %04x | ", offset)
		for index := 0; index < 16; index++ {
			if offset+index < len(data) {
				fmt.Fprintf(fp, "%02x ", data[offset+index])
			} else {
				fmt.Fprintf(fp, "   ")
			}
		}
		fmt.Fprintf(fp, "| ")
		for index := 0; index < 16; index++ {
			if offset+index < len(data) {
				ch := data[offset+index]
				if unicode.IsPrint(rune(ch)) {
					fmt.Fprintf(fp, "%c", ch)
				} else {
					fmt.Fprintf(fp, ".")
				}
			} else {
				fmt.Fprintf(fp, " ")
			}
		}
		fmt.Fprintf(fp, " |\n")
	}
	fmt.Fprintf(fp, "+------+-------------------------------------------------+------------------+\n")
}
