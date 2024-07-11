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

func SearchCtr(c *gin.Context) {
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
		logger.GlobalLogger.Log(logger.ERROR, "Failed to fetch images")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch images"})
		return
	}
	logger.GlobalLogger.Log(logger.INFO, "Fetched images successfully")
	c.JSON(http.StatusOK, images)
}

func UploadImage(c *gin.Context) {
	// 从请求中读取文件
	file, err := c.FormFile("file")
	if err != nil {
		logger.GlobalLogger.Log(logger.ERROR, "Failed to get form file")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get form file"})
		return
	}
	// 创建文件保存路径
	savePath := filepath.Join("/tmp", file.Filename)
	// 保存文件到指定路径
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		logger.GlobalLogger.Log(logger.ERROR, "Failed to save uploaded file")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save uploaded file"})
		return
	}
	// 上传镜像
	if err := docker.UploadImage(savePath); err != nil {
		logger.GlobalLogger.Log(logger.ERROR, "Failed to upload image")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image"})
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
		logger.GlobalLogger.Log(logger.ERROR, "Failed to delete image")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete image"})
		return
	}
	logger.GlobalLogger.Log(logger.INFO, "Image deleted successfully")
	c.JSON(http.StatusOK, gin.H{"message": "Image deleted successfully"})
}
