package logger

import (
	"log"
	"os"
)

var (
	Warn  *log.Logger
	Info  *log.Logger
	Error *log.Logger
)

func InitLogger() {

	Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warn = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
