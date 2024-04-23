package web

import (
	"VDController/config"
	"VDController/docker"
	"VDController/logger"

	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	// Web端日志记录器
	wLogger *logger.Logger
	// 互斥锁，保证线程安全
	mutex sync.Mutex
)

func StartWeb() {
	// 加锁，确保线程安全
	mutex.Lock()
	defer mutex.Unlock()
	wLogger = logger.NewLogger(logger.INFO)
	wLogger.Log(logger.INFO, "启动Web程序")
	listeningAddr := config.ConfigData.ListeningAddr
	// ======
	// 此日志仅用于记录 Gin 框架本在终端出现的回显，开发测试用途
	file, err := os.OpenFile("./logs/web.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Failed to open log file:", err)
		return
	}
	defer file.Close()
	logrus.SetOutput(file)
	gin.DefaultWriter = file
	// ======
	// Gin 路由设置
	router := gin.Default()
	// 加载静态文件
	router.Static("/web", "./web/static")
	// 加载模板
	router.LoadHTMLGlob("./web/template/*")
	router.GET("/", vdIndex)
	router.GET("/json", jsonIndex)
	router.GET("/search", search)
	// 创建监听端口
	if err := router.Run(listeningAddr); err != nil {
		wLogger.Log(logger.ERROR, "创建监听端口失败")
		panic("ListenAndServe: " + err.Error())
	} else {
		msg := "Web端在" + listeningAddr + "上开启"
		wLogger.Log(logger.INFO, msg)
	}
}

func vdIndex(c *gin.Context) {
	envInfo := docker.GetEnvInfo()
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"DockerV":        envInfo.DockerVersion,
		"DockerComposeV": envInfo.DockerCVersion,
	})
}

func jsonIndex(c *gin.Context) {
	envInfo := docker.GetEnvInfo()
	c.JSON(http.StatusOK, envInfo)
}

func search(c *gin.Context) {
	imgName, ok := c.GetQuery("image")
	outPut := make(map[string]interface{})
	if !ok {
		outPut["error"] = "No Such Resource."
	} else {
		outPut, _ = docker.Dockerclient.DockerLsByImg(imgName)
	}
	c.JSON(http.StatusOK, outPut)
}
