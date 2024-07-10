package kubernetes

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"io"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/util/retry"
	"os"
	"path/filepath"
	"strings"
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
	return nil
}

func GetDeployments(namespace string) (interface{}, error) {
	deployments, err := kubeClient.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return deployments, nil
}

func GetDeployment(name, namespace string) (interface{}, error) {
	deployment, err := kubeClient.AppsV1().Deployments(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return deployment, nil
}
func GetServices(namespace string) (interface{}, error) {
	if namespace == "" {
		namespace = metav1.NamespaceAll
	}
	services, err := kubeClient.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return services, nil
}

func GetPods(namespace string) (*v1.PodList, error) {
	pods, err := kubeClient.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return pods, nil
}

func GetPod(name, namespace string) (*v1.Pod, error) {
	pod, err := kubeClient.CoreV1().Pods(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return pod, nil
}
func GetNamespaces() (interface{}, error) {
	namespaces, err := kubeClient.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return namespaces, nil
}

func IsYAML(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".yml" || ext == ".yaml"
}

// ApplyYAML 通过 yaml 文件动态创建集群资源
func ApplyYAML(filePath string) error {
	if !IsYAML(filePath) {
		return fmt.Errorf("file %s is not a valid YAML file", filePath)
	}
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()
	decoder := yaml.NewYAMLOrJSONDecoder(file, 4096)
	for {
		var u unstructured.Unstructured
		if err := decoder.Decode(&u); err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("failed to decode YAML: %v", err)
		}
		gvk := u.GroupVersionKind()
		resource := schema.GroupVersionResource{
			Group:    gvk.Group,
			Version:  gvk.Version,
			Resource: strings.ToLower(gvk.Kind) + "s",
		}
		namespace := u.GetNamespace()
		if namespace == "" {
			namespace = "default"
		}
		resourceClient := dynamicClient.Resource(resource).Namespace(namespace)
		_, err = resourceClient.Create(context.TODO(), &u, metav1.CreateOptions{})
		if errors.IsAlreadyExists(err) {
			retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
				existing, getErr := resourceClient.Get(context.TODO(), u.GetName(), metav1.GetOptions{})
				if getErr != nil {
					return getErr
				}
				u.SetResourceVersion(existing.GetResourceVersion())
				_, updateErr := resourceClient.Update(context.TODO(), &u, metav1.UpdateOptions{})
				return updateErr
			})
			if retryErr != nil {
				return fmt.Errorf("update failed: %v", retryErr)
			}
		} else if err != nil {
			return fmt.Errorf("create failed: %v", err)
		}
	}
	return nil
}

func DeleteYAML(filePath string) error {
	if !IsYAML(filePath) {
		return fmt.Errorf("file %s is not a valid YAML file", filePath)
	}
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	decoder := yaml.NewYAMLOrJSONDecoder(file, 4096)
	for {
		var u unstructured.Unstructured
		if err := decoder.Decode(&u); err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("failed to decode YAML: %v", err)
		}

		gvk := u.GroupVersionKind()
		resource := schema.GroupVersionResource{
			Group:    gvk.Group,
			Version:  gvk.Version,
			Resource: strings.ToLower(gvk.Kind) + "s",
		}

		namespace := u.GetNamespace()
		if namespace == "" {
			namespace = "default"
		}

		resourceClient := dynamicClient.Resource(resource).Namespace(namespace)

		err = resourceClient.Delete(context.TODO(), u.GetName(), metav1.DeleteOptions{})
		if err != nil {
			return fmt.Errorf("delete failed: %v", err)
		}
	}

	return nil
}
