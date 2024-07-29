package routes

import (
	"GoToKube/web/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// 开发模式：DebugMode 线上模式：ReleaseMode
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	// 路由设置
	router.GET("/")
	registerKubeRoutes(router)
	registerDockerRoutes(router)
	return router
}

// registerKubeRoutes groups and registers Kube related routes.
func registerKubeRoutes(router *gin.Engine) {
	kube := router.Group("/kube")
	kube.GET("/deployment/:namespace", controller.GetDeployments)
	kube.GET("/deployment/:namespace/:name", controller.GetDeployment)
	kube.DELETE("/deployment/:namespace/:name", controller.DeleteDeployment)
	kube.GET("/service/:namespace", controller.GetServices)
	kube.DELETE("/service/:namespace/:name", controller.DeleteService)
	kube.GET("/pods/:namespace", controller.GetPods)
	kube.GET("/pod/:namespace/:name", controller.GetPod)
	kube.DELETE("/pod/:namespace/:name", controller.DeletePod)
	kube.GET("/namespace", controller.GetNameSpaces)
	kube.POST("/namespace/create/:name", controller.CreateNamespace)
	kube.DELETE("/namespace/delete/:name", controller.DeleteNamespace)
	kube.POST("/uploadYaml", controller.UploadYaml)
	kube.DELETE("/deleteYaml/:file", controller.DeleteYaml)
	kube.GET("/listYaml", controller.ListYamlFiles)
}

// registerDockerRoutes groups and registers Docker related routes.
func registerDockerRoutes(router *gin.Engine) {
	docker := router.Group("/docker")
	docker.GET("/image", controller.GetImages)
	docker.POST("/uploadImage", controller.UploadImage)
	docker.DELETE("/image/:id", controller.DeleteImage)
	docker.GET("/search", controller.SearchContainer)
	docker.POST("/ctr/create", controller.CreateContainer)
	docker.DELETE("/ctr/delete/:id", controller.DeleteContainer)
	docker.POST("/ctr/stop/:id", controller.StopContainer)
	docker.POST("/ctr/start/:id", controller.StartContainer)
}
