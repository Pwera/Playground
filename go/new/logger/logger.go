package logger

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
)

type Level int

const (
	LevelDebug Level = iota
	LevelInfo
	LevelError
)
var loggerData = map[Level]string{
	LevelDebug: "Debug",
	LevelInfo: "Info",
	LevelError: "Error",
}
const defaultLogLevel = LevelInfo

type Logger struct {
	mu     sync.Mutex
	prefix string
	Level  Level
	w      io.Writer
	buf    bytes.Buffer
}

func New(w io.Writer, prefix string) *Logger {
	return &Logger{
		prefix: prefix,
		Level:  defaultLogLevel,
		w:      w,
	}
}

var Console = New(os.Stdout, "")

func (l *Logger) Debug(v ...interface{}) {
	l.WriteEntry(LevelDebug, fmt.Sprintln(v...))
}
func (l *Logger) Info(v ...interface{}) {
	if LevelInfo < l.Level {
		return
	}
	l.WriteEntry(LevelInfo, fmt.Sprintln(v...))
}
func (l *Logger) Error(v ...interface{}) {
	if LevelError < l.Level {
		return
	}
	l.WriteEntry(LevelError, fmt.Sprintln(v...))
}
func (l *Logger) WriteEntry(lvl Level, msg string) error{
	_, err := l.w.Write([]byte(loggerData[lvl] +" "+ msg))
	return err
}
func (l *Logger) SetCurrentLevel(lvl Level){
	l.Level=lvl
}
