package docker

import (
	"GoToKube/logger"
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"io"
	"os"
	"strings"
)

var (
	dockerClient *client.Client
)

func initDocker() {
	var err error
	dockerClient, err = newClient()
	if err != nil {
		logger.GlobalLogger.Error("Docker Client creation failed" + err.Error())
	} else {
		logger.GlobalLogger.Info("Docker Client was successfully created")
	}
}

// NewClient 创建一个包含 Docker 客户端的新实例
func newClient() (*client.Client, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}
	return cli, nil
}

// Close 关闭 Docker 客户端连接
func Close() error {
	return dockerClient.Close()
}

// GetCtr 获取当前容器
func GetCtr() ([]types.Container, error) {
	containers, err := dockerClient.ContainerList(context.Background(), container.ListOptions{})
	if err != nil {
		logger.GlobalLogger.Error("Failed to get containers" + err.Error())
	} else {
		logger.GlobalLogger.Info("Success to get containers")
	}
	return containers, err
}

// GetCtrByImg 通过镜像名获得容器
func GetCtrByImg(imgName string) ([]types.Container, error) {
	containers, err := GetCtr()
	if err != nil {
		return nil, err
	}
	var output []types.Container
	for _, ctr := range containers {
		if strings.Contains(ctr.Image, imgName) {
			output = append(output, ctr)
		}
	}
	if len(output) == 0 {
		return output, fmt.Errorf("no container matches this condition")
	}
	return output, nil
}

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
