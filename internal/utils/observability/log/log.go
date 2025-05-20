package log

import (
	"log/slog"
	"os"
)

type Level int8

const (
	LevelDebug Level = -4
	LevelInfo  Level = 0
	LevelWarn  Level = 4
	LevelError Level = 8
)

func convertToSlogLevel(level Level) slog.Level {
	switch level {
	case LevelDebug:
		return slog.LevelDebug
	case LevelInfo:
		return slog.LevelInfo
	case LevelWarn:
		return slog.LevelWarn
	case LevelError:
		return slog.LevelError
	default:
		return slog.LevelInfo

	}
}

type Logger struct {
	*slog.Logger
}

func NewLogger(level Level) *Logger {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: convertToSlogLevel(level),
	})

	return &Logger{Logger: slog.New(handler)}
}

func (l *Logger) Error(msg string, args ...any) {
	args = append([]any{"error"}, args...)
	l.Logger.Error(msg, args...)
}

func (l *Logger) Info(msg string, args ...any) {
	args = append([]any{"info"}, args...)
	l.Logger.Info(msg, args...)
}

func (l *Logger) Debug(msg string, args ...any) {
	args = append([]any{"debug"}, args...)
	l.Logger.Debug(msg, args...)
}

func (l *Logger) Warn(msg string, args ...any) {
	args = append([]any{"warn"}, args...)
	l.Logger.Warn(msg, args...)
}
