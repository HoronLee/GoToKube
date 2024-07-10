package kubernetes

import (
	"VDController/config"
	"VDController/logger"
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
	EnvInfo       = Info{}
)

type Info struct {
	KubeVersion string `json:"kubeVersion"`
}

func CheckStatus() bool {
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
	dynamicClient, err = dynamic.NewForConfig(kubeConfig)
	if err != nil {
		logger.GlobalLogger.Log(logger.ERROR, err.Error())
		return false
	} else {
		err = Getk8sVersion()
		if err != nil {
			logger.GlobalLogger.Log(logger.ERROR, err.Error())
			return false
		}
		return true
	}
}
