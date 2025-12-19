package logger

import (
	"context"
	"encoding/json"
	"fmt"

	gormLogger "gorm.io/gorm/logger"
)

func InfoWithFields(msg string, fields map[string]interface{}) {
	GetLogger().InfoWithFields(context.Background(), msg, fields)
}

func WarnWithFields(msg string, fields map[string]interface{}) {
	GetLogger().WarnWithFields(context.Background(), msg, fields)
}

func ErrorWithFields(msg string, fields map[string]interface{}) {
	GetLogger().ErrorWithFields(context.Background(), msg, fields)
}

func DebugWithFields(msg string, fields map[string]interface{}) {
	GetLogger().DebugWithFields(context.Background(), msg, fields)
}

func (l *Logger) InfoWithFields(ctx context.Context, msg string, fields map[string]interface{}) {
	if l.LogLevel >= gormLogger.Info {
		logEntry := formatLogEntry("INFO", msg, fields, l.config.Colorful, Green)
		l.Logger.Print(logEntry)
	}
}

func (l *Logger) WarnWithFields(ctx context.Context, msg string, fields map[string]interface{}) {
	if l.LogLevel >= gormLogger.Warn {
		logEntry := formatLogEntry("WARN", msg, fields, l.config.Colorful, Yellow)
		l.Logger.Print(logEntry)
	}
}

func (l *Logger) ErrorWithFields(ctx context.Context, msg string, fields map[string]interface{}) {
	if l.LogLevel >= gormLogger.Error {
		logEntry := formatLogEntry("ERROR", msg, fields, l.config.Colorful, Red)
		l.Logger.Print(logEntry)
	}
}

func (l *Logger) DebugWithFields(ctx context.Context, msg string, fields map[string]interface{}) {
	if l.LogLevel >= gormLogger.Info {
		logEntry := formatLogEntry("DEBUG", msg, fields, l.config.Colorful, Cyan)
		l.Logger.Print(logEntry)
	}
}

func formatLogEntry(level, msg string, fields map[string]interface{}, colorful bool, color string) string {
	fieldsJSON, _ := json.Marshal(fields)

	if colorful {
		return fmt.Sprintf("%s[%s]%s %s | %s", color, level, Reset, msg, string(fieldsJSON))
	}
	return fmt.Sprintf("[%s] %s | %s", level, msg, string(fieldsJSON))
}

func WithRequestID(requestID string, fields map[string]interface{}) map[string]interface{} {
	if fields == nil {
		fields = make(map[string]interface{})
	}
	fields["request_id"] = requestID
	return fields
}

func WithUserID(userID interface{}, fields map[string]interface{}) map[string]interface{} {
	if fields == nil {
		fields = make(map[string]interface{})
	}
	fields["user_id"] = userID
	return fields
}

func WithError(err error, fields map[string]interface{}) map[string]interface{} {
	if fields == nil {
		fields = make(map[string]interface{})
	}
	if err != nil {
		fields["error"] = err.Error()
	}
	return fields
}
