package kubernetes

import (
	"VDController/config"
	"VDController/logger"
	"flag"
	"fmt"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var (
	kubeClient *kubernetes.Clientset
	EnvInfo    = Info{}
)

type Info struct {
	KubeVersion string `json:"kubeVersion"`
}

func CheckStatus() bool {
	if config.ConfigData.KubeEnable {
		fmt.Println("⚓️已启用 kubenetes 控制器")
	} else {
		fmt.Println("⚓️不启用 kubenetes 控制器")
		return true
	}
	// 获取 kubernetes 配置文件
	kubeconfig := config.ConfigData.KubeconfigPath
	if kubeconfig == "" {
		if home := homedir.HomeDir(); home != "" {
			kubeconfig = filepath.Join(home, ".kube", "config")
		}
		flag.StringVar(&kubeconfig, "kubeconfig", kubeconfig, "(optional) absolute path to the kubeconfig file")
	}
	flag.Parse()
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		logger.GlobalLogger.Log(logger.ERROR, err.Error())
		return false
	}
	// 创建 kubernetes 客户端
	kubeClient, err = kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		logger.GlobalLogger.Log(logger.ERROR, err.Error())
		return false
	} else {
		logger.GlobalLogger.Log(logger.INFO, "Kubernetes Client Create Success")
		GetK8sVersion()
		return true
	}
}
