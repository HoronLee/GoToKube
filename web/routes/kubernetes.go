package routes

import (
	"GoToKube/web/controller"
	"github.com/gin-gonic/gin"
)

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
