package web

import (
	"GoToKube/config"
	"GoToKube/logger"
	"GoToKube/web/routes"
	"fmt"

	"sync"
)

var (
	// 互斥锁，保证线程安全
	mutex sync.Mutex
)

func CheckStatus(wg *sync.WaitGroup) {
	if config.Data.Web.Enable {
		fmt.Println("✅ 在 http://" + config.Data.Web.ListeningAddr + " 上启动 Web 服务")
		wg.Add(1)
		go StartWeb()
	} else {
		fmt.Println("❎ 不启动 Web 服务")
	}
}

func StartWeb() {
	mutex.Lock()
	defer mutex.Unlock()

	logger.GlobalLogger.Info("Launching the Web Application")
	listeningAddr := config.Data.Web.ListeningAddr

	router := routes.SetupRouter()
	if router == nil {
		logger.GlobalLogger.Error("Failed to setup router")
		return
	}

	if err := router.Run(listeningAddr); err != nil {
		logger.GlobalLogger.Error("Failed to create listening port")
		panic("ListenAndServe: " + err.Error())
	} else {
		msg := "Listening and serving HTTP on " + listeningAddr
		logger.GlobalLogger.Error(msg)
	}
}
