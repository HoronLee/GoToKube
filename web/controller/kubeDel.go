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
