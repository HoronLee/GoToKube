package main

import (
	"GoToKube/config"
	"GoToKube/database"
	"GoToKube/docker"
	"GoToKube/kubernetes"
	"GoToKube/logger"
	"GoToKube/terminal"
	"GoToKube/web"
	"fmt"
	"sync"
)

var mainWg sync.WaitGroup

func main() {
	err := logger.InitGlobalLogger(logger.INFO)
	if err != nil {
		return
	}
	checkStatus()
}

func checkStatus() {
	// 检查组件状态
	if database.CheckStatus() {
		_, err := database.GetDBConnection()
		if err != nil {
			logger.GlobalLogger.Error("Database connection failed")
			panic(err)
		} else if !config.ConfigData.KubeEnable {
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
	if config.ConfigData.TermEnable {
		mainWg.Add(1)
		go terminal.Terminal(&mainWg)
	}
	mainWg.Add(1)
	web.CheckStatus(&mainWg)
	mainWg.Wait()
}
