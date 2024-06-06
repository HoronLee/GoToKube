package main

import (
	"VDController/database"
	"VDController/docker"
	"VDController/kubernetes"
	"VDController/logger"
	"VDController/terminal"
	"VDController/web"
	"sync"
)

func main() {
	logger.InitGlobalLogger(logger.INFO)
	// 检查组件状态
	database.CheckStatus()
	docker.CheckStatus()
	kubernetes.CheckStatus()
	web.CheckStatus()
	// 控制台协程
	var mainWg sync.WaitGroup
	mainWg.Add(1)
	go terminal.Terminal(&mainWg)
	mainWg.Wait()
}
