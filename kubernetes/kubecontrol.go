package kubernetes

import (
	"VDController/database"
	"VDController/web/models"
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetAllPods() {
	pods, err := kubeClient.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
	for _, pod := range pods.Items {
		fmt.Printf("Namespace: %s, Name: %s\n", pod.Namespace, pod.Name)
	}
}

func GetK8sVersion() error {
	version, err := kubeClient.Discovery().ServerVersion()
	if err != nil {
		return err
	}
	EnvInfo.KubeVersion = version.String()
	database.SaveOrUpdateStatusInfo(models.StatusInfo{Component: "Kubernetes", Version: version.String(), Status: "OK"})
	return nil
}
