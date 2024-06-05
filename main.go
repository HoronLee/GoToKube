package main

import (
	"VDController/config"
	"VDController/docker"
	"VDController/kubernetes"
	"VDController/logger"
	"VDController/terminal"
	"VDController/web"
	"fmt"
	"sync"
)

func main() {
	logger.InitGlobalLogger(logger.INFO)
	// 检查Docker状态
	docker.Checkstatus()
	if config.ConfigData.KubeEnable {
		fmt.Println("⚓️已启用 kubenetes 控制器")
		kubernetes.Checkstatus()
	} else {
		fmt.Println("⚓️不启用 kubenetes 控制器")
	}
	if config.ConfigData.WebEnable {
		fmt.Println("✅在 http://" + config.ConfigData.ListeningAddr + " 上启动 Web 服务")
		go web.StartWeb()
	} else {
		fmt.Println("❎不启动 Web 服务")
	}
	// 控制台协程
	var mainWg sync.WaitGroup
	mainWg.Add(1)
	go terminal.Terminal(&mainWg)
	mainWg.Wait()
}
