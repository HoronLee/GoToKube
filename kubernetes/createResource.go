package kubernetes

import (
	"context"
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CreateNamespace 创建指定名称的命名空间
func CreateNamespace(namespace string) (*v1.Namespace, error) {
	if namespace == "" {
		return &v1.Namespace{}, fmt.Errorf("namespace name is required")
	}

	ns := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: namespace,
		},
	}
	result, err := kubeClient.CoreV1().Namespaces().Create(context.TODO(), ns, metav1.CreateOptions{})
	if err != nil {
		return &v1.Namespace{}, err
	}
	return result, nil
}
