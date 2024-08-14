package routes

import (
	"GoToKube/web/controller"
	"github.com/gin-gonic/gin"
)

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
