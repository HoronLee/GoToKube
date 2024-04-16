package docker

import (
	"VDController/logger"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/docker/docker/client"
)

var tLogger = logger.NewLogger(logger.INFO)

func CheckState() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		fmt.Println("Failed to create Docker client:", err)
		os.Exit(1)
	}

	ifok, state := dockerChecks(cli)
	if ifok {
		fmt.Println(state)
	} else {
		fmt.Println(state)
		os.Exit(1)
	}
}

var eInfo = make(map[string]interface{})

func dockerChecks(cli *client.Client) (ifok bool, state string) {
	ctx := context.Background()

	// 检查 Docker 是否在运行
	_, err := cli.Info(ctx)
	if err != nil {
		ifok, state = false, "Docker 不在运行，请先启动 Docker! \n"+"请参考 https://docs.docker.com/engine/install/ 安装 Docker。"
		tLogger.Log(logger.ERROR, state)
		return ifok, state
	}
	// 检查 Docker 版本
	sVersion, err := cli.ServerVersion(ctx)
	if err != nil {
		ifok, state = false, "无法获取 Docker 版本。"
		tLogger.Log(logger.ERROR, state)
		return ifok, state
	} else {
		eInfo["DockerVersion"] = sVersion.Version
		dstate := "Docker 版本:" + sVersion.Version
		// 检查 Docker Compose 版本
		dockerCompV, err := exec.Command("docker-compose", "version").Output()
		if err != nil {
			ifok, state = false, "无法获取 Docker Compose 版本，将无法使用Docker Compose功能，\n"+"请参考 https://docs.docker.com/compose/install/ 安装 Docker Compose。"
			tLogger.Log(logger.WARNING, state)
			return ifok, state
		} else {
			versionIndex := strings.Index(string(dockerCompV), "version")
			if versionIndex != -1 {
				versionStr := strings.TrimSpace(string(dockerCompV)[versionIndex+len("version v"):])
				ifok, state = true, dstate+", "+"Docker Compose 版本:"+versionStr
				eInfo["DcomposeVersion"] = versionStr
			} else {
				ifok, state = true, dstate+"\n"+"无法获取 Docker Compose 版本"
			}
		}
		tLogger.Log(logger.INFO, state)
		return ifok, state
	}
}

func GetEnvInfo() map[string]interface{} {
	return eInfo
}
