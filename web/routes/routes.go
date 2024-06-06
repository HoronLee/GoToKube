package routes

import (
	"VDController/web/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// 开发模式
	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	// 加载模板
	router.LoadHTMLGlob("webSrc/template/*")

	// 路由设置
	router.GET("/", controller.Index)
	router.GET("/search", controller.Search)

	jsonIndex := router.Group("/json")
	{
		jsonIndex.GET("/docker", controller.DockerJson)
		jsonIndex.GET("/kube", controller.KubeJson)
	}

	return router
}
