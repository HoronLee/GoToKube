package controller

import (
	"VDController/kubernetes"
	"github.com/gin-gonic/gin"
	"net/http"
)

func KubeJson(c *gin.Context) {
	c.JSON(http.StatusOK, kubernetes.EnvInfo)
}
