package loggerApp

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"io"
	"os"
)

type Logger struct {
	logFile    string
	permission os.FileMode
	File       *os.File
	Log        *logrus.Logger
}

func NewLogger(logFile string, permission os.FileMode) *Logger {
	return &Logger{
		logFile:    logFile,
		permission: permission,
	}
}

func (l *Logger) InitLogger() error {
	var fileErr error
	l.File, fileErr = os.OpenFile(l.logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, l.permission)
	if fileErr != nil {
		return errors.Wrap(fileErr, "Failed to create log file")
	}

	logrus.SetOutput(l.File)

	l.Log = &logrus.Logger{
		Out:   io.MultiWriter(l.File, os.Stdout),
		Level: logrus.DebugLevel,
		Formatter: &easy.Formatter{
			TimestampFormat: "2006-01-02 15:04:05",
			LogFormat:       "[%lvl%]: %time% - %msg%\n",
		},
	}

	return nil
}
