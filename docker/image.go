package docker

import (
	"GoToKube/logger"
	"context"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types/image"
)

// GetImages 获取当前的 Docker 镜像列表
func GetImages() ([]image.Summary, error) {
	images, err := dockerClient.ImageList(context.Background(), image.ListOptions{})
	if err != nil {
		logger.GlobalLogger.Error(err.Error())
	} else {
		logger.GlobalLogger.Info("Success to get images")
	}
	return images, err
}

// UploadImage 上传镜像
func UploadImage(filePath string) error {
	// 打开镜像文件
	file, err := os.Open(filePath)
	if err != nil {
		logger.GlobalLogger.Error(fmt.Sprintf("Failed to open image file: %s", err))
		return err
	}
	defer file.Close()
	// 上传镜像
	response, err := dockerClient.ImageLoad(context.Background(), file, true)
	if err != nil {
		logger.GlobalLogger.Error(fmt.Sprintf("Failed to load image: %s", err))
		return err
	}
	defer response.Body.Close()
	// 读取上传响应
	_, err = io.Copy(os.Stdout, response.Body)
	if err != nil {
		logger.GlobalLogger.Error(fmt.Sprintf("Failed to read response: %s", err))
		return err
	}
	logger.GlobalLogger.Info("Image was successfully uploaded")
	return nil
}

// DeleteImage 删除镜像
func DeleteImage(imageID string) error {
	_, err := dockerClient.ImageRemove(context.Background(), imageID, image.RemoveOptions{Force: true, PruneChildren: true})
	if err != nil {
		logger.GlobalLogger.Error(fmt.Sprintf("Failed to delete image: %s", err))
		return err
	}
	logger.GlobalLogger.Info("Image deleted successfully")
	return nil
}
