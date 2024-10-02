package logger

import "context"

type Logger interface {
	Info(format string, v ...interface{})
	Debug(format string, v ...interface{})
	Warn(format string, v ...interface{})
	Error(format string, v ...interface{})
	Print(format string, v ...interface{})
	With(key string, value interface{}) Logger
	WithContext(ctx context.Context) context.Context
}
