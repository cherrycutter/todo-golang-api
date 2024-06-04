package logger

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
)

var (
	Info  *log.Logger
	Error *log.Logger
)

// InitLogger initializes logger and creates logs in files logs/info.log for info and logs/error.log for errors
func InitLogger() {
	infoLog := &lumberjack.Logger{
		Filename:   "logs/info.log",
		MaxSize:    5, // Megabytes
		MaxBackups: 3,
		MaxAge:     10,   // Days
		Compress:   true, // Compression of old files
	}

	errorLog := &lumberjack.Logger{
		Filename:   "logs/error.log",
		MaxSize:    5, // Megabytes
		MaxBackups: 3,
		MaxAge:     10,   // Days
		Compress:   true, // Compression of old files
	}

	Info = log.New(infoLog, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(errorLog, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
