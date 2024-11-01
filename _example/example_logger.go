package _example

import (
	"fmt"

	"github.com/leap-fish/clay/pkg/logger"
	"github.com/sirupsen/logrus"
)

type MyLogger struct {
	log *logrus.Logger
}

type MyLoggerContext struct {
	level  logger.Level
	log    *logrus.Logger
	fields logrus.Fields
	err    error
}

func (l *MyLoggerContext) Field(fields ...logger.LoggerField) logger.LoggerContext {
	if l.fields == nil {
		l.fields = make(logrus.Fields)
	}
	for _, field := range fields {
		l.fields[field.Key()] = field.Value()
	}
	return l
}

func (l *MyLoggerContext) Msg(msg string) {
	entry := l.log.WithFields(l.fields)
	if l.err != nil {
		entry = entry.WithError(l.err)
	}

	switch l.level {
	case logger.DebugLevel:
		entry.Debug(msg)
	case logger.InfoLevel:
		entry.Info(msg)
	case logger.WarnLevel:
		entry.Warn(msg)
	case logger.ErrorLevel:
		entry.Error(msg)
	case logger.FatalLevel:
		entry.Fatal(msg)
	case logger.PanicLevel:
		entry.Panic(msg)
	case logger.TraceLevel:
		entry.Trace(msg)
	default:
		entry.Info(msg)
	}
}

func (l *MyLoggerContext) Msgf(msg string, args ...interface{}) {
	l.Msg(fmt.Sprintf(msg, args...))
}

func (l *MyLoggerContext) Caller() logger.LoggerContext {
	if l.fields == nil {
		l.fields = make(logrus.Fields)
	}
	l.log.SetReportCaller(true)
	return l
}

func (l *MyLoggerContext) Err(err error) logger.LoggerContext {
	l.err = err
	return l
}

func (l *MyLoggerContext) Stack() logger.LoggerContext {
	return l
}

func (l *MyLoggerContext) Send() {
	l.Msg("")
}

func (l *MyLogger) Debug() logger.LoggerContext {
	return &MyLoggerContext{
		level:  logger.DebugLevel,
		log:    l.log,
		fields: make(logrus.Fields),
	}
}

func (l *MyLogger) Info() logger.LoggerContext {
	return &MyLoggerContext{
		level:  logger.InfoLevel,
		log:    l.log,
		fields: make(logrus.Fields),
	}
}

func (l *MyLogger) Warn() logger.LoggerContext {
	return &MyLoggerContext{
		level:  logger.WarnLevel,
		log:    l.log,
		fields: make(logrus.Fields),
	}
}

func (l *MyLogger) Error() logger.LoggerContext {
	return &MyLoggerContext{
		level:  logger.ErrorLevel,
		log:    l.log,
		fields: make(logrus.Fields),
	}
}

func (l *MyLogger) Fatal() logger.LoggerContext {
	return &MyLoggerContext{
		level:  logger.FatalLevel,
		log:    l.log,
		fields: make(logrus.Fields),
	}
}

func (l *MyLogger) Panic() logger.LoggerContext {
	return &MyLoggerContext{
		level:  logger.PanicLevel,
		log:    l.log,
		fields: make(logrus.Fields),
	}
}

func (l *MyLogger) Trace() logger.LoggerContext {
	return &MyLoggerContext{
		level:  logger.TraceLevel,
		log:    l.log,
		fields: make(logrus.Fields),
	}
}

func logrusLevel(level logger.Level) logrus.Level {
	switch level {
	case logger.DebugLevel:
		return logrus.DebugLevel
	case logger.InfoLevel:
		return logrus.InfoLevel
	case logger.WarnLevel:
		return logrus.WarnLevel
	case logger.ErrorLevel:
		return logrus.ErrorLevel
	case logger.FatalLevel:
		return logrus.FatalLevel
	case logger.PanicLevel:
		return logrus.PanicLevel
	case logger.TraceLevel:
		return logrus.TraceLevel
	default:
		return logrus.InfoLevel
	}
}

func (l *MyLogger) SetLevel(level logger.Level) {
	l.log.SetLevel(logrusLevel(level))
}

func (l *MyLogger) EnableColors(enable bool) {
	l.log.SetFormatter(&logrus.TextFormatter{
		ForceColors: enable,
	})
}

func NewMyLogger() *MyLogger {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	return &MyLogger{
		log: log,
	}
}
