package docker

import (
	"VDController/logger"
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

var dLogger = logger.NewLogger(logger.INFO)

// dockerClient 包含 Docker 客户端
type dockerClient struct {
	Client *client.Client
}

// NewClient 创建一个包含 Docker 客户端的新实例
func newClient() (*dockerClient, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}
	return &dockerClient{Client: cli}, nil
}

// Close 关闭 Docker 客户端连接
func (dc *dockerClient) Close() error {
	return dc.Client.Close()
}

// 全局 docker 客户端
var Dockerclient *dockerClient

func init() {
	var err error
	Dockerclient, err = newClient()
	if err != nil {
		dLogger.Log(logger.ERROR, "DockerClient 创建失败 ")
	} else {
		dLogger.Log(logger.INFO, "DockerClient 成功创建 ")
	}
}

// 列出当前容器
func (dc *dockerClient) Dockerls() ([]types.Container, error) {
	containers, err := dc.Client.ContainerList(context.Background(), container.ListOptions{})
	if err != nil {
		dLogger.Log(logger.ERROR, "列出容器失败")
	} else {
		dLogger.Log(logger.INFO, "列出当前容器")
	}
	return containers, err
}
