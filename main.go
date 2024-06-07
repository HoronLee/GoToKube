package main

import (
	"VDController/database"
	"VDController/docker"
	"VDController/kubernetes"
	"VDController/logger"
	"VDController/terminal"
	"VDController/web"
	"VDController/web/models"
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
		} else if docker.CheckStatus() && kubernetes.CheckStatus() {
			logger.GlobalLogger.Info("All components are running")
		}
	} else {
		logger.GlobalLogger.Error("Some components are not running")
	}
	web.CheckStatus()
}
