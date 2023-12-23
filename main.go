package main

import (
	"VDController/docker"
	"VDController/terminal"
	"sync"
)

func main() {
	mainLogger := terminal.NewLogger(terminal.INFO)
	mainLogger.Log(terminal.INFO, "启动主程序")
	// 检查Docker状态
	docker.CheckState()
	// 控制台协程
	var terminalWg sync.WaitGroup
	terminalWg.Add(1)
	go terminal.Terminal(&terminalWg)
	terminalWg.Wait()
}
