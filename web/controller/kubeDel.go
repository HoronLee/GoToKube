package controller

import (
	"GoToKube/kubernetes"
	"github.com/gin-gonic/gin"
)

func DeleteNamespace(c *gin.Context) {
	name := c.Param("name")
	if err := kubernetes.DeleteNamespace(name); err != nil {
		c.JSON(500, gin.H{
			"msg": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"msg": "namespace deleted successfully",
		})
	}
}

func DeleteDeployment(c *gin.Context) {
	name := c.Param("name")
	namespace := c.Param("namespace")
	if err := kubernetes.DeleteDeployment(namespace, name); err != nil {
		c.JSON(500, gin.H{
			"msg": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"msg": "deployment deleted successfully",
		})
	}
}

func DeleteService(c *gin.Context) {
	name := c.Param("name")
	namespace := c.Param("namespace")
	if err := kubernetes.DeleteService(namespace, name); err != nil {
		c.JSON(500, gin.H{
			"msg": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"msg": "service deleted successfully",
		})
	}
}

func DeletePod(c *gin.Context) {
	name := c.Param("name")
	namespace := c.Param("namespace")
	if err := kubernetes.DeletePod(namespace, name); err != nil {
		c.JSON(500, gin.H{
			"msg": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"msg": "pod deleted successfully",
		})
	}
}
