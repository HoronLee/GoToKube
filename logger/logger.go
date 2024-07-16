package logger

import (
	"errors"
	"github.com/sirupsen/logrus"
	"os"
	"sync"
)

var (
	GlobalLogger *Logger
	mu           sync.Mutex
)

type Logger struct {
	logger *logrus.Logger
}

func getConfigFromEnv(key string, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func InitGlobalLogger(level logrus.Level) error {
	mu.Lock()
	defer mu.Unlock()
	if GlobalLogger != nil {
		return errors.New("global logger already initialized")
	}
	GlobalLogger = NewLogger(level)
	return nil
}

func NewLogger(level logrus.Level) *Logger {
	logger := logrus.New()
	logger.SetLevel(level)

	logDir := getConfigFromEnv("LOG_DIR", ".")
	logFileName := logDir + "/app.log"

	err := os.MkdirAll(logDir, os.ModePerm)
	if err != nil {
		logger.Fatalf("failed to create log directory: %v", err)
	}

	// 使用logrus的RotatingFileWriter实现日志滚动
	file, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		os.Exit(0)
	}
	logger.SetOutput(file)

	// 设置时间格式
	logger.SetFormatter(&logrus.TextFormatter{
		DisableTimestamp: false,
		TimestampFormat:  "2006-01-02 15:04:05",
	})

	return &Logger{logger: logger}
}

func (l *Logger) Log(level logrus.Level, msg string) {
	if l.logger.IsLevelEnabled(level) {
		switch level {
		case logrus.InfoLevel:
			l.logger.Info(msg)
		case logrus.WarnLevel:
			l.logger.Warn(msg)
		case logrus.ErrorLevel:
			l.logger.Error(msg)
		default:
			panic("unhandled default case")
		}
	}
}

func (l *Logger) Info(msg string) {
	l.logger.Info(msg)
}

func (l *Logger) Warn(msg string) {
	l.logger.Warn(msg)
}

func (l *Logger) Error(msg string) {
	l.logger.Error(msg)
}

func (l *Logger) Fatal(msg string) {
	l.logger.Fatal(msg)
}
