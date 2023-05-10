package pkg

import "github.com/sirupsen/logrus"

type Logger struct {
	Logger *logrus.Logger
}

func NewLogger() *Logger {
	l := logrus.New()
	l.SetFormatter(&logrus.JSONFormatter{})

	return &Logger{
		Logger: l,
	}
}

func (l *Logger) Info(msg string, fields map[string]interface{}) {
	l.Logger.WithFields(fields).Info(msg)
}

func (l *Logger) Warn(msg string, fields map[string]interface{}) {
	l.Logger.WithFields(fields).Warn(msg)
}

func (l *Logger) Error(msg string, fields map[string]interface{}) {
	l.Logger.WithFields(fields).Error(msg)
}

func (l *Logger) Fatal(msg string, fields map[string]interface{}) {
	l.Logger.WithFields(fields).Fatal(msg)
}

func (l *Logger) Panic(msg string, fields map[string]interface{}) {
	l.Logger.WithFields(fields).Panic(msg)
}

func (l *Logger) Debug(msg string, fields map[string]interface{}) {
	l.Logger.WithFields(fields).Debug(msg)
}

func (l *Logger) Trace(msg string, fields map[string]interface{}) {
	l.Logger.WithFields(fields).Trace(msg)
}
