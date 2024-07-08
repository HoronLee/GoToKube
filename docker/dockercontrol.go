package docker

import (
	"VDController/logger"
	"context"
	"fmt"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

var (
	dockerClient *client.Client
)

func initDocker() {
	var err error
	dockerClient, err = newClient()
	if err != nil {
		logger.GlobalLogger.Log(logger.ERROR, "Docker Client creation failed")
	} else {
		logger.GlobalLogger.Log(logger.INFO, "Docker Client was successfully created")
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

// ContainerLs 获取当前容器
func ContainerLs() ([]types.Container, error) {
	containers, err := dockerClient.ContainerList(context.Background(), container.ListOptions{})
	if err != nil {
		logger.GlobalLogger.Log(logger.ERROR, "Failed to get containers")
	} else {
		logger.GlobalLogger.Log(logger.INFO, "Success to get containers")
	}
	return containers, err
}

// ContainerLsByImg 通过镜像名获得容器
func ContainerLsByImg(imgName string) ([]types.Container, error) {
	containers, err := ContainerLs()
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
