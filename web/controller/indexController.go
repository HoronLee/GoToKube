package controller

import (
	"VDController/docker"
	"VDController/kubernetes"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"DockerV":        docker.EnvInfo.DockerVersion,
		"DockerComposeV": docker.EnvInfo.DockerCVersion,
		"KubeVersion":    kubernetes.EnvInfo.KubeVersion,
	})
}

func Search(c *gin.Context) {
	imgName, ok := c.GetQuery("image")
	outPut := make(map[string]interface{})
	if !ok {
		outPut["error"] = "No Such Resource."
	} else {
		outPut, _ = docker.DockerlsByImg(imgName)
	}
	c.JSON(http.StatusOK, outPut)
}
