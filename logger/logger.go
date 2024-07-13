package logger

import (
	"errors"
	"os"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

type LogLevel int

var (
	GlobalLogger *Logger
	mu           sync.Mutex
)

const (
	INFO LogLevel = iota
	WARNING
	ERROR
)

type Logger struct {
	logger *logrus.Logger
}

// 通过环境变量获取配置
func getConfig(key string, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func InitGlobalLogger(level LogLevel) error {
	mu.Lock()
	defer mu.Unlock()
	if GlobalLogger != nil {
		return errors.New("global logger already initialized")
	}
	GlobalLogger = NewLogger(level)
	return nil
}

func NewLogger(level LogLevel) *Logger {
	logger := logrus.New()

	switch level {
	case INFO:
		logger.SetLevel(logrus.InfoLevel)
	case WARNING:
		logger.SetLevel(logrus.WarnLevel)
	case ERROR:
		logger.SetLevel(logrus.ErrorLevel)
	}

	logDir := getConfig("LOG_DIR", "logs")
	logFileName := logDir + "/" + time.Now().Format("2006-01-02") + ".log"

	// 目录不存在则新建
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		if err := os.Mkdir(logDir, 0755); err != nil {
			os.Exit(0)
		}
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

func (l *Logger) Log(level LogLevel, msg string) {
	if l.logger.IsLevelEnabled(logrus.Level(level)) {
		switch level {
		case INFO:
			l.logger.Info(msg)
		case WARNING:
			l.logger.Warn(msg)
		case ERROR:
			l.logger.Error(msg)
		}
	}
}

// 也可以直接使用logrus提供的方法
func (l *Logger) Info(msg string) {
	l.logger.Info(msg)
}

func (l *Logger) Warn(msg string) {
	l.logger.Warn(msg)
}

func (l *Logger) Error(msg string) {
	l.logger.Error(msg)
}
