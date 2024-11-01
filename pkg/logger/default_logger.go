package logger

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/rs/zerolog"
)

type logger struct {
	log zerolog.Logger
}

type loggerContext struct {
	event    *zerolog.Event
	msg      string
	formated bool
}

func (l *loggerContext) Field(fields ...LoggerField) LoggerContext {
	for _, field := range fields {
		l.event.Any(field.key, field.value)
	}

	return l
}

func (l *loggerContext) Msg(msg string) {
	l.event.Msg(msg)
}

func (l *loggerContext) Msgf(msg string, args ...interface{}) {
	l.event.Msgf(msg, args...)
}

func (l *loggerContext) Caller() LoggerContext {
	l.event = l.event.Caller(1)
	return l
}

func (l *loggerContext) Err(err error) LoggerContext {
	l.event = l.event.Err(err)
	return l
}

func (l *loggerContext) Stack() LoggerContext {
	l.event = l.event.Stack()
	return l
}

func (l *loggerContext) Send() {
	if l.msg != "" {
		if l.formated {
			l.event.Msgf(l.msg, l.msg)
			return
		}

		l.event.Msg(l.msg)
	}
	l.event.Send()
}

func (l *logger) Info() LoggerContext {
	return &loggerContext{
		event: l.log.Info(),
	}
}

func (l *logger) Warn() LoggerContext {
	return &loggerContext{
		event: l.log.Warn(),
	}
}

func (l *logger) Error() LoggerContext {
	return &loggerContext{
		event: l.log.Error(),
	}
}

func (l *logger) Debug() LoggerContext {
	return &loggerContext{
		event: l.log.Debug(),
	}
}

func (l *logger) Trace() LoggerContext {
	return &loggerContext{
		event: l.log.Trace(),
	}
}

func (l *logger) Panic() LoggerContext {
	return &loggerContext{
		event: l.log.Panic(),
	}
}

func (l *logger) Fatal() LoggerContext {
	return &loggerContext{
		event: l.log.Fatal(),
	}
}

func zeroLevel(level Level) zerolog.Level {
	switch level {
	case DebugLevel:
		return zerolog.DebugLevel
	case InfoLevel:
		return zerolog.InfoLevel
	case WarnLevel:
		return zerolog.WarnLevel
	case ErrorLevel:
		return zerolog.ErrorLevel
	case FatalLevel:
		return zerolog.FatalLevel
	case PanicLevel:
		return zerolog.PanicLevel
	case TraceLevel:
		return zerolog.TraceLevel
	default:
		return zerolog.InfoLevel
	}
}

func (l *logger) SetLevel(level Level) {
	l.log.Level(zeroLevel(level))
}

func (l *logger) EnableColors(enable bool) {
	l.log = l.log.Output(zerolog.ConsoleWriter{Out: os.Stdout, NoColor: !enable})
}

func NewDefaultLogger() *logger {
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		return filepath.Base(file) + ":" + strconv.Itoa(line)
	}

	return &logger{
		log: zerolog.New(
			zerolog.ConsoleWriter{Out: os.Stdout, NoColor: true},
		).With().Timestamp().Logger(),
	}
}
