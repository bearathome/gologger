package gologger

import (
	"os"
	"testing"
)

func TestSetLevel(t *testing.T) {
	t.Log("Testing default log level")
	if loggerLevel != LogLevelWarn {
		t.Error("Default level should be warn")
	}

	t.Log("Testing changing log level")
	SetLogLevel(LogLevelDebug)
	if loggerLevel != LogLevelDebug {
		t.Error("Change log level fail")
	}

	os.Setenv("BATH_LOGGER_LEVEL", "0")
	loadLevelFromEnv()
	if loggerLevel != LogLevelError {
		t.Error("Load level from env fail")
	}

	os.Setenv("BATH_LOGGER_LEVEL", "1")
	loadLevelFromEnv()
	if loggerLevel != LogLevelWarn {
		t.Error("Load level from env fail")
	}

	os.Setenv("BATH_LOGGER_LEVEL", "2")
	loadLevelFromEnv()
	if loggerLevel != LogLevelInfo {
		t.Error("Load level from env fail")
	}

	os.Setenv("BATH_LOGGER_LEVEL", "3")
	loadLevelFromEnv()
	if loggerLevel != LogLevelDebug {
		t.Error("Load level from env fail")
	}

	os.Setenv("BATH_LOGGER_LEVEL", "4")
	loadLevelFromEnv()
	if loggerLevel != LogLevelTrace {
		t.Error("Load level from env fail")
	}
}

func checkLogShouldWrite(t *testing.T, level LogLevel) {
	logger := level.GetLogger()
	n, _ := logger("testing")
	if n == 0 {
		t.Errorf("Log in [%s] level should be written", level.String())
	}
}

func checkLogShouldNotWrite(t *testing.T, level LogLevel) {
	logger := level.GetLogger()
	n, _ := logger("testing")
	if n != 0 {
		t.Errorf("Log in [%s] level should not be written", level.String())
	}
}

func TestLevelErrorLog(t *testing.T) {
	SetLogLevel(LogLevelError)
	checkLogShouldWrite(t, LogLevelError)
	checkLogShouldNotWrite(t, LogLevelWarn)
	checkLogShouldNotWrite(t, LogLevelInfo)
	checkLogShouldNotWrite(t, LogLevelDebug)
	checkLogShouldNotWrite(t, LogLevelTrace)
}

func TestLevelWarnLog(t *testing.T) {
	SetLogLevel(LogLevelWarn)
	checkLogShouldWrite(t, LogLevelError)
	checkLogShouldWrite(t, LogLevelWarn)
	checkLogShouldNotWrite(t, LogLevelInfo)
	checkLogShouldNotWrite(t, LogLevelDebug)
	checkLogShouldNotWrite(t, LogLevelTrace)
}

func TestLevelInfoLog(t *testing.T) {
	SetLogLevel(LogLevelInfo)
	checkLogShouldWrite(t, LogLevelError)
	checkLogShouldWrite(t, LogLevelWarn)
	checkLogShouldWrite(t, LogLevelInfo)
	checkLogShouldNotWrite(t, LogLevelDebug)
	checkLogShouldNotWrite(t, LogLevelTrace)
}

func TestLevelDebugLog(t *testing.T) {
	SetLogLevel(LogLevelDebug)
	checkLogShouldWrite(t, LogLevelError)
	checkLogShouldWrite(t, LogLevelWarn)
	checkLogShouldWrite(t, LogLevelInfo)
	checkLogShouldWrite(t, LogLevelDebug)
	checkLogShouldNotWrite(t, LogLevelTrace)
}

func TestLevelTraceLog(t *testing.T) {
	SetLogLevel(LogLevelTrace)
	checkLogShouldWrite(t, LogLevelError)
	checkLogShouldWrite(t, LogLevelWarn)
	checkLogShouldWrite(t, LogLevelInfo)
	checkLogShouldWrite(t, LogLevelDebug)
	checkLogShouldWrite(t, LogLevelTrace)
}

func BenchmarkRealLog(b *testing.B) {
	SetLogLevel(LogLevelError)
	// testing performance with log is written
	for i := 0; i < b.N; i++ {
		Error("testing")
	}
}
func BenchmarkLogNotWritten(b *testing.B) {
	SetLogLevel(LogLevelError)
	// testing performance with log is written
	for i := 0; i < b.N; i++ {
		Warn("testing")
	}
}
