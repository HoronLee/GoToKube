package controller

import (
	"VDController/docker"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DockerJson(c *gin.Context) {
	c.JSON(http.StatusOK, docker.EnvInfo)
}
