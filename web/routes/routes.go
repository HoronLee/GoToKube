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
