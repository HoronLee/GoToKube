package docker

import (
	"GoToKube/logger"
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"os/exec"
	"strings"

	"github.com/docker/docker/client"
)

func CheckStatus() bool {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		logger.GlobalLogger.Error(err.Error())
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
		logger.GlobalLogger.Log(logrus.ErrorLevel, status)
		return ifok, status
	}
	// 检查 Docker 版本
	sVersion, err := cli.ServerVersion(ctx)
	// TODO: 将 Docker 信息写入数据表
	dockerStatus := "Docker version:" + sVersion.Version
	if err != nil {
		ifok, status = false, "Unable to get Docker version: "+err.Error()
		logger.GlobalLogger.Log(logrus.ErrorLevel, status)
		return ifok, status
	} else {
		// 检查 Docker Compose 版本
		dockerCVersion, err := exec.Command("docker", "compose", "version").Output()
		if err != nil {
			ifok, status = true, "Unable to get the Docker Compose version, you will not be able to use the Docker Compose feature，\n"+"See https://docs.docker.com/compose/install/ to install Docker Compose:"+err.Error()
			logger.GlobalLogger.Log(logrus.WarnLevel, status)
			return ifok, status
		} else {
			versionIndex := strings.Index(string(dockerCVersion), "version")
			if versionIndex != -1 {
				versionStr := strings.TrimSpace(string(dockerCVersion)[versionIndex+len("version v"):])
				ifok, status = true, dockerStatus+", "+"Docker Compose 版本:"+versionStr
				// TODO: 将 DockerCompose 信息写入数据表
			} else {
				ifok, status = true, dockerStatus+"\n"+"Unable to get Docker Compose version."
			}
		}
		logger.GlobalLogger.Log(logrus.InfoLevel, status)
		return ifok, status
	}
}
