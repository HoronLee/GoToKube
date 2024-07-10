package controller

import (
	"VDController/kubernetes"
	"VDController/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)

func KubeJson(c *gin.Context) {
	c.JSON(http.StatusOK, kubernetes.EnvInfo)
}
func GetDeployments(c *gin.Context) {
	namespace := c.Param("namespace")
	deployments, err := kubernetes.GetDeployments(namespace)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, deployments)
}
func GetDeployment(c *gin.Context) {
	name := c.Param("name")
	namespace := c.Param("namespace")
	deployment, err := kubernetes.GetDeployment(name, namespace)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, deployment)
}

func GetServices(c *gin.Context) {
	namespace := c.Param("namespace")
	services, err := kubernetes.GetServices(namespace)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, services)
}

func GetPods(c *gin.Context) {
	namespace := c.Param("namespace")
	services, err := kubernetes.GetPods(namespace)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, services)
}

func GetPod(c *gin.Context) {
	name := c.Param("name")
	namespace := c.Param("namespace")
	deployment, err := kubernetes.GetPod(name, namespace)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, deployment)
}

func GetNameSpaces(c *gin.Context) {
	services, err := kubernetes.GetNamespaces()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, services)
}

// UploadYaml 单文件创建资源
func UploadYaml(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	if !kubernetes.IsYAML(file.Filename) {
		c.String(http.StatusBadRequest, "only .yml or .yaml files are allowed")
		return
	}
	logger.GlobalLogger.Info("upload file: " + file.Filename)
	uploadDir := "./uploads"
	dst := uploadDir + "/" + file.Filename
	// 目录不存在则新建
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		if err := os.Mkdir(uploadDir, 0755); err != nil {
			os.Exit(0)
		}
	}
	// 检查是否存在同名文件
	if _, err := os.Stat(dst); err == nil {
		c.String(http.StatusConflict, fmt.Sprintf("file '%s' already exists", file.Filename))
		return
	} else if !os.IsNotExist(err) {
		c.String(http.StatusInternalServerError, fmt.Sprintf("failed to check file existence: %s", err.Error()))
		return
	}
	// 保存文件
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}
	err = kubernetes.ApplyYAML(dst)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("failed to apply YAML file: %s", err.Error()))
		err := os.Remove(dst)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		logger.GlobalLogger.Info("delete file: " + dst)
		return
	}
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded and resources created!", file.Filename))
}

func ListYamlFiles(c *gin.Context) {
	files, err := os.ReadDir("./uploads")
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("read dir err: %s", err.Error()))
		return
	}
	var fileNames []string
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}
	c.JSON(http.StatusOK, fileNames)
}

func DeleteYaml(c *gin.Context) {
	fileName := c.Param("file")
	filePath := filepath.Join("./uploads", fileName)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "file not found"})
		return
	}

	err := kubernetes.DeleteYAML(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		err := os.Remove(filePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		logger.GlobalLogger.Info("delete file: " + fileName)
	}

	c.JSON(http.StatusOK, gin.H{"message": "resource deleted successfully"})
}
