package main

import (
	"VDController/docker"
	"VDController/logger"
	"VDController/terminal"
	"sync"
)

func main() {
	mLogger := logger.NewLogger(logger.INFO)
	mLogger.Log(logger.INFO, "启动主程序")
	// 检查Docker状态
	docker.CheckState()
	// 控制台协程
	var terminalWg sync.WaitGroup
	terminalWg.Add(1)
	go terminal.Terminal(&terminalWg)
	terminalWg.Wait()
}
