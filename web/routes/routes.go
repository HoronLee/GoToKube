package routes

import (
	"GoToKube/web/auth"
	"GoToKube/web/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// 设置为测试模式或线上模式
	gin.SetMode(gin.TestMode) // 根据环境改为 gin.ReleaseMode 或 gin.DebugMode
	router := gin.Default()

	// 身份认证相关路由
	router.POST("/register", controller.Register)
	router.POST("/login", controller.Login)

	// 使用 JWT 中间件保护需要身份认证的路由
	kube := router.Group("/kube")
	kube.Use(auth.JWTMiddleware())
	registerKubeRoutes(kube)

	docker := router.Group("/docker")
	docker.Use(auth.JWTMiddleware())
	registerDockerRoutes(docker)

	return router
}

// registerKubeRoutes groups and registers Kubernetes related routes.
func registerKubeRoutes(kube *gin.RouterGroup) {
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
func registerDockerRoutes(docker *gin.RouterGroup) {
	docker.GET("/image", controller.GetImages)
	docker.POST("/uploadImage", controller.UploadImage)
	docker.DELETE("/image/:id", controller.DeleteImage)
	docker.GET("/search", controller.SearchContainer)
	docker.POST("/ctr/create", controller.CreateContainer)
	docker.DELETE("/ctr/delete/:id", controller.DeleteContainer)
	docker.POST("/ctr/stop/:id", controller.StopContainer)
	docker.POST("/ctr/start/:id", controller.StartContainer)
}
