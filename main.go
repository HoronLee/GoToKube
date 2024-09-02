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
	"fmt"
	"log"
	"sync"

	"github.com/sirupsen/logrus"
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
	// 检查数据库状态
	if err := database.CheckStatus(); err != nil {
		logger.GlobalLogger.Error(err.Error())
		panic(err.Error())
	}
	// 获取数据库连接
	if _, err := database.GetDBConnection(); err != nil {
		logger.GlobalLogger.Error("Database connection failed")
		panic(err)
	}
	switch config.Data.Kubernetes.Enable {
	case false:
		fmt.Println("⚓️不启用 kubernetes 控制器")
		if err := docker.CheckStatus(); err != nil {
			panic("Docker is not healthy: " + err.Error())
		}
	case true:
		fmt.Println("⚓️已启用 kubernetes 控制器")
		if err := docker.CheckStatus(); err != nil {
			panic("Docker is not healthy: " + err.Error())
		}
		if err := kubernetes.CheckStatus(); err != nil {
			panic("Kubernetes is not healthy: " + err.Error())
		}
	}
	logger.GlobalLogger.Info("All components are running")
}
