package golog

import (
	"os"
	"testing"
)

func TestLog(t *testing.T) {
	l := NewLogger(os.Stdout, "Test", "trace")
	l.Fatal("This is a fatal message")
	l.Error("This is an error message")
	l.Warn("This is a warn message")
	l.Info("This is an info message")
	l.Debug("This is a debug message")
	l.Trace("This is a trace message")
}

func TestSetLevel(t *testing.T) {
	l := NewLogger(os.Stdout, "Test", "off")
	l.SetLevel("Fatal")
	l.SetLevel("Error")
	l.SetLevel("Warn")
	l.SetLevel("Info")
	l.SetLevel("Debug")
	l.SetLevel("Trace")
	l.SetLevel("Default")
}

func TestNonLevelMsg(t *testing.T) {
	l := NewLogger(os.Stdout, "Test", "error")
	l.Warn("Warning, this doesn't show up")
}
