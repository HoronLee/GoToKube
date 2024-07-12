package controller

import (
	"GoToKube/kubernetes"
	"GoToKube/logger"
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
	uploadDir := "./uploads/yaml"
	dst := uploadDir + "/" + file.Filename
	// 目录不存在则新建
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		if err := os.Mkdir(uploadDir, 0755); err != nil {
			os.Exit(0)
		}
	}
	// 检查是否存在同名文件
	if _, err := os.Stat(dst); err == nil {
		// 删除同名文件
		if err := os.Remove(dst); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("failed to delete existing file: %s", err.Error()))
			return
		}
		logger.GlobalLogger.Info("deleted existing file: " + dst)
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
		logger.GlobalLogger.Info("deleted file after apply failure: " + dst)
		return
	}
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded and resources created!", file.Filename))
}

// DeleteYaml 删除 YAML 文件并删除集群资源
func DeleteYaml(c *gin.Context) {
	fileName := c.Param("file")
	filePath := filepath.Join("./uploads/yaml", fileName)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "file not found"})
		return
	}
	err := kubernetes.DeleteYAML(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "file not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 删除文件
	if err := os.Remove(filePath); err != nil {
		if os.IsNotExist(err) {
			c.JSON(http.StatusOK, gin.H{"message": "resource deleted successfully and file already deleted"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	logger.GlobalLogger.Info("deleted file: " + fileName)
	c.JSON(http.StatusOK, gin.H{"message": "resource and file deleted successfully"})
}

func ListYamlFiles(c *gin.Context) {
	uploadDir := "./uploads/yaml"
	// 检查目录是否存在
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		// 目录不存在则新建
		if err := os.MkdirAll(uploadDir, 0755); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("failed to create directory: %s", err.Error()))
			return
		}
	}

	// 读取目录内容
	files, err := os.ReadDir(uploadDir)
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
