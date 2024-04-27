package main

import (
	"VDController/config"
	"VDController/docker"
	//"VDController/docker"
	"VDController/logger"
	"VDController/terminal"
	"VDController/web"
	"fmt"
	"sync"
)

func main() {
	mLogger := logger.NewLogger(logger.INFO)
	mLogger.Log(logger.INFO, "启动主程序")
	// 检查Docker状态
	docker.CheckState()
	if config.ConfigData.WebEnable {
		fmt.Println("✅启动 Web 服务在：" + config.ConfigData.ListeningAddr)
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
