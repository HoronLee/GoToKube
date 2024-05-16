package kubernetes

import (
	//"k8s.io/client-go/kubernetes"
	"VDController/logger"
	"encoding/json"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var kLogger = logger.NewLogger(logger.INFO)

func InitK8s() {
	config, _ := createKubeConfig()
	outPut, _ := json.Marshal(config)
	kLogger.Log(logger.INFO, string(outPut))
}
func createKubeConfig() (*rest.Config, error) {
	kubeconfigPath := "~/.kube/config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		return nil, err
	}
	return config, nil
}
