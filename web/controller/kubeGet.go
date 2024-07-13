package controller

import (
	"GoToKube/kubernetes"
	"github.com/gin-gonic/gin"
	"net/http"
)

func KubeJson(c *gin.Context) {
	c.JSON(http.StatusOK, kubernetes.EnvInfo)
}
func GetDeployments(c *gin.Context) {
	namespace := c.Param("namespace")
	deployments, err := kubernetes.GetDeployments(namespace)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, deployments)
}
func GetDeployment(c *gin.Context) {
	name := c.Param("name")
	namespace := c.Param("namespace")
	deployment, err := kubernetes.GetDeployment(name, namespace)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, deployment)
}

func GetServices(c *gin.Context) {
	namespace := c.Param("namespace")
	services, err := kubernetes.GetServices(namespace)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, services)
}

func GetPods(c *gin.Context) {
	namespace := c.Param("namespace")
	services, err := kubernetes.GetPods(namespace)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, services)
}

func GetPod(c *gin.Context) {
	name := c.Param("name")
	namespace := c.Param("namespace")
	deployment, err := kubernetes.GetPod(name, namespace)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, deployment)
}

func GetNameSpaces(c *gin.Context) {
	services, err := kubernetes.GetNamespaces()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, services)
}
