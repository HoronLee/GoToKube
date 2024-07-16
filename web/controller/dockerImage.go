package controller

import (
	"GoToKube/docker"
	"GoToKube/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

func GetImages(c *gin.Context) {
	images, err := docker.GetImages()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	logger.GlobalLogger.Info("Fetched images successfully")
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
	logger.GlobalLogger.Info("Image uploaded successfully")
	c.JSON(http.StatusOK, gin.H{"message": "Image uploaded successfully"})
}

func DeleteImage(c *gin.Context) {
	imageID := c.Param("id")
	if imageID == "" {
		logger.GlobalLogger.Error("Image ID not provided")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image ID not provided"})
		return
	}
	if err := docker.DeleteImage(imageID); err != nil {
		logger.GlobalLogger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Image deleted successfully"})
}
