package main

import (
	"VDController/docker"
	"VDController/logger"
	"VDController/terminal"
	"VDController/web"
	"sync"
)

func main() {
	mLogger := logger.NewLogger(logger.INFO)
	mLogger.Log(logger.INFO, "启动主程序")
	// 检查Docker状态
	docker.CheckState()
	// 控制台协程
	var mianWg sync.WaitGroup
	mianWg.Add(1)
	go terminal.Terminal(&mianWg)
	// Web 端
	web.StartWeb()
	mianWg.Wait()
}
