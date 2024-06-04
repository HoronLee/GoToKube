package web

import (
	"VDController/config"
	"VDController/docker"
	"VDController/kubernetes"
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
	wLogger.Log(logger.INFO, "Launching the Web Application")
	listeningAddr := config.ConfigData.ListeningAddr
	// ============
	// 此日志仅用于记录 Gin 框架本在终端出现的回显，开发测试用途
	file, err := os.OpenFile("./logs/web.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Failed to open log file:", err)
		return
	}
	defer file.Close()
	logrus.SetOutput(file)
	gin.DefaultWriter = file
	// ============
	// Gin 路由设置
	router := gin.Default()
	// 加载静态文件
	//router.Static("/web", "./webSrc/static")
	// 加载模板
	router.LoadHTMLGlob("./webSrc/template/*")
	// 路由设置
	router.GET("/", index)
	router.GET("/search", search)
	jsonIndex := router.Group("/json")
	{
		jsonIndex.GET("/docker", dockerJson)	
		jsonIndex.GET("/kube", kubeJson)
	}
	// 创建监听端口
	if err := router.Run(listeningAddr); err != nil {
		wLogger.Log(logger.ERROR, "Failed to create listening port")
		panic("ListenAndServe: " + err.Error())
	} else {
		msg := "Listening and serving HTTP on" + listeningAddr
		wLogger.Log(logger.INFO, msg)
	}
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"DockerV":        docker.EnvInfo.DockerVersion,
		"DockerComposeV": docker.EnvInfo.DockerCVersion,
		"KubeVersion":    kubernetes.EnvInfo.KubeVersion,
	})
}

func dockerJson(c *gin.Context) {
	c.JSON(http.StatusOK, docker.EnvInfo)
}
func kubeJson(c *gin.Context) {
	c.JSON(http.StatusOK, kubernetes.EnvInfo)
}

func search(c *gin.Context) {
	imgName, ok := c.GetQuery("image")
	outPut := make(map[string]interface{})
	if !ok {
		outPut["error"] = "No Such Resource."
	} else {
		outPut, _ = docker.Dockerclient.DockerlsByImg(imgName)
	}
	c.JSON(http.StatusOK, outPut)
}
