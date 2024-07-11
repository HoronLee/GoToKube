package main

import (
	"GoToKube/config"
	"GoToKube/database"
	"GoToKube/docker"
	"GoToKube/kubernetes"
	"GoToKube/logger"
	"GoToKube/terminal"
	"GoToKube/web"
	"GoToKube/web/models"
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
		db, _ := database.GetDBConnection()
		if err := db.AutoMigrate(&models.StatusInfo{}); err != nil {
			logger.GlobalLogger.Error("Migrate table failed")
			panic(err)
		} else if !config.ConfigData.KubeEnable {
			fmt.Println("⚓️不启用 kubenetes 控制器")
			if !docker.CheckStatus() {
				logger.GlobalLogger.Error("Docker are not health")
				panic("Docker are not health")
			}
		} else {
			fmt.Println("⚓️已启用 kubenetes 控制器")
			if kubernetes.CheckStatus() && docker.CheckStatus() {
				logger.GlobalLogger.Info("All components are running")
			} else {
				logger.GlobalLogger.Error("Kubernetes or Docker are not health")
				panic("Kubernetes or Docker are not health")
			}
		}
	} else {
		logger.GlobalLogger.Error("Database are not health,please check the relevant configuration of the database")
		panic("Database are not health")
	}
	if config.ConfigData.TermEnable {
		mainWg.Add(1)
		go terminal.Terminal(&mainWg)
	}
	mainWg.Add(1)
	web.CheckStatus(&mainWg)
	mainWg.Wait()
}
