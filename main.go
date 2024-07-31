package main

import (
	"GoToKube/config"
	"GoToKube/database"
	"GoToKube/docker"
	"GoToKube/kubernetes"
	"GoToKube/logger"
	"GoToKube/terminal"
	"GoToKube/web"
	"GoToKube/web/auth"
	"GoToKube/web/models"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"sync"
)

var mainWg sync.WaitGroup

func main() {
	err := logger.InitGlobalLogger(logrus.InfoLevel)
	if err != nil {
		log.Fatalf("Failed to initialize global logger: %v", err)
	}
	config.InitConfig()

	// 检查组件状态
	checkStatus()

	db, err := database.GetDBConnection()
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	// 自动迁移模型
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return
	}

	// 初始化根用户
	if err := auth.InitRootUser(); err != nil {
		log.Fatalf("failed to initialize root user: %v", err)
	}

	// 启动服务
	mainWg.Add(1)
	web.CheckStatus(&mainWg)
	if config.Data.Common.TermEnable {
		mainWg.Add(1)
		go terminal.Terminal(&mainWg)
	}
	mainWg.Wait()
}

func checkStatus() {
	// 检查组件状态
	if database.CheckStatus() {
		db, err := database.GetDBConnection()
		err = db.AutoMigrate(&models.User{})
		if err := auth.InitRootUser(); err != nil {
			log.Fatalf("failed to initialize root user: %v", err)
		}
		if err != nil {
			logger.GlobalLogger.Error("Database connection failed")
			panic(err)
		} else if !config.Data.Kubernetes.Enable {
			fmt.Println("⚓️不启用 kubernetes 控制器")
			if !docker.CheckStatus() {
				panic("Docker is not healthy,please start docker")
			}
		} else {
			fmt.Println("⚓️已启用 kubernetes 控制器")
			if kubernetes.CheckStatus() && docker.CheckStatus() {
				logger.GlobalLogger.Info("All components are running")
			} else {
				panic("Kubernetes or Docker is not healthy")
			}
		}
	} else {
		logger.GlobalLogger.Error("Database is not healthy, please check the relevant configuration of the database")
		panic("Database is not healthy")
	}
}
