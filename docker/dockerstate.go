package docker

import (
	"VDController/logger"
	"context"
	"fmt"
	"os/exec"
	"strings"

	"github.com/docker/docker/client"
)

var EnvInfo = Info{}

type Info struct {
	DockerVersion  string `json:"dockerVersion"`
	DockerCVersion string `json:"dockerComposeVersion"`
}

func CheckStatus() bool {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		fmt.Println("Failed to create Docker client:", err)
		return false
	}
	ifok, status := dockerChecks(cli)
	if !ifok {
		fmt.Println(status)
		return false
	}
	cli.Close()
	initDocker()
	return true
}

func dockerChecks(cli *client.Client) (ifok bool, status string) {
	ctx := context.Background()
	// 检查 Docker 是否在运行
	_, err := cli.Info(ctx)
	if err != nil {
		ifok, status = false, fmt.Sprint(err)
		logger.GlobalLogger.Log(logger.ERROR, status)
		return ifok, status
	}
	// 检查 Docker 版本
	sVersion, err := cli.ServerVersion(ctx)
	if err != nil {
		ifok, status = false, "Unable to get Docker version."
		logger.GlobalLogger.Log(logger.ERROR, status)
		return ifok, status
	} else {
		EnvInfo.DockerVersion = string(sVersion.Version)
		dstatus := "Docker version:" + sVersion.Version
		// 检查 Docker Compose 版本
		dockerCompV, err := exec.Command("docker", "compose", "version").Output()
		if err != nil {
			ifok, status = false, "Unable to get the Docker Compose version, you will not be able to use the Docker Compose feature，\n"+"See https://docs.docker.com/compose/install/ to install Docker Compose."
			logger.GlobalLogger.Log(logger.WARNING, status)
			return ifok, status
		} else {
			versionIndex := strings.Index(string(dockerCompV), "version")
			if versionIndex != -1 {
				versionStr := strings.TrimSpace(string(dockerCompV)[versionIndex+len("version v"):])
				ifok, status = true, dstatus+", "+"Docker Compose 版本:"+versionStr
				EnvInfo.DockerCVersion = versionStr
			} else {
				ifok, status = true, dstatus+"\n"+"Unable to get Docker Compose version."
			}
		}
		logger.GlobalLogger.Log(logger.INFO, status)
		return ifok, status
	}
}
