package main

import (
	"VDController/config"
	"VDController/database"
	"VDController/docker"
	"VDController/kubernetes"
	"VDController/logger"
	"VDController/terminal"
	"VDController/web"
	"VDController/web/models"
	"fmt"
	"sync"
)

func main() {
	logger.InitGlobalLogger(logger.INFO)
	checkStatus()
	// 控制台协程
	var mainWg sync.WaitGroup
	mainWg.Add(1)
	go terminal.Terminal(&mainWg)
	mainWg.Wait()
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
			}
		} else {
			fmt.Println("⚓️已启用 kubenetes 控制器")
			if docker.CheckStatus() || kubernetes.CheckStatus() {
				logger.GlobalLogger.Info("All components are running")
			}
		}
	} else {
		logger.GlobalLogger.Error("Database components are not health")
	}
	web.CheckStatus()
}
