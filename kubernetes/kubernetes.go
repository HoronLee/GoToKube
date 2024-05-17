package kubernetes

import (
	"VDController/config"
	"VDController/logger"
	"encoding/json"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var kLogger = logger.NewLogger(logger.INFO)

func InitK8s() {
	config, err := createKubeConfig()
	if err != nil {
		kLogger.Log(logger.ERROR, err.Error())
		return
	}
	outPut, _ := json.Marshal(config)
	// fmt.Println(outPut)
	kLogger.Log(logger.INFO, string(outPut))
}
func createKubeConfig() (*rest.Config, error) {
	kubeconfigPath := config.ConfigData.KubeconfigPath
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		return nil, err
	}
	return config, nil
}
