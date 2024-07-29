package kubernetes

import (
	"GoToKube/config"
	"GoToKube/logger"
	"flag"
	"k8s.io/client-go/dynamic"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var (
	kubeClient    *kubernetes.Clientset
	dynamicClient dynamic.Interface
)

func CheckStatus() bool {
	// 获取 kubernetes 配置文件
	KubeConfig := config.Data.Kube.ConfigPath
	if KubeConfig == "" {
		if home := homedir.HomeDir(); home != "" {
			KubeConfig = filepath.Join(home, ".kube", "config")
		}
		flag.StringVar(&KubeConfig, "KubeConfig", KubeConfig, "(optional) absolute path to the KubeConfig file")
	}
	flag.Parse()
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", KubeConfig)
	if err != nil {
		logger.GlobalLogger.Error(err.Error())
		return false
	}
	// 创建 kubernetes 客户端
	kubeClient, err = kubernetes.NewForConfig(kubeConfig)
	dynamicClient, err = dynamic.NewForConfig(kubeConfig)
	if err != nil {
		logger.GlobalLogger.Error(err.Error())
		return false
	} else {
		err = GetK8sVersion()
		if err != nil {
			logger.GlobalLogger.Error(err.Error())
			return false
		}
		return true
	}
}

func GetK8sVersion() error {
	_, err := kubeClient.Discovery().ServerVersion()
	if err != nil {
		return err
	}
	// TODO: 将 DockerCompose 信息写入数据表
	return nil
}
