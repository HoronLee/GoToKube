package logger

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

type LogLevel int

const (
	INFO LogLevel = iota
	WARNING
	ERROR
)

// 日志记录器
type Logger struct {
	logger *logrus.Logger
}

// 创建日志记录器
func NewLogger(level LogLevel) *Logger {
	// 初始化logrus logger
	logger := logrus.New()

	// 设置日志级别
	switch level {
	case INFO:
		logger.SetLevel(logrus.InfoLevel)
	case WARNING:
		logger.SetLevel(logrus.WarnLevel)
	case ERROR:
		logger.SetLevel(logrus.ErrorLevel)
	}

	// 检查并创建 logs 文件夹
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		err := os.Mkdir("logs", 0755)
		if err != nil {
			logger.Fatalf("Error creating logs directory: %v", err)
		}
	}

	// 设置日志输出
	logFileName := "./logs/vdc_" + time.Now().Format("2006-01-02") + ".log"
	file, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logger.Fatalf("Error opening log file: %v", err)
	}
	logger.SetOutput(file)
	// 设置时间格式
	logger.SetFormatter(&logrus.TextFormatter{
		DisableTimestamp: false,
		TimestampFormat:  "2006-01-02",
	})
	return &Logger{logger: logger}
}

// 日志记录
func (l *Logger) Log(level LogLevel, msg string) {
	switch level {
	case INFO:
		l.logger.Info(msg)
	case WARNING:
		l.logger.Warn(msg)
	case ERROR:
		l.logger.Error(msg)
	}
}

