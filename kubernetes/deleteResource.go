package kubernetes

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func DeleteNamespace(namespace string) error {
	err := kubeClient.CoreV1().Namespaces().Delete(context.TODO(), namespace, metav1.DeleteOptions{})
	if err != nil {
		return err
	}
	return nil
}

func DeleteDeployment(namespace, deploymentName string) error {
	err := kubeClient.AppsV1().Deployments(namespace).Delete(context.TODO(), deploymentName, metav1.DeleteOptions{})
	if err != nil {
		return err
	}
	return nil
}

func DeleteService(namespace, serviceName string) error {
	err := kubeClient.CoreV1().Services(namespace).Delete(context.TODO(), serviceName, metav1.DeleteOptions{})
	if err != nil {
		return err
	}
	return nil
}

func DeletePod(namespace, podName string) error {
	err := kubeClient.CoreV1().Pods(namespace).Delete(context.TODO(), podName, metav1.DeleteOptions{})
	if err != nil {
		return err
	}
	return nil
}
