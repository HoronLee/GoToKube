package controller

import (
	"VDController/docker"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DockerJson(c *gin.Context) {
	c.JSON(http.StatusOK, docker.EnvInfo)
}

func SearchCtr(c *gin.Context) {
	ctrName, ok := c.GetQuery("ctr")
	outPut := make(map[string]interface{})
	if !ok {
		outPut["error"] = "No Such Resource."
	} else {
		containers, err := docker.ContainerLsByImg(ctrName)
		if err != nil {
			outPut["error"] = err.Error()
		} else {
			outPut["containers"] = containers // 直接将容器列表放入 map
		}
	}
	c.JSON(http.StatusOK, outPut)
}
