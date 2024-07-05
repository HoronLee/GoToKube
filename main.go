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
	web.CheckStatus()
}
