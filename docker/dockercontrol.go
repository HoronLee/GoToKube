package docker

import (
	"VDController/logger"
	"context"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

var (
	dLogger = logger.NewLogger(logger.INFO)
	// 全局 docker 客户端
	Dockerclient *dockerClient
)

// dockerClient 包含 Docker 客户端
type dockerClient struct {
	Client *client.Client
}

func init() {
	var err error
	Dockerclient, err = newClient()
	if err != nil {
		dLogger.Log(logger.ERROR, "DockerClient 创建失败")
	} else {
		dLogger.Log(logger.INFO, "DockerClient 成功创建")
	}
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

// 获取当前容器
func (dc *dockerClient) Dockerls() ([]types.Container, error) {
	containers, err := dc.Client.ContainerList(context.Background(), container.ListOptions{})
	if err != nil {
		dLogger.Log(logger.ERROR, "获取容器失败")
	} else {
		dLogger.Log(logger.INFO, "获取当前容器")
	}
	return containers, err
}

// 通过镜像名获得容器
func (dc *dockerClient) DockerlsByImg(imgName string) (outPut map[string]interface{}, error error) {
	outPut = make(map[string]interface{})
	ctrInfo, _ := dc.Dockerls()
	for _, ctr := range ctrInfo {
		if strings.Contains(ctr.Image, imgName) {
			outPut[ctr.Image] = ctr
		}
	}
	if len(outPut) == 0 {
		outPut["WARN"] = "No Container matches this condition."
	}
	return outPut, error
}
