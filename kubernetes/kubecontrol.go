package kubernetes

import (
	"VDController/database"
	"VDController/web/models"
	"context"
	"fmt"
	"gorm.io/gorm"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	db *gorm.DB
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

func Getk8sVersion() error {
	version, err := kubeClient.Discovery().ServerVersion()
	if err != nil {
		return err
	}
	EnvInfo.KubeVersion = version.String()
	db, _ := database.GetDBConnection()
	db.Create(&models.StatusInfo{Component: "Kubernetes", Version: "UnKnown", Status: "Running"})
	var k8sModel models.StatusInfo
	db.First(&k8sModel, "component = ?", "Kubernetes")
	db.Model(&k8sModel).Update("version", version.String())
	return nil
}
