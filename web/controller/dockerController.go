package controller

import (
	"GoToKube/docker"
	"GoToKube/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

func DockerJson(c *gin.Context) {
	c.JSON(http.StatusOK, docker.EnvInfo)
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

func GetImages(c *gin.Context) {
	images, err := docker.GetImages()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	logger.GlobalLogger.Log(logger.INFO, "Fetched images successfully")
	c.JSON(http.StatusOK, images)
}

func UploadImage(c *gin.Context) {
	// 从请求中读取文件
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 创建文件保存路径
	savePath := filepath.Join("/tmp", file.Filename)
	// 保存文件到指定路径
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 上传镜像
	if err := docker.UploadImage(savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	logger.GlobalLogger.Log(logger.INFO, "Image uploaded successfully")
	c.JSON(http.StatusOK, gin.H{"message": "Image uploaded successfully"})
}

func DeleteImage(c *gin.Context) {
	imageID := c.Param("id")
	if imageID == "" {
		logger.GlobalLogger.Log(logger.ERROR, "Image ID not provided")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image ID not provided"})
		return
	}
	if err := docker.DeleteImage(imageID); err != nil {
		logger.GlobalLogger.Log(logger.ERROR, err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Image deleted successfully"})
}

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

// DeleteContainer 删除指定的 Docker 容器
func DeleteContainer(c *gin.Context) {
	containerID := c.Param("id")
	if containerID == "" {
		logger.GlobalLogger.Log(logger.ERROR, "Container ID not provided")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Container ID not provided"})
		return
	}
	if err := docker.DeleteContainer(containerID); err != nil {
		logger.GlobalLogger.Log(logger.ERROR, "Failed to delete container")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Container deleted successfully"})
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
