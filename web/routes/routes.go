package routes

import (
	"VDController/web/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// 开发模式：DebugMode 线上模式：ReleaseMode
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	// 加载模板
	router.LoadHTMLGlob("webSrc/template/*")
	// 路由设置
	router.GET("/", controller.Index)
	kube := router.Group("/kube")
	{
		kube.GET("/", controller.KubeJson)
		kube.GET("/deployments/:namespace", controller.GetDeployments)
		kube.GET("/deployment/:namespace/:name", controller.GetDeployment)
		kube.GET("/services/:namespace", controller.GetServices)
		kube.GET("/pods/:namespace", controller.GetPods)
		kube.GET("/pod/:namespace/:name", controller.GetPod)
		kube.GET("/namespaces", controller.GetNameSpaces)
		kube.POST("/uploadYaml", controller.UploadYaml)
		kube.DELETE("/deleteYaml/:file", controller.DeleteYaml)
		kube.GET("/listYaml", controller.ListYamlFiles)
	}
	docker := router.Group("/docker")
	{
		docker.GET("/", controller.DockerJson)
		docker.GET("/search", controller.SearchCtr)
	}
	return router
}
