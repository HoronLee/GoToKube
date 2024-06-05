package kubernetes

import (
	"VDController/config"
	"VDController/logger"
	"flag"
	"os"
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

func Checkstatus() {
	// 获取 kubernetes 配置文件
	var kubeconfig string
	kubeconfig = config.ConfigData.KubeconfigPath
	if kubeconfig == "" {
		if home := homedir.HomeDir(); home != "" {
			kubeconfig = filepath.Join(home, ".kube", "config")
		}
		flag.StringVar(&kubeconfig, "kubeconfig", kubeconfig, "(optional) absolute path to the kubeconfig file")
	}
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		logger.GlobalLogger.Log(logger.ERROR, err.Error())
		return
	}
	// 创建 kubernetes 客户端
	kubeClient, err = kubernetes.NewForConfig(config)
	if err != nil {
		logger.GlobalLogger.Log(logger.ERROR, err.Error())
		os.Exit(1)
	} else {
		logger.GlobalLogger.Log(logger.INFO, "Kubernetes Client Create Success")
		GetK8sVersion()
	}
}
