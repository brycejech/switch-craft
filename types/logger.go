package types

import (
	"encoding/json"
	"fmt"
	"time"
)

type LogLevel int

const (
	_               LogLevel = iota // 0
	LogLevelError                   // 1
	LogLevelWarning                 // 2
	LogLevelInfo                    // 3
)

type Logger struct {
	LogLevel LogLevel
}

func NewLogger(level LogLevel) *Logger {
	return &Logger{level}
}

type logMessage struct {
	Level      string    `json:"level"`
	TraceID    string    `json:"traceId"`
	TraceStart time.Time `json:"traceStart"`
	Time       time.Time `json:"time"`
	DurationMS int       `json:"durationMs"`
	Message    string    `json:"message"`
	Data       any       `json:"data"`
}

func newLogMessage(
	level string,
	trace OperationTracer,
	message string,
	data any,
) logMessage {
	var traceStart time.Time
	if trace.StartTime.IsZero() {
		traceStart = time.Now()
	} else {
		traceStart = trace.StartTime
	}

	now := time.Now()
	durationMS := int(now.Sub(traceStart) / time.Millisecond)

	return logMessage{
		Level:      level,
		TraceID:    trace.TraceID,
		TraceStart: traceStart,
		Time:       now,
		DurationMS: durationMS,
		Message:    message,
		Data:       data,
	}
}

func (l *Logger) Error(trace OperationTracer, message string, data any) {
	if l.LogLevel < 1 {
		return
	}

	printJSON(newLogMessage("ERROR", trace, message, data))
}

func (l *Logger) Warn(trace OperationTracer, message string, data any) {
	if l.LogLevel < 2 {
		return
	}

	printJSON(newLogMessage("WARNING", trace, message, data))
}

func (l *Logger) Info(trace OperationTracer, message string, data any) {
	if l.LogLevel < 3 {
		return
	}

	printJSON(newLogMessage("INFO", trace, message, data))
}

func printJSON(v interface{}) {
	bytes, _ := json.Marshal(v)
	fmt.Println(string(bytes))
}
