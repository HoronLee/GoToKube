package kubernetes

import (
	"GoToKube/logger"
	"context"
	"fmt"
	"io"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/yaml"
	"os"
	"path/filepath"
	"strings"
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

// ApplyYAML 通过读取并应用给定路径下的 YAML 文件来创建或更新 Kubernetes 资源。
// filePath 参数指定 YAML 文件的路径。
// 返回错误如果文件不是有效的 YAML 文件，或者在处理文件期间遇到任何问题。
func ApplyYAML(filePath string) error {
	// 检查文件是否为有效的 YAML 文件。
	if !IsYAML(filePath) {
		return fmt.Errorf("file %s is not a valid YAML file", filePath)
	}
	// 打开 YAML 文件。
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close() // 确保文件在函数返回前关闭。

	// 创建一个解码器，用于解析 YAML 或 JSON 格式的文件。
	decoder := yaml.NewYAMLOrJSONDecoder(file, 4096)

	for {
		// 解码器读取并解析 YAML 文件中的下一个对象。
		var u unstructured.Unstructured
		if err := decoder.Decode(&u); err != nil {
			if err == io.EOF {
				// 文件结束，正常退出循环。
				break
			}
			// 解析错误，返回错误信息。
			return fmt.Errorf("failed to decode YAML: %v", err)
		}

		// 获取解码后的对象的 GroupVersionKind 信息。
		gvk := u.GroupVersionKind()
		// 根据 GroupVersionKind 信息构造资源的 GroupVersionResource。
		resource := schema.GroupVersionResource{
			Group:    gvk.Group,
			Version:  gvk.Version,
			Resource: strings.ToLower(gvk.Kind) + "s",
		}

		// 获取对象的命名空间，如果未指定，则使用默认命名空间。
		namespace := u.GetNamespace()
		if namespace == "" {
			namespace = "default"
		}

		// 根据资源信息和命名空间，获取动态客户端，用于创建或更新资源。
		resourceClient := dynamicClient.Resource(resource).Namespace(namespace)

		// 尝试创建资源，如果资源已存在，则更新资源。
		_, err = resourceClient.Create(context.TODO(), &u, metav1.CreateOptions{})
		if errors.IsAlreadyExists(err) {
			// 资源已存在，获取现有资源的版本信息，用于更新资源。
			existing, getErr := resourceClient.Get(context.TODO(), u.GetName(), metav1.GetOptions{})
			if getErr != nil {
				// 获取资源时出错，返回错误信息。
				return getErr
			}
			u.SetResourceVersion(existing.GetResourceVersion())
			// 更新资源。
			_, updateErr := resourceClient.Update(context.TODO(), &u, metav1.UpdateOptions{})
			if updateErr != nil {
				// 更新资源时出错，返回错误信息。
				return fmt.Errorf("update failed: %v", updateErr)
			}
		} else if err != nil {
			// 创建资源时出错，返回错误信息。
			return fmt.Errorf("create failed: %v", err)
		}
	}
	// 成功处理文件，无错误返回。
	return nil
}

// DeleteYAML 通过 yaml 文件删除集群资源
// DeleteYAML 删除指定路径下的YAML文件中定义的所有Kubernetes资源。
// filePath: YAML文件的路径。
// 返回值: 如果文件不是有效的YAML文件或删除资源过程中出现错误，则返回错误；否则返回nil。
func DeleteYAML(filePath string) error {
	// 检查文件是否为有效的YAML文件。
	if !IsYAML(filePath) {
		return fmt.Errorf("file %s is not a valid YAML file", filePath)
	}
	// 打开文件以进行读取。
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()
	// 使用YAML或JSON解码器解析文件内容。
	decoder := yaml.NewYAMLOrJSONDecoder(file, 4096)
	for {
		// 解码YAML对象到Unstructured对象。
		var u unstructured.Unstructured
		if err := decoder.Decode(&u); err != nil {
			if err == io.EOF {
				// 文件读取结束，退出循环。
				break
			}
			// 解码失败，返回错误。
			return fmt.Errorf("failed to decode YAML: %v", err)
		}
		// 获取Unstructured对象的GroupVersionKind信息。
		gvk := u.GroupVersionKind()
		// 根据GVK信息构造GroupVersionResource，用于资源操作。
		resource := schema.GroupVersionResource{
			Group:    gvk.Group,
			Version:  gvk.Version,
			Resource: strings.ToLower(gvk.Kind) + "s",
		}
		// 获取资源的命名空间，默认为"default"。
		namespace := u.GetNamespace()
		if namespace == "" {
			namespace = "default"
		}
		// 创建动态客户端，用于删除资源。
		resourceClient := dynamicClient.Resource(resource).Namespace(namespace)
		// 删除资源，并处理可能的错误。
		err = resourceClient.Delete(context.TODO(), u.GetName(), metav1.DeleteOptions{})
		if errors.IsNotFound(err) {
			// 资源不存在，记录日志并继续处理下一个资源。
			logger.GlobalLogger.Info(fmt.Sprintf("resource %s not found, skipping", u.GetName()))
			continue
		} else if err != nil {
			// 删除资源失败，返回错误。
			return fmt.Errorf("delete failed: %v", err)
		}
	}
	// 所有资源都成功删除后，删除YAML文件。
	return os.Remove(filePath)
}
