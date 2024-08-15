package docker

import (
	"GoToKube/logger"

	"github.com/docker/docker/client"
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
