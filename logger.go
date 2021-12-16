package logger

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type LogLevel int

// LogLevelError.
// Error level should used in error condition happen.
const LogLevelError LogLevel = 0

// LogLevelWarn.
// Warn level should used when something is go strange. Add a warn log for
// tracing the possible problem
const LogLevelWarn LogLevel = 1

// LogLevelInfo
// Info level should used only for simple flow check.
const LogLevelInfo LogLevel = 2

// LogLevelDebug
// Debug level should used for flow in detail, or key value when debugging
const LogLevelDebug LogLevel = 3

// LogLevelTrace
// Trace level can used in detail value or every step in action for remote debugging
const LogLevelTrace LogLevel = 4

type OutputType int

const OutputTypeStdout OutputType = 0
const OutputTypeStderr OutputType = 1

// String will get wording of level
func (level LogLevel) String() string {
	switch level {
	case LogLevelError:
		return "ERROR"
	case LogLevelWarn:
		return "WARN "
	case LogLevelInfo:
		return "INFO "
	case LogLevelDebug:
		return "DEBUG"
	case LogLevelTrace:
		fallthrough
	default:
		return "TRACE"
	}
}

func (level LogLevel) GetLogger() func(string, ...interface{}) (int, error) {
	switch level {
	case LogLevelError:
		return Error
	case LogLevelWarn:
		return Warn
	case LogLevelInfo:
		return Info
	case LogLevelDebug:
		return Debug
	case LogLevelTrace:
		fallthrough
	default:
		return Trace
	}
}

var loggerLevel LogLevel = LogLevelWarn
var outputType OutputType = OutputTypeStdout

// init will load level from env, if there has no any valid setting, default will be warning
func init() {
	loadLevelFromEnv()
}

func loadLevelFromEnv() {
	envLevelStr := os.Getenv("BATH_LOGGER_LEVEL")
	if level, err := strconv.ParseInt(envLevelStr, 10, 32); err == nil {
		if level >= int64(LogLevelError) && level <= int64(LogLevelTrace) {
			loggerLevel = LogLevel(level)
		}
	}
}

// SetLogLevel let user can change level
func SetLogLevel(level LogLevel) {
	loggerLevel = level
}

func SetUsingOutput(output OutputType) {
	outputType = output
}

// getCurrentTime will return time in RFC3339 (2006-01-02T15:04:05Z07:00)
func getCurrentTime() string {
	now := time.Now()
	return now.Format(time.RFC3339)
}

// writeLog will write log in format [level][time] log
func writeLog(level LogLevel, format string, args ...interface{}) (n int, err error) {
	if level <= loggerLevel {
		log := fmt.Sprintf(format, args...)
		channel := os.Stdout
		if outputType == OutputTypeStderr {
			channel = os.Stderr
		}
		return fmt.Fprintf(channel, "[%s][%s] %s\n", level.String(), getCurrentTime(), log)
	}
	return 0, nil
}

// Error will write log if level is equal or higher then LogLevelError
func Error(format string, args ...interface{}) (n int, err error) {
	return writeLog(LogLevelError, format, args...)
}

// Warn will write log if level is equal or higher then LogLevelWarn
func Warn(format string, args ...interface{}) (n int, err error) {
	return writeLog(LogLevelWarn, format, args...)
}

// Info will write log if level is equal or higher then LogLevelInfo
func Info(format string, args ...interface{}) (n int, err error) {
	return writeLog(LogLevelInfo, format, args...)
}

// Debug will write log if level is equal or higher then LogLevelDebug
func Debug(format string, args ...interface{}) (n int, err error) {
	return writeLog(LogLevelDebug, format, args...)
}

// Trace will write log if level is equal or higher then LogLevelTrace
func Trace(format string, args ...interface{}) (n int, err error) {
	return writeLog(LogLevelTrace, format, args...)
}
