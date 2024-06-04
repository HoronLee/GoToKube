package kubernetes

import (
	"VDController/config"
	"VDController/logger"
	"flag"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var (
	kLogger    = logger.NewLogger(logger.INFO)
	kubeClient *kubernetes.Clientset
	EnvInfo = Info{}
)

type Info struct {
	KubeVersion string `json:"kubeVersion"`
}
func InitKubernetes() {
	// 获取 kubernetes 配置文件
	var kubeconfig string
	kubeConfigPath := config.ConfigData.KubeconfigPath
	if kubeConfigPath == "" {
		if home := homedir.HomeDir(); home != "" {
			kubeconfig = filepath.Join(home, ".kube", "config")
		}
	} else {
		kubeconfig = kubeConfigPath
	}
	// 解析 kubeconfig 文件路径
	flag.StringVar(&kubeconfig, "kubeconfig", kubeconfig, "(optional) absolute path to the kubeconfig file")
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		kLogger.Log(logger.ERROR, err.Error())
		return
	}
	// 创建 kubernetes 客户端
	kubeClient, err = kubernetes.NewForConfig(config)
	if err != nil {
		kLogger.Log(logger.ERROR, err.Error())
		return
	} else {
		kLogger.Log(logger.INFO, "Kubernetes Client Create Success")
	}
	GetK8sVersion()
}
