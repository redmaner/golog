package golog

import (
	"fmt"
	"io"
	"strings"
	"sync"
	"time"
)

// Logger defines a logger instance
type Logger struct {
	mu    sync.RWMutex
	name  string
	level int
	out   io.Writer
}

func NewLogger(out io.Writer, name string, level string) *Logger {
	l := &Logger{
		name: name,
		out:  out,
	}

	l.SetLevel(level)
	return l
}

func (l *Logger) Log(level int, msg interface{}) (n int, err error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if level <= l.level && msg != nil {
		n, err = l.output(fmt.Sprintf("%s%v", l.formatHeader(level, time.Now()), msg))
	}
	return
}

// Fatal logs msg as a fatal message
func (l *Logger) Fatal(msg interface{}) (n int, err error) {
	return l.Log(1, msg)
}

// Error logs msg as a error message
func (l *Logger) Error(msg interface{}) (n int, err error) {
	return l.Log(2, msg)
}

// Warn logs msg as a warn message
func (l *Logger) Warn(msg interface{}) (n int, err error) {
	return l.Log(3, msg)
}

// Info logs msg as an info message
func (l *Logger) Info(msg interface{}) (n int, err error) {
	return l.Log(4, msg)
}

// Debug logs msg as a debug message
func (l *Logger) Debug(msg interface{}) (n int, err error) {
	return l.Log(5, msg)
}

// Trace logs msg as a trace message
func (l *Logger) Trace(msg interface{}) (n int, err error) {
	return l.Log(6, msg)
}

// Output writes msg to Logger
func (l *Logger) output(msg string) (n int, err error) {

	// Append a new line to output if it does not exist yet
	if msg[len(msg)-1] != '\n' {
		msg = msg + "\n"
	}

	// Write msg to out
	n, err = fmt.Fprint(l.out, msg)
	return
}

func (l *Logger) formatHeader(loglevel int, logTime time.Time) string {

	var level string
	switch loglevel {
	case 1:
		level = "FATAL"
	case 2:
		level = "ERROR"
	case 3:
		level = "WARN"
	case 4:
		level = "INFO"
	case 5:
		level = "DEBUG"
	case 6:
		level = "TRACE"
	}

	return fmt.Sprintf(" %s %s | [%s] ", logTime.Format(LogTimeFormat), l.name, level)
}

func (l *Logger) SetLevel(logLevel string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	logLevel = strings.ToLower(logLevel)
	switch logLevel {
	case LevelOff:
		l.level = 0
	case LevelFatal:
		l.level = 1
	case LevelError:
		l.level = 2
	case LevelWarn:
		l.level = 3
	case LevelInfo:
		l.level = 4
	case LevelDebug:
		l.level = 5
	case LevelTrace:
		l.level = 6
	default:
		l.level = 4
	}
}
