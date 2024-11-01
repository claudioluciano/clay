package logger

// Default logger instance
var defaultLogger Logger = NewDefaultLogger()

// Level represents the severity level of a log message
type Level int8

const (
	// DebugLevel defines debug log level.
	DebugLevel Level = iota
	// InfoLevel defines info log level.
	InfoLevel
	// WarnLevel defines warn log level.
	WarnLevel
	// ErrorLevel defines error log level.
	ErrorLevel
	// FatalLevel defines fatal log level.
	FatalLevel
	// PanicLevel defines panic log level.
	PanicLevel
	// TraceLevel defines trace log level.
	TraceLevel
)

// Logger defines the interface for logging operations at different severity levels
type Logger interface {
	Debug() LoggerContext
	Info() LoggerContext
	Warn() LoggerContext
	Error() LoggerContext
	Fatal() LoggerContext
	Panic() LoggerContext
	Trace() LoggerContext
	SetLevel(level Level)
	EnableColors(enable bool)
}

// LoggerField represents a key-value pair for structured logging
type LoggerField struct {
	key   string
	value interface{}
}

func (l *LoggerField) Key() string {
	return l.key
}

func (l *LoggerField) Value() interface{} {
	return l.value
}

// Field creates a new LoggerField with the given key and value
func Field(key string, value interface{}) LoggerField {
	return LoggerField{
		key:   key,
		value: value,
	}
}

// LoggerContext defines the interface for building and sending log messages
type LoggerContext interface {
	// Field adds fields to the logging context
	Field(fields ...LoggerField) LoggerContext
	// Msg logs a message
	Msg(msg string)
	// Msgf logs a formatted message
	Msgf(msg string, args ...interface{})
	// Caller adds caller information (file and line) to the log
	Caller() LoggerContext
	// Err adds error information to the log
	Err(err error) LoggerContext
	// Stack adds stack trace to the log
	Stack() LoggerContext
	// Send sends the log entry
	Send()
}

// SetLevel sets the logging level for the default logger
func SetLevel(level Level) {
	defaultLogger.SetLevel(level)
}

// EnableColors enables or disables colors for the default logger
func EnableColors(enable bool) {
	defaultLogger.EnableColors(enable)
}

// SetLogger sets a new default logger
func SetLogger(logger Logger) {
	defaultLogger = logger
}

// Info returns a LoggerContext for info level logging
func Info() LoggerContext {
	return defaultLogger.Info()
}

// Warn returns a LoggerContext for warning level logging
func Warn() LoggerContext {
	return defaultLogger.Warn()
}

// Error returns a LoggerContext for error level logging
func Error() LoggerContext {
	return defaultLogger.Error()
}

// Debug returns a LoggerContext for debug level logging
func Debug() LoggerContext {
	return defaultLogger.Debug()
}

// Trace returns a LoggerContext for trace level logging
func Trace() LoggerContext {
	return defaultLogger.Trace()
}

// Panic returns a LoggerContext for panic level logging
func Panic() LoggerContext {
	return defaultLogger.Panic()
}

// Fatal returns a LoggerContext for fatal level logging
func Fatal() LoggerContext {
	return defaultLogger.Fatal()
}
