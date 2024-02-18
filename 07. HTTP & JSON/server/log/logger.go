package log

import (
	"log"
	"os"
)

type ServerLogger struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
	warnLogger  *log.Logger
}

func NewLogger() *ServerLogger {
	return &ServerLogger{
		infoLogger:  log.New(os.Stdout, "[Info]: ", log.Ldate|log.Ltime),
		errorLogger: log.New(os.Stderr, "[Error]: ", log.Ldate|log.Ltime),
		warnLogger:  log.New(os.Stdout, "[Warn]: ", log.Ldate|log.Ltime),
	}
}

func (l *ServerLogger) WriteInfo(text string) {
	l.infoLogger.Println(text)
}

func (l *ServerLogger) WriteWarning(text string) {
	l.warnLogger.Println(text)
}

func (l *ServerLogger) WriteErrorString(text string) {
	l.warnLogger.Println(text)
}
func (l *ServerLogger) WriteErrorObject(error error) {
	l.errorLogger.Println(error.Error())
}
