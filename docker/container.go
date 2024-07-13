package docker

import (
	"GoToKube/logger"
	"context"
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/go-connections/nat"
	"time"
)

// CreateContainer 创建容器
func CreateContainer(imageName string, containerName string, cmd []string, portBindings map[string]string, volumes map[string]string) (string, error) {
	// 配置容器配置
	config := &container.Config{
		Image: imageName,
	}
	// 如果提供了命令，则设置命令
	if len(cmd) > 0 {
		config.Cmd = cmd
	}
	// 配置端口映射
	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{},
		Mounts:       []mount.Mount{},
	}
	for hostPort, containerPort := range portBindings {
		hostConfig.PortBindings[nat.Port(containerPort)] = []nat.PortBinding{
			{
				HostPort: hostPort,
			},
		}
	}
	// 如果提供了卷挂载，则设置卷挂载
	if volumes != nil {
		for hostPath, containerPath := range volumes {
			hostConfig.Mounts = append(hostConfig.Mounts, mount.Mount{
				Type:   mount.TypeBind,
				Source: hostPath,
				Target: containerPath,
			})
		}
	}
	// 创建容器
	resp, err := dockerClient.ContainerCreate(context.Background(), config, hostConfig, nil, nil, containerName)
	if err != nil {
		logger.GlobalLogger.Log(logger.ERROR, fmt.Sprintf("Failed to create container: %s", err))
		return "", err
	}
	// 启动容器
	if err := dockerClient.ContainerStart(context.Background(), resp.ID, container.StartOptions{}); err != nil {
		logger.GlobalLogger.Log(logger.ERROR, fmt.Sprintf("Failed to start container: %s", err))
		return "", err
	}
	logger.GlobalLogger.Log(logger.INFO, "Container created and started successfully")
	return resp.ID, nil
}

// DeleteContainer 删除指定的 Docker 容器
func DeleteContainer(containerID string) error {
	// 停止容器
	if err := dockerClient.ContainerStop(context.Background(), containerID, container.StopOptions{}); err != nil {
		logger.GlobalLogger.Log(logger.ERROR, "Failed to stop container: "+err.Error())
		return err
	}
	// 删除容器
	if err := dockerClient.ContainerRemove(context.Background(), containerID, container.RemoveOptions{Force: true}); err != nil {
		logger.GlobalLogger.Log(logger.ERROR, "Failed to remove container: "+err.Error())
		return err
	}
	logger.GlobalLogger.Log(logger.INFO, "Container deleted successfully")
	return nil
}

// StopContainer 停止指定的 Docker 容器
func StopContainer(containerID string) (string, error) {
	// 创建带超时的上下文，10s 过期
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 停止容器
	if err := dockerClient.ContainerStop(ctx, containerID, container.StopOptions{}); err != nil {
		logger.GlobalLogger.Log(logger.ERROR, fmt.Sprintf("Failed to stop container %s: %v", containerID, err))
		return containerID, err
	}

	logger.GlobalLogger.Log(logger.INFO, fmt.Sprintf("Container %s stopped successfully", containerID))
	return containerID, nil
}

// StartContainer 启动指定的 Docker 容器
func StartContainer(containerID string) (string, error) {
	// 创建带超时的上下文，10s 过期
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 启动容器
	if err := dockerClient.ContainerStart(ctx, containerID, container.StartOptions{}); err != nil {
		logger.GlobalLogger.Log(logger.ERROR, fmt.Sprintf("Failed to start container %s: %v", containerID, err))
		return containerID, err
	}
	return containerID, nil
}
