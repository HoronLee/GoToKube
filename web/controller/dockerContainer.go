package controller

import (
	"GoToKube/docker"
	"GoToKube/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateContainer 创建 Docker 容器
func CreateContainer(c *gin.Context) {
	var request struct {
		ImageName     string            `json:"imageName"`
		ContainerName string            `json:"containerName"`
		Cmd           []string          `json:"cmd,omitempty"`
		PortBindings  map[string]string `json:"portBindings"`
		Volumes       map[string]string `json:"volumes"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	containerID, err := docker.CreateContainer(request.ImageName, request.ContainerName, request.Cmd, request.PortBindings, request.Volumes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"containerID": containerID})
}

func SearchContainer(c *gin.Context) {
	ctrName, ok := c.GetQuery("ctr")
	outPut := make(map[string]interface{})
	if !ok {
		outPut["error"] = "No Such Resource."
	} else {
		containers, err := docker.GetCtrByImg(ctrName)
		if err != nil {
			outPut["error"] = err.Error()
		} else {
			outPut["containers"] = containers // 直接将容器列表放入 map
		}
	}
	c.JSON(http.StatusOK, outPut)
}

// StartContainer 启动指定的 Docker 容器
func StartContainer(c *gin.Context) {
	containerID := c.Param("id")
	if containerID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Container ID not provided"})
		return
	}
	if _, err := docker.StartContainer(containerID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Container started successfully"})
}

// StopContainer 停止指定的 Docker 容器
func StopContainer(c *gin.Context) {
	containerID := c.Param("id")
	if containerID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Container ID not provided"})
		return
	}
	if _, err := docker.StopContainer(containerID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Container stopped successfully"})
}

// DeleteContainer 删除指定的 Docker 容器
func DeleteContainer(c *gin.Context) {
	containerID := c.Param("id")
	if containerID == "" {
		logger.GlobalLogger.Error("Container ID not provided")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Container ID not provided"})
		return
	}
	if err := docker.DeleteContainer(containerID); err != nil {
		logger.GlobalLogger.Error("Failed to delete container")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Container deleted successfully"})
}
