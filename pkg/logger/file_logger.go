package logger

import (
	"context"
	"log"
	"os"
	"time"

	"gorm.io/gorm/logger"
)

type FileLogger struct {
	logFile       *log.Logger
	SlowThreshold time.Duration
	LogLevel      logger.LogLevel
}

func NewFileLogger(path string, slow time.Duration, level logger.LogLevel) *FileLogger {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	return &FileLogger{
		logFile:       log.New(f, "", log.LstdFlags|log.Lmicroseconds),
		SlowThreshold: slow,
		LogLevel:      level,
	}
}

func (l *FileLogger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *l
	newLogger.LogLevel = level
	return &newLogger
}

func (l *FileLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Info {
		l.logFile.Printf("[INFO] "+msg, data...)
	}
}

func (l *FileLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Warn {
		l.logFile.Printf("[WARN] "+msg, data...)
	}
}

func (l *FileLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Error {
		l.logFile.Printf("[ERROR] "+msg, data...)
	}
}

func (l *FileLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()
	if err != nil && l.LogLevel >= logger.Error {
		l.logFile.Printf("[ERROR] %s [rows:%d] %v\n", sql, rows, err)
	} else if elapsed > l.SlowThreshold && l.LogLevel >= logger.Warn {
		l.logFile.Printf("[SLOW SQL] %v [rows:%d] %s\n", elapsed, rows, sql)
	} else if l.LogLevel >= logger.Info {
		l.logFile.Printf("[SQL] %v [rows:%d] %s\n", elapsed, rows, sql)
	}
}
