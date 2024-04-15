package web

import (
	"VDController/docker"
	"VDController/logger"
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

var (
	// Web端日志记录器
	wLogger *logger.Logger
	// 互斥锁，保证线程安全
	mutex sync.Mutex
)

func vdIndex(c *gin.Context) {
	envInfo := docker.GetEnvInfo()
	// 渲染模板
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"DockerVersion":   envInfo.DockerVersion,
		"DcomposeVersion": envInfo.DcomposeVersion,
	})
}

func StartWeb() {
	// 加锁，确保线程安全
	mutex.Lock()
	defer mutex.Unlock()
	// web 日志
	wLogger = logger.NewLogger(logger.INFO)
	wLogger.Log(logger.INFO, "启动Web程序")
	// 此日志仅用于记录 Gin 框架本在终端出现的回显，开发测试用途
	file, err := os.OpenFile("./logs/web.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Failed to open log file:", err)
		return
	}
	defer file.Close()
	// 设置日志输出到文件
	logrus.SetOutput(file)
	// 启用Gin的默认日志中间件，将日志输出到日志文件中
	//gin.DefaultWriter = file
	// Gin 路由设置
	router := gin.Default()
	// 加载静态文件
	router.Static("/web", "./web/static")
	// 加载模板
	router.LoadHTMLGlob("./web/template/*")
	router.GET("/", vdIndex)
	// 创建监听端口
	if err := router.Run(":8080"); err != nil {
		wLogger.Log(logger.ERROR, "创建监听端口失败")
		panic("ListenAndServe: " + err.Error())
	}
}
