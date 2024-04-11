package logger

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

type LogLevel int

const (
	INFO LogLevel = iota
	WARNING
	ERROR
)

// 日志记录器
type Logger struct {
	level LogLevel
}

// 创建日志记录器
func NewLogger(level LogLevel) *Logger {
	// 检查并创建 logs 文件夹
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		err := os.Mkdir("logs", 0755)
		if err != nil {
			log.Fatalf("Error creating logs directory: %v", err)
		}
	}
	return &Logger{level: level}
}

// 日志记录
func (l *Logger) Log(level LogLevel, msg string) {
	if level >= l.level {
		filename := fmt.Sprintf("./logs/vdc_%s.log", time.Now().Format("2006-01-02"))
		file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		log.SetOutput(file)
		switch level {
		case INFO:
			log.Println("[INFO] " + msg)
		case WARNING:
			log.Println("[WARNING] " + msg)
		case ERROR:
			log.Println("[ERROR] " + msg)
		}
	}
}

func (l *Logger) countLogs() {
	files, err := ioutil.ReadDir("./logs")
	if err != nil {
		log.Fatal(err)
	}

	count := 0
	for _, file := range files {
		if strings.HasPrefix(file.Name(), "vdc_") {
			count++
		}
	}

	if count > 5 {
		l.cleanupLogs()
	}
}

// 清理三天前的日志
func (l *Logger) cleanupLogs() {
	files, err := ioutil.ReadDir("./logs")
	if err != nil {
		log.Fatal(err)
	}

	threeDaysAgo := time.Now().AddDate(0, 0, -3).Format("2006-01-02")
	for _, file := range files {
		if strings.Compare(file.Name()[4:14], threeDaysAgo) < 0 && strings.HasPrefix(file.Name(), "vdc_") {
			os.Remove("./logs/" + file.Name())
		}
	}
}