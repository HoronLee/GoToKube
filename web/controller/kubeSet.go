package controller

import (
	"GoToKube/kubernetes"
	"github.com/gin-gonic/gin"
)

func CreateNamespace(c *gin.Context) {
	name := c.Param("name")
	if ns, err := kubernetes.CreateNamespace(name); err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"message": ns,
		})
	}
}
